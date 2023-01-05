package types

import "time"

type Post struct {
	ID               int              `json:"id,omitempty" url:"id,omitempty"`
	Name             string           `json:"name,omitempty" url:"name,omitempty"`
	URL              Optional[string] `json:"url,omitempty" url:"url,omitempty"`
	Body             Optional[string] `json:"body,omitempty" url:"body,omitempty"`
	CreatorID        int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	CommunityID      int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed          bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Locked           bool             `json:"locked,omitempty" url:"locked,omitempty"`
	Published        time.Time        `json:"published,omitempty" url:"published,omitempty"`
	Updated          time.Time        `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted          bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	NSFW             bool             `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	Stickied         bool             `json:"stickied,omitempty" url:"stickied,omitempty"`
	EmbedTitle       Optional[string] `json:"embed_title,omitempty" url:"embed_title,omitempty"`
	EmbedDescription Optional[string] `json:"embed_description,omitempty" url:"embed_description,omitempty"`
	EmbedHtml        Optional[string] `json:"embed_html,omitempty" url:"embed_html,omitempty"`
	ThumbnailURL     Optional[string] `json:"thumbnail_url,omitempty" url:"thumbnail_url,omitempty"`
	ApID             string           `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local            bool             `json:"local,omitempty" url:"local,omitempty"`
}
type PostForm struct {
	Name             string           `json:"name,omitempty" url:"name,omitempty"`
	CreatorID        int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	CommunityID      int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	NSFW             Optional[bool]   `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	URL              Optional[string] `json:"url,omitempty" url:"url,omitempty"`
	Body             Optional[string] `json:"body,omitempty" url:"body,omitempty"`
	Removed          Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	Locked           Optional[bool]   `json:"locked,omitempty" url:"locked,omitempty"`
	Published        time.Time        `json:"published,omitempty" url:"published,omitempty"`
	Updated          time.Time        `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted          Optional[bool]   `json:"deleted,omitempty" url:"deleted,omitempty"`
	Stickied         Optional[bool]   `json:"stickied,omitempty" url:"stickied,omitempty"`
	EmbedTitle       Optional[string] `json:"embed_title,omitempty" url:"embed_title,omitempty"`
	EmbedDescription Optional[string] `json:"embed_description,omitempty" url:"embed_description,omitempty"`
	EmbedHtml        Optional[string] `json:"embed_html,omitempty" url:"embed_html,omitempty"`
	ThumbnailURL     Optional[string] `json:"thumbnail_url,omitempty" url:"thumbnail_url,omitempty"`
	ApID             Optional[string] `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local            Optional[bool]   `json:"local,omitempty" url:"local,omitempty"`
}
type PostLike struct {
	ID        int32     `json:"id,omitempty" url:"id,omitempty"`
	PostID    int       `json:"post_id,omitempty" url:"post_id,omitempty"`
	PersonID  int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	Score     int16     `json:"score,omitempty" url:"score,omitempty"`
	Published time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type PostLikeForm struct {
	PostID   int   `json:"post_id,omitempty" url:"post_id,omitempty"`
	PersonID int   `json:"person_id,omitempty" url:"person_id,omitempty"`
	Score    int16 `json:"score,omitempty" url:"score,omitempty"`
}
type PostSaved struct {
	ID        int32     `json:"id,omitempty" url:"id,omitempty"`
	PostID    int       `json:"post_id,omitempty" url:"post_id,omitempty"`
	PersonID  int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	Published time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type PostSavedForm struct {
	PostID   int `json:"post_id,omitempty" url:"post_id,omitempty"`
	PersonID int `json:"person_id,omitempty" url:"person_id,omitempty"`
}
type PostRead struct {
	ID        int32     `json:"id,omitempty" url:"id,omitempty"`
	PostID    int       `json:"post_id,omitempty" url:"post_id,omitempty"`
	PersonID  int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	Published time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type PostReadForm struct {
	PostID   int `json:"post_id,omitempty" url:"post_id,omitempty"`
	PersonID int `json:"person_id,omitempty" url:"person_id,omitempty"`
}
