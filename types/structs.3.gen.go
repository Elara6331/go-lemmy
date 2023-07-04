//  Source: lemmy/crates/db_schema/src/aggregates/structs.rs
// Code generated by go.elara.ws/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type CommentAggregates struct {
	ID         int32     `json:"id" url:"id,omitempty"`
	CommentID  int       `json:"comment_id" url:"comment_id,omitempty"`
	Score      int64     `json:"score" url:"score,omitempty"`
	Upvotes    int64     `json:"upvotes" url:"upvotes,omitempty"`
	Downvotes  int64     `json:"downvotes" url:"downvotes,omitempty"`
	Published  LemmyTime `json:"published" url:"published,omitempty"`
	ChildCount int32     `json:"child_count" url:"child_count,omitempty"`
	HotRank    int32     `json:"hot_rank" url:"hot_rank,omitempty"`
}
type CommunityAggregates struct {
	ID                  int32     `json:"id" url:"id,omitempty"`
	CommunityID         int       `json:"community_id" url:"community_id,omitempty"`
	Subscribers         int64     `json:"subscribers" url:"subscribers,omitempty"`
	Posts               int64     `json:"posts" url:"posts,omitempty"`
	Comments            int64     `json:"comments" url:"comments,omitempty"`
	Published           LemmyTime `json:"published" url:"published,omitempty"`
	UsersActiveDay      int64     `json:"users_active_day" url:"users_active_day,omitempty"`
	UsersActiveWeek     int64     `json:"users_active_week" url:"users_active_week,omitempty"`
	UsersActiveMonth    int64     `json:"users_active_month" url:"users_active_month,omitempty"`
	UsersActiveHalfYear int64     `json:"users_active_half_year" url:"users_active_half_year,omitempty"`
	HotRank             int32     `json:"hot_rank" url:"hot_rank,omitempty"`
}
type PersonAggregates struct {
	ID           int32 `json:"id" url:"id,omitempty"`
	PersonID     int   `json:"person_id" url:"person_id,omitempty"`
	PostCount    int64 `json:"post_count" url:"post_count,omitempty"`
	PostScore    int64 `json:"post_score" url:"post_score,omitempty"`
	CommentCount int64 `json:"comment_count" url:"comment_count,omitempty"`
	CommentScore int64 `json:"comment_score" url:"comment_score,omitempty"`
}
type PostAggregates struct {
	ID                     int32     `json:"id" url:"id,omitempty"`
	PostID                 int       `json:"post_id" url:"post_id,omitempty"`
	Comments               int64     `json:"comments" url:"comments,omitempty"`
	Score                  int64     `json:"score" url:"score,omitempty"`
	Upvotes                int64     `json:"upvotes" url:"upvotes,omitempty"`
	Downvotes              int64     `json:"downvotes" url:"downvotes,omitempty"`
	Published              LemmyTime `json:"published" url:"published,omitempty"`
	NewestCommentTimeNecro LemmyTime `json:"newest_comment_time_necro" url:"newest_comment_time_necro,omitempty"`
	NewestCommentTime      LemmyTime `json:"newest_comment_time" url:"newest_comment_time,omitempty"`
	FeaturedCommunity      bool      `json:"featured_community" url:"featured_community,omitempty"`
	FeaturedLocal          bool      `json:"featured_local" url:"featured_local,omitempty"`
	HotRank                int32     `json:"hot_rank" url:"hot_rank,omitempty"`
	HotRankActive          int32     `json:"hot_rank_active" url:"hot_rank_active,omitempty"`
}
type PersonPostAggregates struct {
	ID           int32     `json:"id" url:"id,omitempty"`
	PersonID     int       `json:"person_id" url:"person_id,omitempty"`
	PostID       int       `json:"post_id" url:"post_id,omitempty"`
	ReadComments int64     `json:"read_comments" url:"read_comments,omitempty"`
	Published    LemmyTime `json:"published" url:"published,omitempty"`
}
type PersonPostAggregatesForm struct {
	PersonID     int       `json:"person_id" url:"person_id,omitempty"`
	PostID       int       `json:"post_id" url:"post_id,omitempty"`
	ReadComments int64     `json:"read_comments" url:"read_comments,omitempty"`
	Published    LemmyTime `json:"published" url:"published,omitempty"`
}
type SiteAggregates struct {
	ID                  int32 `json:"id" url:"id,omitempty"`
	SiteID              int   `json:"site_id" url:"site_id,omitempty"`
	Users               int64 `json:"users" url:"users,omitempty"`
	Posts               int64 `json:"posts" url:"posts,omitempty"`
	Comments            int64 `json:"comments" url:"comments,omitempty"`
	Communities         int64 `json:"communities" url:"communities,omitempty"`
	UsersActiveDay      int64 `json:"users_active_day" url:"users_active_day,omitempty"`
	UsersActiveWeek     int64 `json:"users_active_week" url:"users_active_week,omitempty"`
	UsersActiveMonth    int64 `json:"users_active_month" url:"users_active_month,omitempty"`
	UsersActiveHalfYear int64 `json:"users_active_half_year" url:"users_active_half_year,omitempty"`
}
