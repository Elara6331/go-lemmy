package lemmy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
	"go.elara.ws/go-lemmy/types"
)

// Client is a client for Lemmy's HTTP API
type Client struct {
	client  *http.Client
	baseURL *url.URL
	Token   string
}

// New creates a new Lemmy client with the default HTTP client.
func New(baseURL string) (*Client, error) {
	return NewWithClient(baseURL, http.DefaultClient)
}

// NewWithClient creates a new Lemmy client with the given HTTP client
func NewWithClient(baseURL string, client *http.Client) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u = u.JoinPath("/api/v3")

	return &Client{baseURL: u, client: client}, nil
}

// ClientLogin logs in to Lemmy by sending an HTTP request to the
// login endpoint. It stores the returned token in the client
// for future use.
func (c *Client) ClientLogin(ctx context.Context, l types.Login) error {
	lr, err := c.Login(ctx, l)
	if err != nil {
		return err
	}

	c.Token = lr.JWT.MustValue()
	return nil
}

// req makes a request to the server
func (c *Client) req(ctx context.Context, method string, path string, data, resp any) (*http.Response, error) {
	data = c.setAuth(data)

	var r io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		r = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL.JoinPath(path).String(),
		r,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(resp)
	if err == io.EOF {
		return res, nil
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

// getReq makes a get request to the Lemmy server.
// It is separate from req() because it uses query
// parameters rather than a JSON request body.
func (c *Client) getReq(ctx context.Context, method string, path string, data, resp any) (*http.Response, error) {
	data = c.setAuth(data)

	getURL := c.baseURL.JoinPath(path)
	vals, err := query.Values(data)
	if err != nil {
		return nil, err
	}
	getURL.RawQuery = vals.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		getURL.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(resp)
	if err == io.EOF {
		return res, nil
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

// resError returns an error if the given response is an error
func resError(res *http.Response, lr types.LemmyResponse) error {
	if lr.Error.IsValid() {
		return types.LemmyError{
			Code:   res.StatusCode,
			ErrStr: lr.Error.MustValue(),
		}
	} else if res.StatusCode != http.StatusOK {
		return types.HTTPError{
			Code: res.StatusCode,
		}
	} else {
		return nil
	}
}

// setAuth uses reflection to automatically
// set struct fields called Auth of type
// string or types.Optional[string] to the
// authentication token, then returns the
// updated struct
func (c *Client) setAuth(data any) any {
	if data == nil {
		return data
	}

	val := reflect.New(reflect.TypeOf(data))
	val.Elem().Set(reflect.ValueOf(data))

	authField := val.Elem().FieldByName("Auth")
	if !authField.IsValid() {
		return data
	}

	switch authField.Type().String() {
	case "string":
		authField.SetString(c.Token)
	case "types.Optional[string]":
		setMtd := authField.MethodByName("Set")
		out := setMtd.Call([]reflect.Value{reflect.ValueOf(c.Token)})
		authField.Set(out[0])
	default:
		return data
	}

	return val.Elem().Interface()
}
