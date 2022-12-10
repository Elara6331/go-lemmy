package types

type CreateComment struct {
	Content    string           `json:"content" url:"content,omitempty"`
	PostID     int              `json:"post_id" url:"post_id,omitempty"`
	ParentID   Optional[int]    `json:"parent_id" url:"parent_id,omitempty"`
	LanguageID Optional[int]    `json:"language_id" url:"language_id,omitempty"`
	FormID     Optional[string] `json:"form_id" url:"form_id,omitempty"`
	Auth       string           `json:"auth" url:"auth,omitempty"`
}

type GetComment struct {
	ID   int              `json:"id" url:"id,omitempty"`
	Auth Optional[string] `json:"auth" url:"auth,omitempty"`
}

type EditComment struct {
	CommentID     int              `json:"comment_id" url:"comment_id,omitempty"`
	Content       Optional[string] `json:"content" url:"content,omitempty"`
	Distinguished Optional[bool]   `json:"distinguished" url:"distinguished,omitempty"`
	LanguageID    Optional[int]    `json:"language_id" url:"language_id,omitempty"`
	FormID        Optional[string] `json:"form_id" url:"form_id,omitempty"`
	Auth          string           `json:"auth" url:"auth,omitempty"`
}

type DeleteComment struct {
	CommentID int    `json:"comment_id" url:"comment_id,omitempty"`
	Deleted   bool   `json:"deleted" url:"deleted,omitempty"`
	Auth      string `json:"auth" url:"auth,omitempty"`
}

type RemoveComment struct {
	CommentID int              `json:"comment_id" url:"comment_id,omitempty"`
	Removed   bool             `json:"removed" url:"removed,omitempty"`
	Reason    Optional[string] `json:"reason" url:"reason,omitempty"`
	Auth      string           `json:"auth" url:"auth,omitempty"`
}

type SaveComment struct {
	CommentID int    `json:"comment_id" url:"comment_id,omitempty"`
	Save      bool   `json:"save" url:"save,omitempty"`
	Auth      string `json:"auth" url:"auth,omitempty"`
}

type CommentResponse struct {
	CommentView  CommentView      `json:"comment_view" url:"comment_view,omitempty"`
	RecipientIDs []int            `json:"recipient_i_ds" url:"recipient_i_ds,omitempty"`
	FormID       Optional[string] `json:"form_id" url:"form_id,omitempty"`
	LemmyResponse
}

type CreateCommentLike struct {
	CommentID int    `json:"comment_id" url:"comment_id,omitempty"`
	Score     int16  `json:"score" url:"score,omitempty"`
	Auth      string `json:"auth" url:"auth,omitempty"`
}

type GetComments struct {
	Type          Optional[ListingType]     `json:"type" url:"type,omitempty"`
	Sort          Optional[CommentSortType] `json:"sort" url:"sort,omitempty"`
	MaxDepth      Optional[int]             `json:"max_depth" url:"max_depth,omitempty"`
	Page          Optional[int]             `json:"page" url:"page,omitempty"`
	Limit         Optional[int]             `json:"limit" url:"limit,omitempty"`
	CommunityID   Optional[int]             `json:"community_id" url:"community_id,omitempty"`
	CommunityName Optional[string]          `json:"community_name" url:"community_name,omitempty"`
	PostID        Optional[int]             `json:"post_id" url:"post_id,omitempty"`
	ParentID      Optional[int]             `json:"parent_id" url:"parent_id,omitempty"`
	SavedOnly     Optional[bool]            `json:"saved_only" url:"saved_only,omitempty"`
	Auth          Optional[string]          `json:"auth" url:"auth,omitempty"`
}

type GetCommentsResponse struct {
	Comments []CommentView `json:"comments" url:"comments,omitempty"`
	LemmyResponse
}

type CreateCommentReport struct {
	CommentID int    `json:"comment_id" url:"comment_id,omitempty"`
	Reason    string `json:"reason" url:"reason,omitempty"`
	Auth      string `json:"auth" url:"auth,omitempty"`
}

type CommentReportResponse struct {
	CommentReportView CommentReportView `json:"comment_report_view" url:"comment_report_view,omitempty"`
	LemmyResponse
}

type ResolveCommentReport struct {
	ReportID int    `json:"report_id" url:"report_id,omitempty"`
	Resolved bool   `json:"resolved" url:"resolved,omitempty"`
	Auth     string `json:"auth" url:"auth,omitempty"`
}

type ListCommentReports struct {
	Page           Optional[int]  `json:"page" url:"page,omitempty"`
	Limit          Optional[int]  `json:"limit" url:"limit,omitempty"`
	UnresolvedOnly Optional[bool] `json:"unresolved_only" url:"unresolved_only,omitempty"`
	CommunityID    Optional[int]  `json:"community_id" url:"community_id,omitempty"`
	Auth           string         `json:"auth" url:"auth,omitempty"`
}

type ListCommentReportsResponse struct {
	CommentReports []CommentReportView `json:"comment_reports" url:"comment_reports,omitempty"`
	LemmyResponse
}
