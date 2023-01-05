package generator

import (
	"io"
	"strings"

	"github.com/dave/jennifer/jen"
	"go.arsenm.dev/go-lemmy/cmd/gen/parser"
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
					g.Id(field.Name).Id(field.Type).Tag(map[string]string{
						"json": field.OrigName,
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
					g.Id(e.Name + string(member)).Id(e.Name).Op("=").Lit(string(member))
				}
			})
		}
	}

	return f.Render(s.w)
}
