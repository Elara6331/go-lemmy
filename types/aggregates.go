package types

type PersonAggregates struct {
	ID           int `json:"id" url:"id,omitempty"`
	PersonID     int `json:"person_id" url:"person_id,omitempty"`
	PostCount    int `json:"post_count" url:"post_count,omitempty"`
	PostScore    int `json:"post_score" url:"post_score,omitempty"`
	CommentCount int `json:"comment_count" url:"comment_count,omitempty"`
	CommentScore int `json:"comment_score" url:"comment_score,omitempty"`
}

type SiteAggregates struct {
	ID                  int `json:"id" url:"id,omitempty"`
	SiteID              int `json:"site_id" url:"site_id,omitempty"`
	Users               int `json:"users" url:"users,omitempty"`
	Posts               int `json:"posts" url:"posts,omitempty"`
	Comments            int `json:"comments" url:"comments,omitempty"`
	Communities         int `json:"communities" url:"communities,omitempty"`
	UsersActiveDay      int `json:"users_active_day" url:"users_active_day,omitempty"`
	UsersActiveWeek     int `json:"users_active_week" url:"users_active_week,omitempty"`
	UsersActiveMonth    int `json:"users_active_month" url:"users_active_month,omitempty"`
	UsersActiveHalfYear int `json:"users_active_half_year" url:"users_active_half_year,omitempty"`
}

type PostAggregates struct {
	ID                     int    `json:"id" url:"id,omitempty"`
	PostID                 int    `json:"post_id" url:"post_id,omitempty"`
	Comments               int    `json:"comments" url:"comments,omitempty"`
	Score                  int    `json:"score" url:"score,omitempty"`
	Upvotes                int    `json:"upvotes" url:"upvotes,omitempty"`
	Downvotes              int    `json:"downvotes" url:"downvotes,omitempty"`
	NewestCommentTimeNecro string `json:"newest_comment_time_necro" url:"newest_comment_time_necro,omitempty"`
	NewestCommentTime      string `json:"newest_comment_time" url:"newest_comment_time,omitempty"`
}

type CommunityAggregates struct {
	ID                  int `json:"id" url:"id,omitempty"`
	CommunityID         int `json:"community_id" url:"community_id,omitempty"`
	Subscribers         int `json:"subscribers" url:"subscribers,omitempty"`
	Posts               int `json:"posts" url:"posts,omitempty"`
	Comments            int `json:"comments" url:"comments,omitempty"`
	UsersActiveDay      int `json:"users_active_day" url:"users_active_day,omitempty"`
	UsersActiveWeek     int `json:"users_active_week" url:"users_active_week,omitempty"`
	UsersActiveMonth    int `json:"users_active_month" url:"users_active_month,omitempty"`
	UsersActiveHalfYear int `json:"users_active_half_year" url:"users_active_half_year,omitempty"`
}

type CommentAggregates struct {
	ID         int `json:"id" url:"id,omitempty"`
	CommentID  int `json:"comment_id" url:"comment_id,omitempty"`
	Score      int `json:"score" url:"score,omitempty"`
	Upvotes    int `json:"upvotes" url:"upvotes,omitempty"`
	Downvotes  int `json:"downvotes" url:"downvotes,omitempty"`
	ChildCount int `json:"child_count" url:"child_count,omitempty"`
}
