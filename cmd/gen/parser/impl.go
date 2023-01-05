package parser

import (
	"bufio"
	"errors"
	"io"
	"regexp"
)

var (
	implRegex     = regexp.MustCompile(`impl Perform.* for (.+) {`)
	respTypeRegex = regexp.MustCompile(`type Response = (.+);`)
)

var ErrNoType = errors.New("type line not found")

type ImplParser struct {
	r *bufio.Reader
}

func NewImpl(r io.Reader) *ImplParser {
	return &ImplParser{
		r: bufio.NewReader(r),
	}
}

func (i *ImplParser) Parse() (map[string]string, error) {
	out := map[string]string{}
	for {
		line, err := i.r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		if implRegex.MatchString(line) {
			im := implRegex.FindStringSubmatch(line)

			line, err := i.r.ReadString('\n')
			if errors.Is(err, io.EOF) {
				return nil, io.ErrUnexpectedEOF
			} else if err != nil {
				return nil, err
			}

			if !respTypeRegex.MatchString(line) {
				return nil, ErrNoType
			}

			rtm := respTypeRegex.FindStringSubmatch(line)

			out[im[1]] = rtm[1]
		}
	}

	return out, nil
}

func (i *ImplParser) Reset(r io.Reader) {
	i.r.Reset(r)
}
