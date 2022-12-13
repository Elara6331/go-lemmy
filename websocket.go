package lemmy

import (
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/recws-org/recws"
	"go.arsenm.dev/go-lemmy/types"
)

type WSClient struct {
	conn   *recws.RecConn
	respCh chan types.LemmyWebSocketMsg
	errCh  chan error
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

	ws.Dial(u.String(), nil)

	out := &WSClient{
		conn:   ws,
		respCh: make(chan types.LemmyWebSocketMsg, 10),
		errCh:  make(chan error, 10),
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

func (c *WSClient) Request(op types.UserOperation, data any) error {
	return c.conn.WriteJSON(types.LemmyWebSocketMsg{
		Op:   op,
		Data: data,
	})
}

func (c *WSClient) Responses() <-chan types.LemmyWebSocketMsg {
	return c.respCh
}

func (c *WSClient) Errors() <-chan error {
	return c.errCh
}

func DecodeResponse(data, out any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  out,
	})
	if err != nil {
		return err
	}
	return dec.Decode(data)
}
