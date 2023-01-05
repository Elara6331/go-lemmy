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

// IsOneOf checks if the message is one of the given operation.
// Operations must be of type UserOperation or UserOperationCrud.
func (msg LemmyWebSocketMsg) IsOneOf(ops ...any) bool {
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
