package types

type CreatePost struct {
	Name        string           `json:"name" url:"name,omitempty"`
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	URL         Optional[string] `json:"url" url:"url,omitempty"`
	Body        Optional[string] `json:"body" url:"body,omitempty"`
	Honeypot    Optional[string] `json:"honeypot" url:"honeypot,omitempty"`
	NSFW        Optional[bool]   `json:"nsfw" url:"nsfw,omitempty"`
	LanguageID  Optional[int]    `json:"language_id" url:"language_id,omitempty"`
	Auth        string           `json:"auth" url:"auth,omitempty"`
}

type PostResponse struct {
	PostView PostView `json:"post_view" url:"post_view,omitempty"`
	LemmyResponse
}

type GetPost struct {
	ID        Optional[int]    `json:"id" url:"id,omitempty"`
	CommentID Optional[int]    `json:"comment_id" url:"comment_id,omitempty"`
	Auth      Optional[string] `json:"auth" url:"auth,omitempty"`
}

type GetPostResponse struct {
	PostView      PostView                 `json:"post_view" url:"post_view,omitempty"`
	CommunityView CommunityView            `json:"community_view" url:"community_view,omitempty"`
	Moderators    []CommunityModeratorView `json:"moderators" url:"moderators,omitempty"`
	Online        uint                     `json:"online" url:"online,omitempty"`
	LemmyResponse
}

type GetPosts struct {
	Type          Optional[ListingType] `json:"type" url:"type,omitempty"`
	Sort          Optional[SortType]    `json:"sort" url:"sort,omitempty"`
	Page          Optional[int64]       `json:"page" url:"page,omitempty"`
	Limit         Optional[int64]       `json:"limit" url:"limit,omitempty"`
	CommunityID   Optional[int]         `json:"community_id" url:"community_id,omitempty"`
	CommunityName Optional[string]      `json:"community_name" url:"community_name,omitempty"`
	SavedOnly     Optional[bool]        `json:"saved_only" url:"saved_only,omitempty"`
	Auth          Optional[string]      `json:"auth" url:"auth,omitempty"`
}

type GetPostsResponse struct {
	Posts []PostView `json:"posts" url:"posts,omitempty"`
	LemmyResponse
}

type CreatePostLike struct {
	PostID int              `json:"post_id" url:"post_id,omitempty"`
	Score  int16            `json:"score" url:"score,omitempty"`
	Auth   Optional[string] `json:"auth" url:"auth,omitempty"`
}

type EditPost struct {
	PostID     int              `json:"post_id" url:"post_id,omitempty"`
	Name       Optional[string] `json:"name" url:"name,omitempty"`
	URL        Optional[string] `json:"url" url:"url,omitempty"`
	Body       Optional[string] `json:"body" url:"body,omitempty"`
	NSFW       Optional[bool]   `json:"nsfw" url:"nsfw,omitempty"`
	LanguageID Optional[int]    `json:"language_id" url:"language_id,omitempty"`
	Auth       Optional[string] `json:"auth" url:"auth,omitempty"`
}

type DeletePost struct {
	PostID  int              `json:"post_id" url:"post_id,omitempty"`
	Deleted bool             `json:"deleted" url:"deleted,omitempty"`
	Auth    Optional[string] `json:"auth" url:"auth,omitempty"`
}

type RemovePost struct {
	PostID  int              `json:"post_id" url:"post_id,omitempty"`
	Removed bool             `json:"removed" url:"removed,omitempty"`
	Reason  Optional[string] `json:"reason" url:"reason,omitempty"`
	Auth    Optional[string] `json:"auth" url:"auth,omitempty"`
}

type MarkPostAsRead struct {
	PostID int              `json:"post_id" url:"post_id,omitempty"`
	Read   bool             `json:"read" url:"read,omitempty"`
	Auth   Optional[string] `json:"auth" url:"auth,omitempty"`
}

type LockPost struct {
	PostID int              `json:"post_id" url:"post_id,omitempty"`
	Locked bool             `json:"locked" url:"locked,omitempty"`
	Auth   Optional[string] `json:"auth" url:"auth,omitempty"`
}

type StickyPost struct {
	PostID   int              `json:"post_id" url:"post_id,omitempty"`
	Stickied bool             `json:"stickied" url:"stickied,omitempty"`
	Auth     Optional[string] `json:"auth" url:"auth,omitempty"`
}

type SavePost struct {
	PostID int              `json:"post_id" url:"post_id,omitempty"`
	Save   bool             `json:"save" url:"save,omitempty"`
	Auth   Optional[string] `json:"auth" url:"auth,omitempty"`
}

type CreatePostReport struct {
	PostID int              `json:"post_id" url:"post_id,omitempty"`
	Reason string           `json:"reason" url:"reason,omitempty"`
	Auth   Optional[string] `json:"auth" url:"auth,omitempty"`
}

type PostReportResponse struct {
	PostReportView PostReportView `json:"post_report_view" url:"post_report_view,omitempty"`
	LemmyResponse
}

type ResolvePostReport struct {
	ReportID int              `json:"report_id" url:"report_id,omitempty"`
	Resolved bool             `json:"resolved" url:"resolved,omitempty"`
	Auth     Optional[string] `json:"auth" url:"auth,omitempty"`
}

type ListPostReports struct {
	Page           Optional[int64]  `json:"page" url:"page,omitempty"`
	Limit          Optional[int64]  `json:"limit" url:"limit,omitempty"`
	UnresolvedOnly Optional[bool]   `json:"unresolved_only" url:"unresolved_only,omitempty"`
	CommunityID    Optional[int]    `json:"community_id" url:"community_id,omitempty"`
	Auth           Optional[string] `json:"auth" url:"auth,omitempty"`
}

type ListPostReportsResponse struct {
	PostReports []PostReportView `json:"post_reports" url:"post_reports,omitempty"`
	LemmyResponse
}

type GetSiteMetadata struct {
	URL string `json:"url" url:"url,omitempty"`
}

type GetSiteMetadataResponse struct {
	Metadata SiteMetadata `json:"metadata" url:"metadata,omitempty"`
	LemmyResponse
}

type SiteMetadata struct {
	Title         Optional[string] `json:"title" url:"title,omitempty"`
	Description   Optional[string] `json:"description" url:"description,omitempty"`
	Image         Optional[string] `json:"image" url:"image,omitempty"`
	EmbedVideoURL Optional[string] `json:"embed_video_url" url:"embed_video_url,omitempty"`
}
