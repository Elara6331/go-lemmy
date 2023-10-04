package main

import (
	"flag"
	"os"
	"path/filepath"

	"go.elara.ws/go-lemmy/cmd/gen/extractor"
	"go.elara.ws/go-lemmy/cmd/gen/generator"
	"go.elara.ws/logger"
	"go.elara.ws/logger/log"
)

func init() {
	log.Logger = logger.NewPretty(os.Stderr)
}

func main() {
	jsonFile := flag.String("json-file", "lemmy.json", "Path to the JSON file generated from lemmy's docs")
	outDir := flag.String("out-dir", "out", "Directory to write output in")
	flag.Parse()

	e, err := extractor.New(*jsonFile)
	if err != nil {
		log.Fatal("Error creating extractor").Err(err).Send()
	}

	routes, structs := e.Extract()

	otf, err := os.Create(filepath.Join(*outDir, "types.gen.go"))
	if err != nil {
		log.Fatal("Error creating types output file").Err(err).Send()
	}
	defer otf.Close()

	err = generator.NewStruct(otf, "lemmy").Generate(structs)
	if err != nil {
		log.Fatal("Error generating output routes file").Err(err).Send()
	}

	orf, err := os.Create(filepath.Join(*outDir, "routes.gen.go"))
	if err != nil {
		log.Fatal("Error creating routes output file").Err(err).Send()
	}
	defer orf.Close()

	err = generator.NewRoutes(orf, "lemmy").Generate(routes)
	if err != nil {
		log.Fatal("Error generating output routes file").Err(err).Send()
	}
}
