package types

import "time"

type CommentAggregates struct {
	ID        int32     `json:"id,omitempty" url:"id,omitempty"`
	CommentID int       `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Score     int64     `json:"score,omitempty" url:"score,omitempty"`
	Upvotes   int64     `json:"upvotes,omitempty" url:"upvotes,omitempty"`
	Downvotes int64     `json:"downvotes,omitempty" url:"downvotes,omitempty"`
	Published time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type CommunityAggregates struct {
	ID                  int32     `json:"id,omitempty" url:"id,omitempty"`
	CommunityID         int       `json:"community_id,omitempty" url:"community_id,omitempty"`
	Subscribers         int64     `json:"subscribers,omitempty" url:"subscribers,omitempty"`
	Posts               int64     `json:"posts,omitempty" url:"posts,omitempty"`
	Comments            int64     `json:"comments,omitempty" url:"comments,omitempty"`
	Published           time.Time `json:"published,omitempty" url:"published,omitempty"`
	UsersActiveDay      int64     `json:"users_active_day,omitempty" url:"users_active_day,omitempty"`
	UsersActiveWeek     int64     `json:"users_active_week,omitempty" url:"users_active_week,omitempty"`
	UsersActiveMonth    int64     `json:"users_active_month,omitempty" url:"users_active_month,omitempty"`
	UsersActiveHalfYear int64     `json:"users_active_half_year,omitempty" url:"users_active_half_year,omitempty"`
}
type PersonAggregates struct {
	ID           int32 `json:"id,omitempty" url:"id,omitempty"`
	PersonID     int   `json:"person_id,omitempty" url:"person_id,omitempty"`
	PostCount    int64 `json:"post_count,omitempty" url:"post_count,omitempty"`
	PostScore    int64 `json:"post_score,omitempty" url:"post_score,omitempty"`
	CommentCount int64 `json:"comment_count,omitempty" url:"comment_count,omitempty"`
	CommentScore int64 `json:"comment_score,omitempty" url:"comment_score,omitempty"`
}
type PostAggregates struct {
	ID                     int32     `json:"id,omitempty" url:"id,omitempty"`
	PostID                 int       `json:"post_id,omitempty" url:"post_id,omitempty"`
	Comments               int64     `json:"comments,omitempty" url:"comments,omitempty"`
	Score                  int64     `json:"score,omitempty" url:"score,omitempty"`
	Upvotes                int64     `json:"upvotes,omitempty" url:"upvotes,omitempty"`
	Downvotes              int64     `json:"downvotes,omitempty" url:"downvotes,omitempty"`
	Stickied               bool      `json:"stickied,omitempty" url:"stickied,omitempty"`
	Published              time.Time `json:"published,omitempty" url:"published,omitempty"`
	NewestCommentTimeNecro time.Time `json:"newest_comment_time_necro,omitempty" url:"newest_comment_time_necro,omitempty"`
	NewestCommentTime      time.Time `json:"newest_comment_time,omitempty" url:"newest_comment_time,omitempty"`
}
type SiteAggregates struct {
	ID                  int32 `json:"id,omitempty" url:"id,omitempty"`
	SiteID              int32 `json:"site_id,omitempty" url:"site_id,omitempty"`
	Users               int64 `json:"users,omitempty" url:"users,omitempty"`
	Posts               int64 `json:"posts,omitempty" url:"posts,omitempty"`
	Comments            int64 `json:"comments,omitempty" url:"comments,omitempty"`
	Communities         int64 `json:"communities,omitempty" url:"communities,omitempty"`
	UsersActiveDay      int64 `json:"users_active_day,omitempty" url:"users_active_day,omitempty"`
	UsersActiveWeek     int64 `json:"users_active_week,omitempty" url:"users_active_week,omitempty"`
	UsersActiveMonth    int64 `json:"users_active_month,omitempty" url:"users_active_month,omitempty"`
	UsersActiveHalfYear int64 `json:"users_active_half_year,omitempty" url:"users_active_half_year,omitempty"`
}
