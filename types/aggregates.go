package types

type PersonAggregates struct {
	ID           int `json:"id,omitempty" url:"id,omitempty"`
	PersonID     int `json:"person_id,omitempty" url:"person_id,omitempty"`
	PostCount    int `json:"post_count,omitempty" url:"post_count,omitempty"`
	PostScore    int `json:"post_score,omitempty" url:"post_score,omitempty"`
	CommentCount int `json:"comment_count,omitempty" url:"comment_count,omitempty"`
	CommentScore int `json:"comment_score,omitempty" url:"comment_score,omitempty"`
}

type SiteAggregates struct {
	ID                  int `json:"id,omitempty" url:"id,omitempty"`
	SiteID              int `json:"site_id,omitempty" url:"site_id,omitempty"`
	Users               int `json:"users,omitempty" url:"users,omitempty"`
	Posts               int `json:"posts,omitempty" url:"posts,omitempty"`
	Comments            int `json:"comments,omitempty" url:"comments,omitempty"`
	Communities         int `json:"communities,omitempty" url:"communities,omitempty"`
	UsersActiveDay      int `json:"users_active_day,omitempty" url:"users_active_day,omitempty"`
	UsersActiveWeek     int `json:"users_active_week,omitempty" url:"users_active_week,omitempty"`
	UsersActiveMonth    int `json:"users_active_month,omitempty" url:"users_active_month,omitempty"`
	UsersActiveHalfYear int `json:"users_active_half_year,omitempty" url:"users_active_half_year,omitempty"`
}

type PostAggregates struct {
	ID                     int    `json:"id,omitempty" url:"id,omitempty"`
	PostID                 int    `json:"post_id,omitempty" url:"post_id,omitempty"`
	Comments               int    `json:"comments,omitempty" url:"comments,omitempty"`
	Score                  int    `json:"score,omitempty" url:"score,omitempty"`
	Upvotes                int    `json:"upvotes,omitempty" url:"upvotes,omitempty"`
	Downvotes              int    `json:"downvotes,omitempty" url:"downvotes,omitempty"`
	NewestCommentTimeNecro string `json:"newest_comment_time_necro,omitempty" url:"newest_comment_time_necro,omitempty"`
	NewestCommentTime      string `json:"newest_comment_time,omitempty" url:"newest_comment_time,omitempty"`
}

type CommunityAggregates struct {
	ID                  int `json:"id,omitempty" url:"id,omitempty"`
	CommunityID         int `json:"community_id,omitempty" url:"community_id,omitempty"`
	Subscribers         int `json:"subscribers,omitempty" url:"subscribers,omitempty"`
	Posts               int `json:"posts,omitempty" url:"posts,omitempty"`
	Comments            int `json:"comments,omitempty" url:"comments,omitempty"`
	UsersActiveDay      int `json:"users_active_day,omitempty" url:"users_active_day,omitempty"`
	UsersActiveWeek     int `json:"users_active_week,omitempty" url:"users_active_week,omitempty"`
	UsersActiveMonth    int `json:"users_active_month,omitempty" url:"users_active_month,omitempty"`
	UsersActiveHalfYear int `json:"users_active_half_year,omitempty" url:"users_active_half_year,omitempty"`
}

type CommentAggregates struct {
	ID         int `json:"id,omitempty" url:"id,omitempty"`
	CommentID  int `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Score      int `json:"score,omitempty" url:"score,omitempty"`
	Upvotes    int `json:"upvotes,omitempty" url:"upvotes,omitempty"`
	Downvotes  int `json:"downvotes,omitempty" url:"downvotes,omitempty"`
	ChildCount int `json:"child_count,omitempty" url:"child_count,omitempty"`
}
