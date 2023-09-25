package extractor

import (
	"fmt"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

type Route struct {
	Name       string
	Summary    string
	Method     string
	Path       string
	ParamsName string
	ParamsID   int64
	ReturnName string
	ReturnID   int64
}

type Struct struct {
	Name       string
	Fields     []Field
	UnionNames []string
}

type Field struct {
	Name       string
	IsArray    bool
	IsOptional bool
	Type       string
}

type Extractor struct {
	root gjson.Result
}

func New(path string) (*Extractor, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &Extractor{gjson.ParseBytes(data)}, nil
}

func (e *Extractor) Routes() []Route {
	var out []Route
	routes := e.root.Get("children.#.children.#(kind==2048)#|@flatten")
	for _, route := range routes.Array() {
		name := route.Get("name").String()
		signature := route.Get(`signatures.0`)

		httpInfo := signature.Get(`comment.summary.#(kind=="code").text`).String()
		if !strings.HasPrefix(httpInfo, "`HTTP") {
			continue
		}
		method, path := parseHTTPInfo(httpInfo)

		summary := strings.TrimSpace(signature.Get(`comment.summary.#(kind=="text").text`).String())
		if summary == "" {
			continue
		}

		paramsID := signature.Get("parameters.0.type.target").Int()
		paramsName := signature.Get("parameters.0.type.name").String()
		returnID := signature.Get("type.typeArguments.0.target").Int()
		returnName := signature.Get("type.typeArguments.0.name").String()

		out = append(out, Route{
			Name:       name,
			Summary:    summary,
			Method:     method,
			Path:       path,
			ParamsName: paramsName,
			ParamsID:   paramsID,
			ReturnName: returnName,
			ReturnID:   returnID,
		})
	}
	return out
}

func (e *Extractor) Structs(routes []Route) []Struct {
	var ids []int64
	for _, route := range routes {
		ids = append(ids, route.ParamsID)
		if route.ReturnID != 0 {
			ids = append(ids, route.ReturnID)
		}
	}

	structs := map[int64]Struct{}
	e.getStructs(ids, structs)
	return getKeys(structs)
}

func (e *Extractor) getStructs(ids []int64, structs map[int64]Struct) {
	for _, id := range ids {
		if _, ok := structs[id]; ok {
			continue
		}

		jstruct := e.root.Get(fmt.Sprintf("children.#(id==%d)", id))
		if !jstruct.Exists() {
			continue
		}

		name := jstruct.Get("name").String()

		if jstruct.Get("type.type").String() == "union" {
			structs[id] = Struct{
				Name:       name,
				UnionNames: e.unionNames(jstruct),
			}
		} else {
			fields, newIDs := e.fields(jstruct)

			structs[id] = Struct{
				Name:   name,
				Fields: fields,
			}

			e.getStructs(newIDs, structs)
		}

	}
}

func (e *Extractor) unionNames(jstruct gjson.Result) []string {
	jnames := jstruct.Get("type.types").Array()
	out := make([]string, len(jnames))
	for i, name := range jnames {
		out[i] = name.Get("value").String()
	}
	return out
}

func (e *Extractor) fields(jstruct gjson.Result) ([]Field, []int64) {
	var fields []Field
	var ids []int64
	jfields := jstruct.Get("children").Array()
	for _, jfield := range jfields {
		var field Field

		field.Name = jfield.Get("name").String()
		field.IsOptional = jfield.Get("flags.isOptional").Bool()

		if jfield.Get("type.type").String() == "array" {
			field.IsArray = true
			field.Type = jfield.Get("type.elementType.name").String()

			switch jfield.Get("type.elementType.type").String() {
			case "reference":
				ids = append(ids, jfield.Get("type.elementType.target").Int())
			case "union":
				field.Type = "string"
			}
		} else {
			field.Type = jfield.Get("type.name").String()

			switch jfield.Get("type.type").String() {
			case "reference":
				ids = append(ids, jfield.Get("type.target").Int())
			case "union":
				field.Type = "string"
			}
		}

		fields = append(fields, field)
	}
	return fields, ids
}

func parseHTTPInfo(httpInfo string) (method, path string) {
	httpInfo = strings.Trim(httpInfo, "`")
	method, path, _ = strings.Cut(httpInfo, " ")
	method = strings.TrimPrefix(method, "HTTP.")
	method = strings.ToUpper(method)
	return method, path
}

func getKeys(m map[int64]Struct) []Struct {
	out := make([]Struct, len(m))
	i := 0
	for _, s := range m {
		out[i] = s
		i++
	}
	return out
}
