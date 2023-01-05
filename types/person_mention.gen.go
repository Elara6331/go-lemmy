package types

import "time"

type PersonMention struct {
	ID          int       `json:"id,omitempty" url:"id,omitempty"`
	RecipientID int       `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	CommentID   int       `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Read        bool      `json:"read,omitempty" url:"read,omitempty"`
	Published   time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type PersonMentionForm struct {
	RecipientID int            `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	CommentID   int            `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Read        Optional[bool] `json:"read,omitempty" url:"read,omitempty"`
}
