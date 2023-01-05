package types

type Search struct {
	Q             string                `json:"q,omitempty" url:"q,omitempty"`
	CommunityID   Optional[int]         `json:"community_id,omitempty" url:"community_id,omitempty"`
	CommunityName Optional[string]      `json:"community_name,omitempty" url:"community_name,omitempty"`
	CreatorID     Optional[int]         `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	Type          Optional[SearchType]  `json:"type_,omitempty" url:"type_,omitempty"`
	Sort          Optional[SortType]    `json:"sort,omitempty" url:"sort,omitempty"`
	ListingType   Optional[ListingType] `json:"listing_type,omitempty" url:"listing_type,omitempty"`
	Page          Optional[int64]       `json:"page,omitempty" url:"page,omitempty"`
	Limit         Optional[int64]       `json:"limit,omitempty" url:"limit,omitempty"`
	Auth          Optional[string]      `json:"auth,omitempty" url:"auth,omitempty"`
}
type SearchResponse struct {
	Type        string           `json:"type_,omitempty" url:"type_,omitempty"`
	Comments    []CommentView    `json:"comments,omitempty" url:"comments,omitempty"`
	Posts       []PostView       `json:"posts,omitempty" url:"posts,omitempty"`
	Communities []CommunityView  `json:"communities,omitempty" url:"communities,omitempty"`
	Users       []PersonViewSafe `json:"users,omitempty" url:"users,omitempty"`
	LemmyResponse
}
type ResolveObject struct {
	Q    string           `json:"q,omitempty" url:"q,omitempty"`
	Auth Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}
type ResolveObjectResponse struct {
	Comment   Optional[CommentView]    `json:"comment,omitempty" url:"comment,omitempty"`
	Post      Optional[PostView]       `json:"post,omitempty" url:"post,omitempty"`
	Community Optional[CommunityView]  `json:"community,omitempty" url:"community,omitempty"`
	Person    Optional[PersonViewSafe] `json:"person,omitempty" url:"person,omitempty"`
	LemmyResponse
}
type GetModlog struct {
	ModPersonID Optional[int]    `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	CommunityID Optional[int]    `json:"community_id,omitempty" url:"community_id,omitempty"`
	Page        Optional[int64]  `json:"page,omitempty" url:"page,omitempty"`
	Limit       Optional[int64]  `json:"limit,omitempty" url:"limit,omitempty"`
	Auth        Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetModlogResponse struct {
	RemovedPosts           []ModRemovePostView        `json:"removed_posts,omitempty" url:"removed_posts,omitempty"`
	LockedPosts            []ModLockPostView          `json:"locked_posts,omitempty" url:"locked_posts,omitempty"`
	StickiedPosts          []ModStickyPostView        `json:"stickied_posts,omitempty" url:"stickied_posts,omitempty"`
	RemovedComments        []ModRemoveCommentView     `json:"removed_comments,omitempty" url:"removed_comments,omitempty"`
	RemovedCommunities     []ModRemoveCommunityView   `json:"removed_communities,omitempty" url:"removed_communities,omitempty"`
	BannedFromCommunity    []ModBanFromCommunityView  `json:"banned_from_community,omitempty" url:"banned_from_community,omitempty"`
	Banned                 []ModBanView               `json:"banned,omitempty" url:"banned,omitempty"`
	AddedToCommunity       []ModAddCommunityView      `json:"added_to_community,omitempty" url:"added_to_community,omitempty"`
	TransferredToCommunity []ModTransferCommunityView `json:"transferred_to_community,omitempty" url:"transferred_to_community,omitempty"`
	Added                  []ModAddView               `json:"added,omitempty" url:"added,omitempty"`
	HiddenCommunities      []ModHideCommunityView     `json:"hidden_communities,omitempty" url:"hidden_communities,omitempty"`
	LemmyResponse
}
type CreateSite struct {
	Name                       string           `json:"name,omitempty" url:"name,omitempty"`
	Sidebar                    Optional[string] `json:"sidebar,omitempty" url:"sidebar,omitempty"`
	Description                Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	Icon                       Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                     Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	EnableDownvotes            Optional[bool]   `json:"enable_downvotes,omitempty" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]   `json:"open_registration,omitempty" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]   `json:"enable_nsfw,omitempty" url:"enable_nsfw,omitempty"`
	CommunityCreationAdminOnly Optional[bool]   `json:"community_creation_admin_only,omitempty" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]   `json:"require_email_verification,omitempty" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]   `json:"require_application,omitempty" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string] `json:"application_question,omitempty" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]   `json:"private_instance,omitempty" url:"private_instance,omitempty"`
	DefaultTheme               Optional[string] `json:"default_theme,omitempty" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string] `json:"default_post_listing_type,omitempty" url:"default_post_listing_type,omitempty"`
	Auth                       string           `json:"auth,omitempty" url:"auth,omitempty"`
}
type EditSite struct {
	Name                       Optional[string] `json:"name,omitempty" url:"name,omitempty"`
	Sidebar                    Optional[string] `json:"sidebar,omitempty" url:"sidebar,omitempty"`
	Description                Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	Icon                       Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                     Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	EnableDownvotes            Optional[bool]   `json:"enable_downvotes,omitempty" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]   `json:"open_registration,omitempty" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]   `json:"enable_nsfw,omitempty" url:"enable_nsfw,omitempty"`
	CommunityCreationAdminOnly Optional[bool]   `json:"community_creation_admin_only,omitempty" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]   `json:"require_email_verification,omitempty" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]   `json:"require_application,omitempty" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string] `json:"application_question,omitempty" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]   `json:"private_instance,omitempty" url:"private_instance,omitempty"`
	DefaultTheme               Optional[string] `json:"default_theme,omitempty" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string] `json:"default_post_listing_type,omitempty" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string] `json:"legal_information,omitempty" url:"legal_information,omitempty"`
	Auth                       string           `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetSite struct {
	Auth Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}
type SiteResponse struct {
	SiteView SiteView `json:"site_view,omitempty" url:"site_view,omitempty"`
	LemmyResponse
}
type GetSiteResponse struct {
	SiteView           Optional[SiteView]           `json:"site_view,omitempty" url:"site_view,omitempty"`
	Admins             []PersonViewSafe             `json:"admins,omitempty" url:"admins,omitempty"`
	Online             uint                         `json:"online,omitempty" url:"online,omitempty"`
	Version            string                       `json:"version,omitempty" url:"version,omitempty"`
	MyUser             Optional[MyUserInfo]         `json:"my_user,omitempty" url:"my_user,omitempty"`
	FederatedInstances Optional[FederatedInstances] `json:"federated_instances,omitempty" url:"federated_instances,omitempty"`
	LemmyResponse
}
type MyUserInfo struct {
	LocalUserView   LocalUserSettingsView    `json:"local_user_view,omitempty" url:"local_user_view,omitempty"`
	Follows         []CommunityFollowerView  `json:"follows,omitempty" url:"follows,omitempty"`
	Moderates       []CommunityModeratorView `json:"moderates,omitempty" url:"moderates,omitempty"`
	CommunityBlocks []CommunityBlockView     `json:"community_blocks,omitempty" url:"community_blocks,omitempty"`
	PersonBlocks    []PersonBlockView        `json:"person_blocks,omitempty" url:"person_blocks,omitempty"`
}
type LeaveAdmin struct {
	Auth string `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetSiteConfig struct {
	Auth string `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetSiteConfigResponse struct {
	ConfigHjson string `json:"config_hjson,omitempty" url:"config_hjson,omitempty"`
	LemmyResponse
}
type SaveSiteConfig struct {
	ConfigHjson string `json:"config_hjson,omitempty" url:"config_hjson,omitempty"`
	Auth        string `json:"auth,omitempty" url:"auth,omitempty"`
}
type FederatedInstances struct {
	Linked  []string           `json:"linked,omitempty" url:"linked,omitempty"`
	Allowed Optional[[]string] `json:"allowed,omitempty" url:"allowed,omitempty"`
	Blocked Optional[[]string] `json:"blocked,omitempty" url:"blocked,omitempty"`
}
type ListRegistrationApplications struct {
	UnreadOnly Optional[bool]  `json:"unread_only,omitempty" url:"unread_only,omitempty"`
	Page       Optional[int64] `json:"page,omitempty" url:"page,omitempty"`
	Limit      Optional[int64] `json:"limit,omitempty" url:"limit,omitempty"`
	Auth       string          `json:"auth,omitempty" url:"auth,omitempty"`
}
type ListRegistrationApplicationsResponse struct {
	RegistrationApplications []RegistrationApplicationView `json:"registration_applications,omitempty" url:"registration_applications,omitempty"`
	LemmyResponse
}
type ApproveRegistrationApplication struct {
	ID         int32            `json:"id,omitempty" url:"id,omitempty"`
	Approve    bool             `json:"approve,omitempty" url:"approve,omitempty"`
	DenyReason Optional[string] `json:"deny_reason,omitempty" url:"deny_reason,omitempty"`
	Auth       string           `json:"auth,omitempty" url:"auth,omitempty"`
}
type RegistrationApplicationResponse struct {
	RegistrationApplication RegistrationApplicationView `json:"registration_application,omitempty" url:"registration_application,omitempty"`
	LemmyResponse
}
type GetUnreadRegistrationApplicationCount struct {
	Auth string `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetUnreadRegistrationApplicationCountResponse struct {
	RegistrationApplications int64 `json:"registration_applications,omitempty" url:"registration_applications,omitempty"`
	LemmyResponse
}
