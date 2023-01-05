package types

type PrivateMessage struct {
	ID          int       `json:"id" url:"id,omitempty"`
	CreatorID   int       `json:"creator_id" url:"creator_id,omitempty"`
	RecipientID int       `json:"recipient_id" url:"recipient_id,omitempty"`
	Content     string    `json:"content" url:"content,omitempty"`
	Deleted     bool      `json:"deleted" url:"deleted,omitempty"`
	Read        bool      `json:"read" url:"read,omitempty"`
	Published   LemmyTime `json:"published" url:"published,omitempty"`
	Updated     LemmyTime `json:"updated" url:"updated,omitempty"`
	ApID        string    `json:"ap_id" url:"ap_id,omitempty"`
	Local       bool      `json:"local" url:"local,omitempty"`
}
type PrivateMessageForm struct {
	CreatorID   int              `json:"creator_id" url:"creator_id,omitempty"`
	RecipientID int              `json:"recipient_id" url:"recipient_id,omitempty"`
	Content     string           `json:"content" url:"content,omitempty"`
	Deleted     Optional[bool]   `json:"deleted" url:"deleted,omitempty"`
	Read        Optional[bool]   `json:"read" url:"read,omitempty"`
	Published   LemmyTime        `json:"published" url:"published,omitempty"`
	Updated     LemmyTime        `json:"updated" url:"updated,omitempty"`
	ApID        Optional[string] `json:"ap_id" url:"ap_id,omitempty"`
	Local       Optional[bool]   `json:"local" url:"local,omitempty"`
}
