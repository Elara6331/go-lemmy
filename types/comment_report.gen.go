package types

import "time"

type CommentReport struct {
	ID                  int           `json:"id,omitempty" url:"id,omitempty"`
	CreatorID           int           `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	CommentID           int           `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	OriginalCommentText string        `json:"original_comment_text,omitempty" url:"original_comment_text,omitempty"`
	Reason              string        `json:"reason,omitempty" url:"reason,omitempty"`
	Resolved            bool          `json:"resolved,omitempty" url:"resolved,omitempty"`
	ResolverID          Optional[int] `json:"resolver_id,omitempty" url:"resolver_id,omitempty"`
	Published           time.Time     `json:"published,omitempty" url:"published,omitempty"`
	Updated             time.Time     `json:"updated,omitempty" url:"updated,omitempty"`
}
type CommentReportForm struct {
	CreatorID           int    `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	CommentID           int    `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	OriginalCommentText string `json:"original_comment_text,omitempty" url:"original_comment_text,omitempty"`
	Reason              string `json:"reason,omitempty" url:"reason,omitempty"`
}
