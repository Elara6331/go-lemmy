package types

import "time"

type Activity struct {
	ID        int32          `json:"id,omitempty" url:"id,omitempty"`
	Data      any            `json:"data,omitempty" url:"data,omitempty"`
	Local     bool           `json:"local,omitempty" url:"local,omitempty"`
	Published time.Time      `json:"published,omitempty" url:"published,omitempty"`
	Updated   time.Time      `json:"updated,omitempty" url:"updated,omitempty"`
	ApID      string         `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Sensitive Optional[bool] `json:"sensitive,omitempty" url:"sensitive,omitempty"`
}
type ActivityForm struct {
	Data      any            `json:"data,omitempty" url:"data,omitempty"`
	Local     Optional[bool] `json:"local,omitempty" url:"local,omitempty"`
	Updated   time.Time      `json:"updated,omitempty" url:"updated,omitempty"`
	ApID      string         `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Sensitive bool           `json:"sensitive,omitempty" url:"sensitive,omitempty"`
}
