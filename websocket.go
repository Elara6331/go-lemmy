package lemmy

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/websocket"
	"go.elara.ws/go-lemmy/types"
)

type authData struct {
	Auth string `json:"auth"`
}

// WSClient is a client for Lemmy's WebSocket API
type WSClient struct {
	conn       *websocket.Conn
	baseURL    *url.URL
	respCh     chan types.LemmyWebSocketMsg
	errCh      chan error
	recHandler func(c *WSClient)
	Token      string
}

// NewWebSocket creates and returns a new WSClient, and
// starts a goroutine to read server responses and errors
func NewWebSocket(baseURL string) (*WSClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u = u.JoinPath("/api/v3")

	conn, _, err := keepaliveDialer().Dial(u.JoinPath("ws").String(), nil)
	if err != nil {
		return nil, err
	}

	out := &WSClient{
		conn:    conn,
		baseURL: u,
		respCh:  make(chan types.LemmyWebSocketMsg, 10),
		errCh:   make(chan error, 10),
	}

	go func() {
		for {
			var msg types.LemmyWebSocketMsg
			err = conn.ReadJSON(&msg)
			if err != nil {
				out.errCh <- err

				conn.Close()
				err = backoff.RetryNotify(
					func() error {
						conn, _, err = keepaliveDialer().Dial(u.JoinPath("ws").String(), nil)
						if err != nil {
							out.errCh <- err
							return err
						}
						out.conn = conn
						out.recHandler(out)
						return nil
					},
					backoff.NewExponentialBackOff(),
					func(err error, d time.Duration) {
						out.errCh <- fmt.Errorf("reconnect backoff (%s): %w", d, err)
					},
				)
				if err != nil {
					out.errCh <- err
				}

				continue
			}
			out.respCh <- msg
		}
	}()

	return out, nil
}

// ClientLogin logs in to Lemmy by sending an HTTP request to the
// login endpoint. It stores the returned token in the client
// for future use.
func (c *WSClient) ClientLogin(ctx context.Context, l types.Login) error {
	u := &url.URL{}
	*u = *c.baseURL

	if u.Scheme == "ws" {
		u.Scheme = "http"
	} else if u.Scheme == "wss" {
		u.Scheme = "https"
	}

	hc := &Client{baseURL: u, client: http.DefaultClient}
	err := hc.ClientLogin(ctx, l)
	if err != nil {
		return err
	}
	c.Token = hc.Token
	return nil
}

// Request sends a request to the server. If data is nil,
// the authentication token will be sent instead. If data
// has an Auth field, it will be set to the authentication
// token automatically.
func (c *WSClient) Request(op types.Operation, data any) error {
	if data == nil {
		data = authData{}
	}

	data = c.setAuth(data)

	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.conn.WriteJSON(types.LemmyWebSocketMsg{
		Op:   op.Operation(),
		Data: d,
	})
}

// Responses returns a channel that receives messages from
// Lemmy.
func (c *WSClient) Responses() <-chan types.LemmyWebSocketMsg {
	return c.respCh
}

// Errors returns a channel that receives errors
// received while attempting to read responses
func (c *WSClient) Errors() <-chan error {
	return c.errCh
}

func (c *WSClient) OnReconnect(rh func(c *WSClient)) {
	c.recHandler = rh
}

// setAuth uses reflection to automatically
// set struct fields called Auth of type
// string or types.Optional[string] to the
// authentication token, then returns the
// updated struct
func (c *WSClient) setAuth(data any) any {
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

func DecodeResponse(data json.RawMessage, out any) error {
	return json.Unmarshal(data, out)
}

func keepaliveDialer() *websocket.Dialer {
	d := &websocket.Dialer{
		NetDial: func(network, addr string) (net.Conn, error) {
			tcpAddr, err := net.ResolveTCPAddr(network, addr)
			if err != nil {
				return nil, err
			}

			conn, err := net.DialTCP(network, nil, tcpAddr)
			if err != nil {
				return nil, err
			}

			err = conn.SetKeepAlive(true)
			if err != nil {
				return nil, err
			}

			err = conn.SetKeepAlivePeriod(10 * time.Second)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}

	d.NetDialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return d.NetDial(network, addr)
	}

	return d
}
