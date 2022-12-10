package types

type CreatePost struct {
	Name        string           `json:"name,omitempty" url:"name,omitempty"`
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	URL         Optional[string] `json:"url,omitempty" url:"url,omitempty"`
	Body        Optional[string] `json:"body,omitempty" url:"body,omitempty"`
	Honeypot    Optional[string] `json:"honeypot,omitempty" url:"honeypot,omitempty"`
	NSFW        Optional[bool]   `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	LanguageID  Optional[int]    `json:"language_id,omitempty" url:"language_id,omitempty"`
	Auth        string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type PostResponse struct {
	PostView PostView `json:"post_view,omitempty" url:"post_view,omitempty"`
	LemmyResponse
}

type GetPost struct {
	ID        Optional[int]    `json:"id,omitempty" url:"id,omitempty"`
	CommentID Optional[int]    `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Auth      Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type GetPostResponse struct {
	PostView      PostView                 `json:"post_view,omitempty" url:"post_view,omitempty"`
	CommunityView CommunityView            `json:"community_view,omitempty" url:"community_view,omitempty"`
	Moderators    []CommunityModeratorView `json:"moderators,omitempty" url:"moderators,omitempty"`
	Online        uint                     `json:"online,omitempty" url:"online,omitempty"`
	LemmyResponse
}

type GetPosts struct {
	Type          Optional[ListingType] `json:"type,omitempty" url:"type,omitempty"`
	Sort          Optional[SortType]    `json:"sort,omitempty" url:"sort,omitempty"`
	Page          Optional[int64]       `json:"page,omitempty" url:"page,omitempty"`
	Limit         Optional[int64]       `json:"limit,omitempty" url:"limit,omitempty"`
	CommunityID   Optional[int]         `json:"community_id,omitempty" url:"community_id,omitempty"`
	CommunityName Optional[string]      `json:"community_name,omitempty" url:"community_name,omitempty"`
	SavedOnly     Optional[bool]        `json:"saved_only,omitempty" url:"saved_only,omitempty"`
	Auth          Optional[string]      `json:"auth,omitempty" url:"auth,omitempty"`
}

type GetPostsResponse struct {
	Posts []PostView `json:"posts,omitempty" url:"posts,omitempty"`
	LemmyResponse
}

type CreatePostLike struct {
	PostID int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Score  int16            `json:"score,omitempty" url:"score,omitempty"`
	Auth   Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type EditPost struct {
	PostID     int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Name       Optional[string] `json:"name,omitempty" url:"name,omitempty"`
	URL        Optional[string] `json:"url,omitempty" url:"url,omitempty"`
	Body       Optional[string] `json:"body,omitempty" url:"body,omitempty"`
	NSFW       Optional[bool]   `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	LanguageID Optional[int]    `json:"language_id,omitempty" url:"language_id,omitempty"`
	Auth       Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type DeletePost struct {
	PostID  int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Deleted bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	Auth    Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type RemovePost struct {
	PostID  int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Removed bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Reason  Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Auth    Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type MarkPostAsRead struct {
	PostID int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Read   bool             `json:"read,omitempty" url:"read,omitempty"`
	Auth   Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type LockPost struct {
	PostID int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Locked bool             `json:"locked,omitempty" url:"locked,omitempty"`
	Auth   Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type StickyPost struct {
	PostID   int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Stickied bool             `json:"stickied,omitempty" url:"stickied,omitempty"`
	Auth     Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type SavePost struct {
	PostID int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Save   bool             `json:"save,omitempty" url:"save,omitempty"`
	Auth   Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type CreatePostReport struct {
	PostID int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Reason string           `json:"reason,omitempty" url:"reason,omitempty"`
	Auth   Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type PostReportResponse struct {
	PostReportView PostReportView `json:"post_report_view,omitempty" url:"post_report_view,omitempty"`
	LemmyResponse
}

type ResolvePostReport struct {
	ReportID int              `json:"report_id,omitempty" url:"report_id,omitempty"`
	Resolved bool             `json:"resolved,omitempty" url:"resolved,omitempty"`
	Auth     Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type ListPostReports struct {
	Page           Optional[int64]  `json:"page,omitempty" url:"page,omitempty"`
	Limit          Optional[int64]  `json:"limit,omitempty" url:"limit,omitempty"`
	UnresolvedOnly Optional[bool]   `json:"unresolved_only,omitempty" url:"unresolved_only,omitempty"`
	CommunityID    Optional[int]    `json:"community_id,omitempty" url:"community_id,omitempty"`
	Auth           Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type ListPostReportsResponse struct {
	PostReports []PostReportView `json:"post_reports,omitempty" url:"post_reports,omitempty"`
	LemmyResponse
}

type GetSiteMetadata struct {
	URL string `json:"url,omitempty" url:"url,omitempty"`
}

type GetSiteMetadataResponse struct {
	Metadata SiteMetadata `json:"metadata,omitempty" url:"metadata,omitempty"`
	LemmyResponse
}

type SiteMetadata struct {
	Title         Optional[string] `json:"title,omitempty" url:"title,omitempty"`
	Description   Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	Image         Optional[string] `json:"image,omitempty" url:"image,omitempty"`
	EmbedVideoURL Optional[string] `json:"embed_video_url,omitempty" url:"embed_video_url,omitempty"`
}
