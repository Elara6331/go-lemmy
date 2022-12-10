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
	"go.arsenm.dev/go-lemmy/types"
)

type Client struct {
	client  *http.Client
	baseURL *url.URL
	token   string
}

func New(baseURL string) (*Client, error) {
	return NewWithClient(baseURL, http.DefaultClient)
}

func NewWithClient(baseURL string, client *http.Client) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u = u.JoinPath("/api/v3")

	return &Client{baseURL: u, client: client}, nil
}

func (c *Client) Login(ctx context.Context, l types.Login) error {
	var lr types.LoginResponse
	res, err := c.req(ctx, http.MethodPost, "/user/login", l, &lr)
	if err != nil {
		return err
	}

	err = resError(res, lr.LemmyResponse)
	if err != nil {
		return err
	}

	c.token = lr.JWT.MustValue()

	return nil
}

func (c *Client) Token() string {
	return c.token
}

func (c *Client) SetToken(t string) {
	c.token = t
}

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

func resError(res *http.Response, lr types.LemmyResponse) error {
	if lr.Error.IsValid() {
		return types.LemmyError{
			Code:   res.StatusCode,
			ErrStr: lr.Error.MustValue(),
		}
	} else if res.StatusCode > 299 {
		return types.HTTPError{
			Code: res.StatusCode,
		}
	} else {
		return nil
	}
}

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

	authType := authField.Kind()
	if authType != reflect.String {
		return data
	}

	authField.SetString(c.token)

	return val.Elem().Interface()
}
