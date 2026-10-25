package lemmy

import (
	"context"
	"errors"
)

// Token is a Lemmy JWT token
type Token string

func (t Token) BearerAuth(ctx context.Context, _ OperationName) (BearerAuth, error) {
	return BearerAuth{Token: string(t)}, nil
}

func (c *Client) Token() (Token, error) {
	if tok, ok := c.sec.(Token); ok {
		return tok, nil
	} else {
		return "", errors.New("client security source is not a token")
	}
}

// NewWithLogin creates a new client and logs in to the Lemmy instance
func NewWithLogin(srvURL string, login *Login, opts ...ClientOption) (*Client, error) {
	c, err := NewClient(srvURL, nil, opts...)
	if err != nil {
		return c, err
	}
	
	res, err := c.Login(context.Background(), login)
	if err != nil {
		return c, err
	}
	
	if res.Jwt.IsSet() {
		c.sec = Token(res.Jwt.Value)
	}
	
	return c, nil
}