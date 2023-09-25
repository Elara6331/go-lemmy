package types

import (
	"encoding/json"
	"time"
)

// LemmyTime represents a time value returned by the Lemmy server
type LemmyTime struct {
	time.Time
}

// MarshalJSON encodes the Lemmy time to its JSON value
func (lt LemmyTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(lt.Time.Format("2006-01-02T15:04:05"))
}

// UnmarshalJSON decodes JSON into the Lemmy time struct
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
