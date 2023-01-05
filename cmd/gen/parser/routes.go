package parser

import (
	"bufio"
	"errors"
	"io"
	"net/url"
	"regexp"
	"strings"
)

var (
	scopeRegex = regexp.MustCompile(`web::(?:scope|resource)\("(.*)"\)\n`)
	routeRegex = regexp.MustCompile(`\.route\(\n?\s*(?:"(.*)",[ \n])?\s*web::(.+)\(\)\.to\(route_.*::<(.+)>`)
)

type Route struct {
	Path   string
	Method string
	Struct string
}

type RoutesParser struct {
	r *bufio.Reader
}

func NewRoutes(r io.Reader) *RoutesParser {
	return &RoutesParser{
		r: bufio.NewReader(r),
	}
}

func (r *RoutesParser) Parse() ([]Route, error) {
	var out []Route
	for {
		line, err := r.r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		if scopeRegex.MatchString(line) {
			scopePath := scopeRegex.FindStringSubmatch(line)[1]
			if scopePath == "/api/v3" {
				continue
			}

			routes, err := r.parseRoutes()
			if err != nil {
				return nil, err
			}

			for i := range routes {
				path, err := url.JoinPath(scopePath, routes[i].Path)
				if err != nil {
					return nil, err
				}
				routes[i].Path = path
			}

			out = append(out, routes...)
		}
	}
	return out, nil
}

func (r *RoutesParser) parseRoutes() ([]Route, error) {
	var out []Route
	for {
		line, err := r.r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			if strings.TrimSpace(line)[:1] == ")" {
				return out, nil
			} else {
				return nil, io.ErrUnexpectedEOF
			}
		} else if err != nil {
			return nil, err
		}

		if strings.TrimSpace(line) == ".route(" {
			lines, err := r.readLines(3)
			if err != nil {
				return nil, err
			}
			line += lines
		}

		if strings.TrimSpace(line)[:1] == ")" {
			return out, nil
		} else if strings.HasPrefix(line, "//") {
			continue
		} else if !routeRegex.MatchString(line) {
			continue
		}

		sm := routeRegex.FindStringSubmatch(line)
		out = append(out, Route{
			Path:   sm[1],
			Method: strings.ToUpper(sm[2]),
			Struct: sm[3],
		})
	}
}

func (r *RoutesParser) readLines(n int) (string, error) {
	out := ""
	for i := 0; i < n; i++ {
		line, err := r.r.ReadString('\n')
		if err != nil {
			return "", err
		}
		out += line
	}
	return out, nil
}

func (r *RoutesParser) Reset(rd io.Reader) {
	r.r.Reset(rd)
}
