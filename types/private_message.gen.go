package types

import "time"

type PrivateMessage struct {
	ID          int       `json:"id,omitempty" url:"id,omitempty"`
	CreatorID   int       `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	RecipientID int       `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	Content     string    `json:"content,omitempty" url:"content,omitempty"`
	Deleted     bool      `json:"deleted,omitempty" url:"deleted,omitempty"`
	Read        bool      `json:"read,omitempty" url:"read,omitempty"`
	Published   time.Time `json:"published,omitempty" url:"published,omitempty"`
	Updated     time.Time `json:"updated,omitempty" url:"updated,omitempty"`
	ApID        string    `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local       bool      `json:"local,omitempty" url:"local,omitempty"`
}
type PrivateMessageForm struct {
	CreatorID   int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	RecipientID int              `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	Content     string           `json:"content,omitempty" url:"content,omitempty"`
	Deleted     Optional[bool]   `json:"deleted,omitempty" url:"deleted,omitempty"`
	Read        Optional[bool]   `json:"read,omitempty" url:"read,omitempty"`
	Published   time.Time        `json:"published,omitempty" url:"published,omitempty"`
	Updated     time.Time        `json:"updated,omitempty" url:"updated,omitempty"`
	ApID        Optional[string] `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local       Optional[bool]   `json:"local,omitempty" url:"local,omitempty"`
}
