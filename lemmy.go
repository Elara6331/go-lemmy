package lemmy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
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

// ClientLogin logs in to Lemmy by calling the Lemmy endpoint,
// and stores the returned token for use in future requests.
func (c *Client) ClientLogin(ctx context.Context, l Login) error {
	lr, err := c.Login(ctx, l)
	if err != nil {
		return err
	}
	c.Token = lr.JWT.ValueOrEmpty()
	return nil
}

// req makes a request to the server
func (c *Client) req(ctx context.Context, method string, path string, data, resp any) (*http.Response, error) {
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

	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
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

// getReq makes a get request to the Lemmy server.
// It's separate from req() because it uses query
// parameters rather than a JSON request body.
func (c *Client) getReq(ctx context.Context, method string, path string, data, resp any) (*http.Response, error) {
	getURL := c.baseURL.JoinPath(path)
	if data != nil {
		vals, err := query.Values(data)
		if err != nil {
			return nil, err
		}
		getURL.RawQuery = vals.Encode()
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		getURL.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
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

// resError returns an error if the the response contains an error
func resError(res *http.Response, lr LemmyResponse) error {
	if lr.Error.IsValid() {
		return LemmyError{
			Code:   res.StatusCode,
			ErrStr: lr.Error.MustValue(),
		}
	} else if res.StatusCode != http.StatusOK {
		return HTTPError{
			Code: res.StatusCode,
		}
	} else {
		return nil
	}
}
