package types

import "time"

type Comment struct {
	ID        int           `json:"id,omitempty" url:"id,omitempty"`
	CreatorID int           `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	PostID    int           `json:"post_id,omitempty" url:"post_id,omitempty"`
	ParentID  Optional[int] `json:"parent_id,omitempty" url:"parent_id,omitempty"`
	Content   string        `json:"content,omitempty" url:"content,omitempty"`
	Removed   bool          `json:"removed,omitempty" url:"removed,omitempty"`
	Read      bool          `json:"read,omitempty" url:"read,omitempty"`
	Published time.Time     `json:"published,omitempty" url:"published,omitempty"`
	Updated   time.Time     `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted   bool          `json:"deleted,omitempty" url:"deleted,omitempty"`
	ApID      string        `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local     bool          `json:"local,omitempty" url:"local,omitempty"`
}
type CommentAlias1 struct {
	ID        int           `json:"id,omitempty" url:"id,omitempty"`
	CreatorID int           `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	PostID    int           `json:"post_id,omitempty" url:"post_id,omitempty"`
	ParentID  Optional[int] `json:"parent_id,omitempty" url:"parent_id,omitempty"`
	Content   string        `json:"content,omitempty" url:"content,omitempty"`
	Removed   bool          `json:"removed,omitempty" url:"removed,omitempty"`
	Read      bool          `json:"read,omitempty" url:"read,omitempty"`
	Published time.Time     `json:"published,omitempty" url:"published,omitempty"`
	Updated   time.Time     `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted   bool          `json:"deleted,omitempty" url:"deleted,omitempty"`
	ApID      string        `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local     bool          `json:"local,omitempty" url:"local,omitempty"`
}
type CommentForm struct {
	CreatorID int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	PostID    int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Content   string           `json:"content,omitempty" url:"content,omitempty"`
	ParentID  Optional[int]    `json:"parent_id,omitempty" url:"parent_id,omitempty"`
	Removed   Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	Read      Optional[bool]   `json:"read,omitempty" url:"read,omitempty"`
	Published time.Time        `json:"published,omitempty" url:"published,omitempty"`
	Updated   time.Time        `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted   Optional[bool]   `json:"deleted,omitempty" url:"deleted,omitempty"`
	ApID      Optional[string] `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local     Optional[bool]   `json:"local,omitempty" url:"local,omitempty"`
}
type CommentLike struct {
	ID        int32     `json:"id,omitempty" url:"id,omitempty"`
	PersonID  int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	CommentID int       `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	PostID    int       `json:"post_id,omitempty" url:"post_id,omitempty"`
	Score     int16     `json:"score,omitempty" url:"score,omitempty"`
	Published time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type CommentLikeForm struct {
	PersonID  int   `json:"person_id,omitempty" url:"person_id,omitempty"`
	CommentID int   `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	PostID    int   `json:"post_id,omitempty" url:"post_id,omitempty"`
	Score     int16 `json:"score,omitempty" url:"score,omitempty"`
}
type CommentSaved struct {
	ID        int32     `json:"id,omitempty" url:"id,omitempty"`
	CommentID int       `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	PersonID  int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	Published time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type CommentSavedForm struct {
	CommentID int `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	PersonID  int `json:"person_id,omitempty" url:"person_id,omitempty"`
}
