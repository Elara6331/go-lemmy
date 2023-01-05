package generator

import (
	"io"
	"strings"

	"go.arsenm.dev/go-lemmy/cmd/gen/parser"
	"github.com/dave/jennifer/jen"
)

type StructGenerator struct {
	w       io.Writer
	PkgName string
}

func NewStruct(w io.Writer, pkgName string) *StructGenerator {
	return &StructGenerator{w, pkgName}
}

func (s *StructGenerator) Generate(items []parser.Item) error {
	f := jen.NewFile(s.PkgName)
	for _, item := range items {
		if item.Struct != nil {
			st := item.Struct
			f.Type().Id(st.Name).StructFunc(func(g *jen.Group) {
				for _, field := range st.Fields {
					var t jen.Code
					if field.Type == "time.Time" {
						t = jen.Qual("time", "Time")
					} else {
						t = jen.Id(field.Type)
					}

					g.Id(field.Name).Add(t).Tag(map[string]string{
						"json": field.OrigName + ",omitempty",
						"url":  field.OrigName + ",omitempty",
					})
				}

				if strings.HasSuffix(st.Name, "Response") {
					g.Id("LemmyResponse")
				}
			})
		} else if item.Enum != nil {
			e := item.Enum
			f.Type().Id(e.Name).String()

			f.Const().DefsFunc(func(g *jen.Group) {
				for _, member := range e.Members {
					g.Id(e.Name + string(member)).Op("=").Lit(string(member))
				}
			})
		}
	}

	return f.Render(s.w)
}
