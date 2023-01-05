package types

import "time"

type PersonBlock struct {
	ID        int       `json:"id,omitempty" url:"id,omitempty"`
	PersonID  int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	TargetID  int       `json:"target_id,omitempty" url:"target_id,omitempty"`
	Published time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type PersonBlockForm struct {
	PersonID int `json:"person_id,omitempty" url:"person_id,omitempty"`
	TargetID int `json:"target_id,omitempty" url:"target_id,omitempty"`
}
