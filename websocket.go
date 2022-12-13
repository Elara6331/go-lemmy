package lemmy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/recws-org/recws"
	"go.arsenm.dev/go-lemmy/types"
)

type authData struct {
	Auth string `json:"auth"`
}

type WSClient struct {
	conn    *recws.RecConn
	baseURL *url.URL
	respCh  chan types.LemmyWebSocketMsg
	errCh   chan error
	token   string
}

func NewWebSocket(baseURL string) (*WSClient, error) {
	ws := &recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u = u.JoinPath("/api/v3")

	ws.Dial(u.JoinPath("ws").String(), nil)

	out := &WSClient{
		conn:    ws,
		baseURL: u,
		respCh:  make(chan types.LemmyWebSocketMsg, 10),
		errCh:   make(chan error, 10),
	}

	go func() {
		for {
			var msg types.LemmyWebSocketMsg
			err = ws.ReadJSON(&msg)
			if err != nil {
				out.errCh <- err
				continue
			}
			out.respCh <- msg
		}
	}()

	return out, nil
}

// SetConnectHandler sets the connection handler function
func (c *WSClient) SetConnectHandler(f func() error) {
	c.conn.SubscribeHandler = f
}

// ConnectHandler invokes the connection handler function
func (c *WSClient) ConnectHandler() error {
	if c.conn.SubscribeHandler == nil {
		return nil
	}
	return c.conn.SubscribeHandler()
}

func (c *WSClient) Login(ctx context.Context, l types.Login) error {
	u := &url.URL{}
	*u = *c.baseURL

	if u.Scheme == "ws" {
		u.Scheme = "http"
	} else if u.Scheme == "wss" {
		u.Scheme = "https"
	}

	hc := &Client{baseURL: u, client: http.DefaultClient}
	err := hc.Login(ctx, l)
	if err != nil {
		return err
	}
	c.token = hc.token
	return nil
}

func (c *WSClient) Request(op types.UserOperation, data any) error {
	if data == nil {
		data = authData{}
	}

	data = c.setAuth(data)

	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.conn.WriteJSON(types.LemmyWebSocketMsg{
		Op:   op,
		Data: d,
	})
}

func (c *WSClient) Responses() <-chan types.LemmyWebSocketMsg {
	return c.respCh
}

func (c *WSClient) Errors() <-chan error {
	return c.errCh
}

func (c *WSClient) setAuth(data any) any {
	val := reflect.New(reflect.TypeOf(data))
	val.Elem().Set(reflect.ValueOf(data))

	authField := val.Elem().FieldByName("Auth")
	if !authField.IsValid() {
		return data
	}

	switch authField.Type().String() {
	case "string":
		authField.SetString(c.token)
	case "types.Optional[string]":
		setMtd := authField.MethodByName("Set")
		out := setMtd.Call([]reflect.Value{reflect.ValueOf(c.token)})
		authField.Set(out[0])
	default:
		return data
	}

	return val.Elem().Interface()
}

func DecodeResponse(data json.RawMessage, out any) error {
	return json.Unmarshal(data, out)
}
