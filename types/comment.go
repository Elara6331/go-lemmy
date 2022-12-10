package types

type CreateComment struct {
	Content    string           `json:"content,omitempty" url:"content,omitempty"`
	PostID     int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	ParentID   Optional[int]    `json:"parent_id,omitempty" url:"parent_id,omitempty"`
	LanguageID Optional[int]    `json:"language_id,omitempty" url:"language_id,omitempty"`
	FormID     Optional[string] `json:"form_id,omitempty" url:"form_id,omitempty"`
	Auth       string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type GetComment struct {
	ID   int              `json:"id,omitempty" url:"id,omitempty"`
	Auth Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type EditComment struct {
	CommentID     int              `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Content       Optional[string] `json:"content,omitempty" url:"content,omitempty"`
	Distinguished Optional[bool]   `json:"distinguished,omitempty" url:"distinguished,omitempty"`
	LanguageID    Optional[int]    `json:"language_id,omitempty" url:"language_id,omitempty"`
	FormID        Optional[string] `json:"form_id,omitempty" url:"form_id,omitempty"`
	Auth          string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type DeleteComment struct {
	CommentID int    `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Deleted   bool   `json:"deleted,omitempty" url:"deleted,omitempty"`
	Auth      string `json:"auth,omitempty" url:"auth,omitempty"`
}

type RemoveComment struct {
	CommentID int              `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Removed   bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Reason    Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Auth      string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type SaveComment struct {
	CommentID int    `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Save      bool   `json:"save,omitempty" url:"save,omitempty"`
	Auth      string `json:"auth,omitempty" url:"auth,omitempty"`
}

type CommentResponse struct {
	CommentView  CommentView      `json:"comment_view,omitempty" url:"comment_view,omitempty"`
	RecipientIDs []int            `json:"recipient_i_ds,omitempty" url:"recipient_i_ds,omitempty"`
	FormID       Optional[string] `json:"form_id,omitempty" url:"form_id,omitempty"`
	LemmyResponse
}

type CreateCommentLike struct {
	CommentID int    `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Score     int16  `json:"score,omitempty" url:"score,omitempty"`
	Auth      string `json:"auth,omitempty" url:"auth,omitempty"`
}

type GetComments struct {
	Type          Optional[ListingType]     `json:"type,omitempty" url:"type,omitempty"`
	Sort          Optional[CommentSortType] `json:"sort,omitempty" url:"sort,omitempty"`
	MaxDepth      Optional[int]             `json:"max_depth,omitempty" url:"max_depth,omitempty"`
	Page          Optional[int]             `json:"page,omitempty" url:"page,omitempty"`
	Limit         Optional[int]             `json:"limit,omitempty" url:"limit,omitempty"`
	CommunityID   Optional[int]             `json:"community_id,omitempty" url:"community_id,omitempty"`
	CommunityName Optional[string]          `json:"community_name,omitempty" url:"community_name,omitempty"`
	PostID        Optional[int]             `json:"post_id,omitempty" url:"post_id,omitempty"`
	ParentID      Optional[int]             `json:"parent_id,omitempty" url:"parent_id,omitempty"`
	SavedOnly     Optional[bool]            `json:"saved_only,omitempty" url:"saved_only,omitempty"`
	Auth          Optional[string]          `json:"auth,omitempty" url:"auth,omitempty"`
}

type GetCommentsResponse struct {
	Comments []CommentView `json:"comments,omitempty" url:"comments,omitempty"`
	LemmyResponse
}

type CreateCommentReport struct {
	CommentID int    `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Reason    string `json:"reason,omitempty" url:"reason,omitempty"`
	Auth      string `json:"auth,omitempty" url:"auth,omitempty"`
}

type CommentReportResponse struct {
	CommentReportView CommentReportView `json:"comment_report_view,omitempty" url:"comment_report_view,omitempty"`
	LemmyResponse
}

type ResolveCommentReport struct {
	ReportID int    `json:"report_id,omitempty" url:"report_id,omitempty"`
	Resolved bool   `json:"resolved,omitempty" url:"resolved,omitempty"`
	Auth     string `json:"auth,omitempty" url:"auth,omitempty"`
}

type ListCommentReports struct {
	Page           Optional[int]  `json:"page,omitempty" url:"page,omitempty"`
	Limit          Optional[int]  `json:"limit,omitempty" url:"limit,omitempty"`
	UnresolvedOnly Optional[bool] `json:"unresolved_only,omitempty" url:"unresolved_only,omitempty"`
	CommunityID    Optional[int]  `json:"community_id,omitempty" url:"community_id,omitempty"`
	Auth           string         `json:"auth,omitempty" url:"auth,omitempty"`
}

type ListCommentReportsResponse struct {
	CommentReports []CommentReportView `json:"comment_reports,omitempty" url:"comment_reports,omitempty"`
	LemmyResponse
}
