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
	ReturnName string
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

// New parses the file at path and returns an extractor with its contents.
func New(path string) (*Extractor, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &Extractor{gjson.ParseBytes(data)}, nil
}

// Extract reads the JSON document and extracts all the routes and structs from it.
func (e *Extractor) Extract() ([]Route, []Struct) {
	structs := map[int64]Struct{}
	var out []Route

	// Get all the routes in the JSON document
	routes := e.root.Get("children.#.children.#(kind==2048)#|@flatten")

	for _, route := range routes.Array() {
		name := route.Get("name").String()
		signature := route.Get(`signatures.0`)

		// Get the code part of the route's summary.
		// This will contain the HTTP method and path.
		httpInfo := signature.Get(`comment.summary.#(kind=="code").text`).String()
		if !strings.HasPrefix(httpInfo, "`HTTP") {
			continue
		}
		method, path := parseHTTPInfo(httpInfo)

		summary := strings.TrimSpace(signature.Get(`comment.summary.#(kind=="text").text`).String())
		if summary == "" {
			continue
		}

		// Get the ID and name of the type this function accepts
		paramsID := signature.Get("parameters.0.type.target").Int()
		paramsName := signature.Get("parameters.0.type.name").String()

		// Get the ID and name of the type this function returns
		returnID := signature.Get("type.typeArguments.0.target").Int()
		returnName := signature.Get("type.typeArguments.0.name").String()

		anyType := false
		if returnName == "any" {
			anyType = true
		}

		// Get the referenced structs from the JSON document
		e.getStructs([]int64{paramsID, returnID}, structs)

		// If the parameters struct contains no fields or union names
		if len(structs[paramsID].Fields) == 0 && len(structs[paramsID].UnionNames) == 0 {
			// Delete the params struct from the structs map
			// to make sure it doesn't get generated
			delete(structs, paramsID)

			// Set paramsName to an empty string to signify that this route
			// has no input parameters.
			paramsName = ""
		}

		// If the return struct contains no fields or union names
		if len(structs[returnID].Fields) == 0 && len(structs[returnID].UnionNames) == 0 {
			// Delete the return struct from the structs map
			// to make sure it doesn't get generated
			delete(structs, returnID)

			// Set paramsName to an empty string to signify that this route
			// has no return value.
			returnName = ""
		}

		if anyType {
			returnName = "map[string]any"
		}

		out = append(out, Route{
			Name:       name,
			Summary:    summary,
			Method:     method,
			Path:       path,
			ParamsName: paramsName,
			ReturnName: returnName,
		})
	}
	return out, getStructSlice(structs)
}

func (e *Extractor) getStructs(ids []int64, structs map[int64]Struct) {
	for _, id := range ids {
		if _, ok := structs[id]; ok {
			continue
		}

		// Get the struct with the given ID from the JSON document
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

			// Recursively get any structs referenced by this one
			e.getStructs(newIDs, structs)
		}

	}
}

// unionNames gets all the names of union type members
func (e *Extractor) unionNames(jstruct gjson.Result) []string {
	jnames := jstruct.Get("type.types").Array()
	out := make([]string, len(jnames))
	for i, name := range jnames {
		out[i] = name.Get("value").String()
	}
	return out
}

// fields gets all the fields in a given struct from the JSON document.
// It returns the fields and the IDs of any types they referenced.
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
				// If this field is referencing another type, add that type's id
				// to the ids slice.
				ids = append(ids, jfield.Get("type.elementType.target").Int())
			case "union":
				// Convert unions to strings
				field.Type = "string"
			}
		} else {
			field.Type = jfield.Get("type.name").String()

			switch jfield.Get("type.type").String() {
			case "reference":
				// If this field is referencing another type, add that type's id
				// to the ids slice.
				ids = append(ids, jfield.Get("type.target").Int())
			case "union":
				// Convert unions to strings
				field.Type = "string"
			}
		}

		fields = append(fields, field)
	}
	return fields, ids
}

// parseHTTPInfo parses the string from a route's summary,
// and returns the method and path it uses
func parseHTTPInfo(httpInfo string) (method, path string) {
	httpInfo = strings.Trim(httpInfo, "`")
	method, path, _ = strings.Cut(httpInfo, " ")
	method = strings.TrimPrefix(method, "HTTP.")
	method = strings.ToUpper(method)
	return method, path
}

// getStructSlice returns all the structs in a map
func getStructSlice(m map[int64]Struct) []Struct {
	out := make([]Struct, len(m))
	i := 0
	for _, s := range m {
		out[i] = s
		i++
	}
	return out
}
