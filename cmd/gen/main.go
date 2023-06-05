package main

import (
	"flag"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"go.elara.ws/go-lemmy/cmd/gen/generator"
	"go.elara.ws/go-lemmy/cmd/gen/parser"
	"go.elara.ws/logger"
	"go.elara.ws/logger/log"
)

func init() {
	log.Logger = logger.NewPretty(os.Stderr)
}

var implDirs = [...]string{
	"crates/api_crud/src",
	"crates/apub/src/api",
	"crates/api/src",
}

var structDirs = [...]string{
	"crates/api_common",
	"crates/db_schema/src/source",
	"crates/db_views_actor/src/structs.rs",
	"crates/db_views/src/structs.rs",
	"crates/db_views_moderator/src/structs.rs",
	"crates/db_schema/src/aggregates/structs.rs",
	"crates/db_schema/src/lib.rs",
	"crates/websocket/src/lib.rs",
}

const routesFile = "src/api_routes_http.rs"

func main() {
	lemmyDir := flag.String("lemmy-dir", "lemmy", "Path to Lemmy repository")
	outDir := flag.String("out-dir", "out", "Directory to write output in")
	flag.Parse()

	baseStructDir := filepath.Join(*outDir, "types")
	sp := parser.NewStruct(nil)
	sp.Skip = []string{"LemmyContext", "Recipient", "WsMessage", "Connect", "SessionInfo"}
	for _, structDir := range structDirs {
		dir := filepath.Join(*lemmyDir, structDir)
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			if d.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".rs" {
				return nil
			}

			name := d.Name()
			if name == "context.rs" ||
				name == "local_user_language.rs" ||
				name == "chat_server.rs" {
				return nil
			}

			fl, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fl.Close()

			sp.Reset(fl)
			fileStructs, err := sp.Parse()
			if err != nil {
				return err
			}

			nameNoExt := strings.TrimSuffix(d.Name(), ".rs")
			goFilePath := filepath.Join(baseStructDir, nameNoExt+".gen.go")

			i := 1
			_, err = os.Stat(goFilePath)
			for err == nil {
				goFilePath = filepath.Join(baseStructDir, nameNoExt+"."+strconv.Itoa(i)+".gen.go")
				_, err = os.Stat(goFilePath)
				i++
			}

			outFl, err := os.Create(goFilePath)
			if err != nil {
				return err
			}
			defer outFl.Close()

			_, err = outFl.WriteString("//  Source: " + path + "\n")
			if err != nil {
				return err
			}

			return generator.NewStruct(outFl, "types").Generate(fileStructs)
		})
		if err != nil {
			log.Fatal("Error walking directory while parsing structs").Err(err).Str("dir", dir).Send()
		}
	}

	ip := parser.NewImpl(nil)
	impls := map[string]string{}
	for _, implDir := range implDirs {
		dir := filepath.Join(*lemmyDir, implDir)
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			if d.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".rs" {
				return nil
			}

			fl, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fl.Close()

			ip.Reset(fl)
			implMap, err := ip.Parse()
			if err != nil {
				return err
			}

			for reqStruct, resStruct := range implMap {
				impls[reqStruct] = resStruct
			}

			return nil
		})
		if err != nil {
			log.Fatal("Error walking directory while parsing impls").Err(err).Str("dir", dir).Send()
		}
	}

	routesPath := filepath.Join(*lemmyDir, routesFile)
	rf, err := os.Open(routesPath)
	if err != nil {
		log.Fatal("Error opening routes file").Err(err).Send()
	}
	defer rf.Close()

	rp := parser.NewRoutes(rf)
	routes, err := rp.Parse()
	if err != nil {
		log.Fatal("Error parsing routes file").Err(err).Send()
	}

	orf, err := os.Create(filepath.Join(*outDir, "routes.gen.go"))
	if err != nil {
		log.Fatal("Error creating routes output file").Err(err).Send()
	}
	defer orf.Close()

	_, err = orf.WriteString("//  Source: " + routesPath + "\n")
	if err != nil {
		log.Fatal("Error writing source string to routes file").Err(err).Send()
	}

	err = generator.NewRoutes(orf, "lemmy").Generate(routes, impls)
	if err != nil {
		log.Fatal("Error generating output routes file").Err(err).Send()
	}
}
