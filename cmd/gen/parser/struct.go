package parser

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

var (
	structRegex = regexp.MustCompile(`pub struct (.+) \{`)
	fieldRegex  = regexp.MustCompile(`(?U) {1,1}([^ ]+): (.+),`)

	enumRegex   = regexp.MustCompile(`pub enum (.+) \{`)
	memberRegex = regexp.MustCompile(`  ([^ #]+),\n`)
)

type Item struct {
	Struct *Struct
	Enum   *Enum
}

type Struct struct {
	Name   string
	Fields []Field
}

type Field struct {
	OrigName string
	Name     string
	Type     string
}

type Enum struct {
	Name    string
	Members []Member
}

type Member string

type StructParser struct {
	r             *bufio.Reader
	Skip          []string
	TransformName func(string) string
	TransformType func(string) string
}

func NewStruct(r io.Reader) *StructParser {
	return &StructParser{
		r:             bufio.NewReader(r),
		TransformName: TransformNameGo,
		TransformType: TransformTypeGo,
	}
}

func (s *StructParser) Parse() ([]Item, error) {
	var out []Item
	for {
		line, err := s.r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		if structRegex.MatchString(line) {
			structName := structRegex.FindStringSubmatch(line)[1]
			if slices.Contains(s.Skip, structName) {
				continue
			}

			// If the line ends with "}", this is a struct with no fields
			if strings.HasSuffix(line, "}\n") {
				out = append(out, Item{
					Struct: &Struct{
						Name: structRegex.FindStringSubmatch(line)[1],
					},
				})
				continue
			}

			fields, err := s.parseStructFields()
			if err != nil {
				return nil, err
			}

			out = append(out, Item{
				Struct: &Struct{
					Name:   structName,
					Fields: fields,
				},
			})
		} else if enumRegex.MatchString(line) {
			enumName := enumRegex.FindStringSubmatch(line)[1]
			if slices.Contains(s.Skip, enumName) {
				continue

			}
			members, err := s.parseEnumMemebers()
			if err != nil {
				return nil, err
			}

			out = append(out, Item{
				Enum: &Enum{
					Name:    enumName,
					Members: members,
				},
			})
		}
	}
	return out, nil
}

func (s *StructParser) parseStructFields() ([]Field, error) {
	var out []Field
	for {
		line, err := s.r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			if strings.HasPrefix(line, "}") {
				return out, nil
			} else {
				return nil, io.ErrUnexpectedEOF
			}
		} else if err != nil {
			return nil, err
		}

		if strings.HasPrefix(line, "}") {
			return out, nil
		} else if strings.HasPrefix(line, "//") {
			continue
		} else if !fieldRegex.MatchString(line) {
			continue
		}

		sm := fieldRegex.FindStringSubmatch(line)
		if sm[1] == "Example" {
			continue
		}

		out = append(out, Field{
			OrigName: sm[1],
			Name:     s.TransformName(sm[1]),
			Type:     s.TransformType(sm[2]),
		})
	}
}

func (s *StructParser) parseEnumMemebers() ([]Member, error) {
	var out []Member
	for {
		line, err := s.r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			if strings.HasPrefix(line, "}") {
				return out, nil
			} else {
				return nil, io.ErrUnexpectedEOF
			}
		} else if err != nil {
			return nil, err
		}

		if strings.HasPrefix(line, "}") {
			return out, nil
		} else if strings.HasPrefix(line, "//") {
			continue
		} else if !memberRegex.MatchString(line) {
			continue
		}

		sm := memberRegex.FindStringSubmatch(line)

		out = append(out, Member(sm[1]))
	}
}

// TransformTypeGo transforms Rust types to Go
//
//	Example: TransformTypeGo("Option<Vec<i64>>") // returns "Optional[[]int64]"
func TransformTypeGo(t string) string {
	prefix := ""
	suffix := ""

	for strings.HasPrefix(t, "Option<") {
		t = strings.TrimPrefix(strings.TrimSuffix(t, ">"), "Option<")
		prefix += "Optional["
		suffix += "]"
	}

	for strings.HasPrefix(t, "Vec<") {
		t = strings.TrimPrefix(strings.TrimSuffix(t, ">"), "Vec<")
		prefix += "[]"
	}

	for strings.HasPrefix(t, "Sensitive<") {
		t = strings.TrimPrefix(strings.TrimSuffix(t, ">"), "Sensitive<")
	}

	if strings.HasSuffix(t, "Id") {
		t = "int"
	}

	switch t {
	case "String", "Url", "DbUrl", "Ltree":
		t = "string"
	case "usize":
		t = "uint"
	case "i64":
		t = "int64"
	case "i32":
		t = "int32"
	case "i16":
		t = "int16"
	case "i8":
		t = "int8"
	case "chrono::NaiveDateTime":
		return "LemmyTime"
	case "Value":
		return "any"
	}

	return prefix + t + suffix
}

// TransformNameGo transforms conventional Rust naming to
// conventional Go naming.
//
//	Example: TransformNameGo("post_id") // returns "PostID"
func TransformNameGo(s string) string {
	out := ""

	splitName := strings.Split(s, "_")
	for _, segment := range splitName {
		switch segment {
		case "id":
			out += "ID"
		case "url":
			out += "URL"
		case "nsfw":
			out += "NSFW"
		default:
			if len(segment) == 0 {
				continue
			}

			out += strings.ToUpper(segment[:1]) + segment[1:]
		}
	}

	return out
}

func (s *StructParser) Reset(r io.Reader) {
	s.r.Reset(r)
}
