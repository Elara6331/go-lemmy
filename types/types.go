package types

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

type LemmyTime struct {
	time.Time
}

func (lt *LemmyTime) UnmarshalJSON(b []byte) error {
	var timeStr string
	err := json.Unmarshal(b, &timeStr)
	if err != nil {
		return err
	}

	if timeStr == "" {
		lt.Time = time.Unix(0, 0)
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05", timeStr)
	if err != nil {
		return err
	}

	lt.Time = t
	return nil
}

type LemmyWebSocketMsg struct {
	Op   string          `json:"op"`
	Data json.RawMessage `json:"data"`
}

// IsOneOf checks if the message is one of the given operations.
func (msg LemmyWebSocketMsg) IsOneOf(ops ...Operation) bool {
	for _, op := range ops {
		if op.Operation() == msg.Op {
			return true
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

func (u UserOperationCRUD) Operation() string {
	return string(u)
}

func (u UserOperationApub) Operation() string {
	return string(u)
}
