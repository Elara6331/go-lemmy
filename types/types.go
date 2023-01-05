package types

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LemmyResponse struct {
	Error Optional[string] `json:"error" url:"error,omitempty"`
}

type HTTPError struct {
	Code int
}

func (he HTTPError) Error() string {
	return fmt.Sprintf("%d %s", he.Code, http.StatusText(he.Code))
}

type LemmyError struct {
	ErrStr string
	Code   int
}

func (le LemmyError) Error() string {
	return fmt.Sprintf("%d %s: %s", le.Code, http.StatusText(le.Code), le.ErrStr)
}

type LemmyWebSocketMsg struct {
	Op   string          `json:"op"`
	Data json.RawMessage `json:"data"`
}

// IsOneOf checks if the message is one of the given operations.
func (msg LemmyWebSocketMsg) IsOneOf(ops ...Operation) bool {
	for _, op := range ops {
		switch op := op.(type) {
		case UserOperation:
			if string(op) == msg.Op {
				return true
			}
		case UserOperationCrud:
			if string(op) == msg.Op {
				return true
			}
		}
	}
	return false
}

type Operation interface {
	Operation() string
}

func (u UserOperation) Operation() string {
	return string(u)
}

func (u UserOperationCrud) Operation() string {
	return string(u)
}
