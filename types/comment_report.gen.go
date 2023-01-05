package types

type CommentReport struct {
	ID                  int           `json:"id" url:"id,omitempty"`
	CreatorID           int           `json:"creator_id" url:"creator_id,omitempty"`
	CommentID           int           `json:"comment_id" url:"comment_id,omitempty"`
	OriginalCommentText string        `json:"original_comment_text" url:"original_comment_text,omitempty"`
	Reason              string        `json:"reason" url:"reason,omitempty"`
	Resolved            bool          `json:"resolved" url:"resolved,omitempty"`
	ResolverID          Optional[int] `json:"resolver_id" url:"resolver_id,omitempty"`
	Published           LemmyTime     `json:"published" url:"published,omitempty"`
	Updated             LemmyTime     `json:"updated" url:"updated,omitempty"`
}
type CommentReportForm struct {
	CreatorID           int    `json:"creator_id" url:"creator_id,omitempty"`
	CommentID           int    `json:"comment_id" url:"comment_id,omitempty"`
	OriginalCommentText string `json:"original_comment_text" url:"original_comment_text,omitempty"`
	Reason              string `json:"reason" url:"reason,omitempty"`
}
