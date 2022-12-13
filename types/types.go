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
	Op   UserOperation   `json:"op"`
	Data json.RawMessage `json:"data"`
}
