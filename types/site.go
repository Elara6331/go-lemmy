package types

type Search struct {
	CommunityID   Optional[int]         `json:"community_id" url:"community_id,omitempty"`
	CommunityName Optional[string]      `json:"community_name" url:"community_name,omitempty"`
	CreatorID     Optional[int]         `json:"creator_id" url:"creator_id,omitempty"`
	Limit         Optional[int]         `json:"limit" url:"limit,omitempty"`
	ListingType   Optional[ListingType] `json:"listing_type" url:"listing_type,omitempty"`
	Page          Optional[int]         `json:"page" url:"page,omitempty"`
	Query         string                `json:"q" url:"q,omitempty"`
	Sort          Optional[SortType]    `json:"sort" url:"sort,omitempty"`
	Type          Optional[SearchType]  `json:"type_" url:"type_,omitempty"`
}

type SearchResponse struct {
	Type        string           `json:"type" url:"type,omitempty"`
	Comments    []CommentView    `json:"comments" url:"comments,omitempty"`
	Posts       []PostView       `json:"posts" url:"posts,omitempty"`
	Communities []CommunityView  `json:"communities" url:"communities,omitempty"`
	Users       []PersonViewSafe `json:"users" url:"users,omitempty"`
	LemmyResponse
}

type ResolveObject struct {
	Query string           `json:"q" url:"q,omitempty"`
	Auth  Optional[string] `json:"auth" url:"auth,omitempty"`
}

type ResolveObjectResponse struct {
	Comment   Optional[CommentView]    `json:"comment" url:"comment,omitempty"`
	Post      Optional[PostView]       `json:"post" url:"post,omitempty"`
	Community Optional[CommunityView]  `json:"community" url:"community,omitempty"`
	Person    Optional[PersonViewSafe] `json:"person" url:"person,omitempty"`
	LemmyResponse
}

type GetModlog struct {
	ModPersonID   Optional[int]              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	CommunityID   Optional[int]              `json:"community_id" url:"community_id,omitempty"`
	Page          Optional[int64]            `json:"page" url:"page,omitempty"`
	Limit         Optional[int64]            `json:"limit" url:"limit,omitempty"`
	Auth          Optional[string]           `json:"auth" url:"auth,omitempty"`
	Type          Optional[ModlogActionType] `json:"type" url:"type,omitempty"`
	OtherPersonID Optional[int]              `json:"other_person_id" url:"other_person_id,omitempty"`
}

type GetModlogResponse struct {
	RemovedPosts           []ModRemovePostView        `json:"removed_posts" url:"removed_posts,omitempty"`
	LockedPosts            []ModLockPostView          `json:"locked_posts" url:"locked_posts,omitempty"`
	StickiedPosts          []ModStickyPostView        `json:"stickied_posts" url:"stickied_posts,omitempty"`
	RemovedComments        []ModRemoveCommentView     `json:"removed_comments" url:"removed_comments,omitempty"`
	RemovedCommunities     []ModRemoveCommunityView   `json:"removed_communities" url:"removed_communities,omitempty"`
	BannedFromCommunity    []ModBanFromCommunityView  `json:"banned_from_community" url:"banned_from_community,omitempty"`
	Banned                 []ModBanView               `json:"banned" url:"banned,omitempty"`
	AddedToCommunity       []ModAddCommunityView      `json:"added_to_community" url:"added_to_community,omitempty"`
	TransferredToCommunity []ModTransferCommunityView `json:"transferred_to_community" url:"transferred_to_community,omitempty"`
	Added                  []ModAddView               `json:"added" url:"added,omitempty"`
	AdminPurgedPersons     []AdminPurgePersonView     `json:"admin_purged_persons" url:"admin_purged_persons,omitempty"`
	AdminPurgedCommunities []AdminPurgeCommunityView  `json:"admin_purged_communities" url:"admin_purged_communities,omitempty"`
	AdminPurgedPosts       []AdminPurgePostView       `json:"admin_purged_posts" url:"admin_purged_posts,omitempty"`
	AdminPurgedComments    []AdminPurgeCommentView    `json:"admin_purged_comments" url:"admin_purged_comments,omitempty"`
	HiddenCommunities      []ModHideCommunityView     `json:"hidden_communities" url:"hidden_communities,omitempty"`
	LemmyResponse
}

type CreateSite struct {
	Name                       string             `json:"name" url:"name,omitempty"`
	Sidebar                    Optional[string]   `json:"sidebar" url:"sidebar,omitempty"`
	Description                Optional[string]   `json:"description" url:"description,omitempty"`
	Icon                       Optional[string]   `json:"icon" url:"icon,omitempty"`
	Banner                     Optional[string]   `json:"banner" url:"banner,omitempty"`
	EnableDownvotes            Optional[bool]     `json:"enable_downvotes" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]     `json:"open_registration" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]     `json:"enable_nsfw" url:"enable_nsfw,omitempty"`
	CommunityCreationAdminOnly Optional[bool]     `json:"community_creation_admin_only" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]     `json:"require_email_verification" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]     `json:"require_application" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string]   `json:"application_question" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]     `json:"private_instance" url:"private_instance,omitempty"`
	DefaultTheme               Optional[string]   `json:"default_theme" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string]   `json:"default_post_listing_type" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string]   `json:"legal_information" url:"legal_information,omitempty"`
	ApplicationEmailAdmins     Optional[bool]     `json:"application_email_admins" url:"application_email_admins,omitempty"`
	HideModlogModNames         Optional[bool]     `json:"hide_modlog_mod_names" url:"hide_modlog_mod_names,omitempty"`
	DiscussionLanguages        Optional[[]int]    `json:"discussion_languages" url:"discussion_languages,omitempty"`
	SlurFilterRegex            Optional[string]   `json:"slur_filter_regex" url:"slur_filter_regex,omitempty"`
	ActorNameMaxLength         Optional[int]      `json:"actor_name_max_length" url:"actor_name_max_length,omitempty"`
	RateLimitMessage           Optional[int]      `json:"rate_limit_message" url:"rate_limit_message,omitempty"`
	RateLimitMessagePerSecond  Optional[int]      `json:"rate_limit_message_per_second" url:"rate_limit_message_per_second,omitempty"`
	RateLimitPost              Optional[int]      `json:"rate_limit_post" url:"rate_limit_post,omitempty"`
	RateLimitPostPerSecond     Optional[int]      `json:"rate_limit_post_per_second" url:"rate_limit_post_per_second,omitempty"`
	RateLimitRegister          Optional[int]      `json:"rate_limit_register" url:"rate_limit_register,omitempty"`
	RateLimitRegisterPerSecond Optional[int]      `json:"rate_limit_register_per_second" url:"rate_limit_register_per_second,omitempty"`
	RateLimitImage             Optional[int]      `json:"rate_limit_image" url:"rate_limit_image,omitempty"`
	RateLimitImagePerSecond    Optional[int]      `json:"rate_limit_image_per_second" url:"rate_limit_image_per_second,omitempty"`
	RateLimitComment           Optional[int]      `json:"rate_limit_comment" url:"rate_limit_comment,omitempty"`
	RateLimitCommentPerSecond  Optional[int]      `json:"rate_limit_comment_per_second" url:"rate_limit_comment_per_second,omitempty"`
	RateLimitSearch            Optional[int]      `json:"rate_limit_search" url:"rate_limit_search,omitempty"`
	RateLimitSearchPerSecond   Optional[int]      `json:"rate_limit_search_per_second" url:"rate_limit_search_per_second,omitempty"`
	FederationEnabled          Optional[bool]     `json:"federation_enabled" url:"federation_enabled,omitempty"`
	FederationDebug            Optional[bool]     `json:"federation_debug" url:"federation_debug,omitempty"`
	FederationWorkerCount      Optional[int]      `json:"federation_worker_count" url:"federation_worker_count,omitempty"`
	CaptchaEnabled             Optional[bool]     `json:"captcha_enabled" url:"captcha_enabled,omitempty"`
	CaptchaDifficulty          Optional[string]   `json:"captcha_difficulty" url:"captcha_difficulty,omitempty"`
	AllowedInstances           Optional[[]string] `json:"allowed_instances" url:"allowed_instances,omitempty"`
	BlockedInstances           Optional[[]string] `json:"blocked_instances" url:"blocked_instances,omitempty"`
	Auth                       string             `json:"auth" url:"auth,omitempty"`
}

type EditSite struct {
	Name                       Optional[string]   `json:"name" url:"name,omitempty"`
	Sidebar                    Optional[string]   `json:"sidebar" url:"sidebar,omitempty"`
	Description                Optional[string]   `json:"description" url:"description,omitempty"`
	Icon                       Optional[string]   `json:"icon" url:"icon,omitempty"`
	Banner                     Optional[string]   `json:"banner" url:"banner,omitempty"`
	EnableDownvotes            Optional[bool]     `json:"enable_downvotes" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]     `json:"open_registration" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]     `json:"enable_nsfw" url:"enable_nsfw,omitempty"`
	CommunityCreationAdminOnly Optional[bool]     `json:"community_creation_admin_only" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]     `json:"require_email_verification" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]     `json:"require_application" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string]   `json:"application_question" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]     `json:"private_instance" url:"private_instance,omitempty"`
	DefaultTheme               Optional[string]   `json:"default_theme" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string]   `json:"default_post_listing_type" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string]   `json:"legal_information" url:"legal_information,omitempty"`
	ApplicationEmailAdmins     Optional[bool]     `json:"application_email_admins" url:"application_email_admins,omitempty"`
	HideModlogModNames         Optional[bool]     `json:"hide_modlog_mod_names" url:"hide_modlog_mod_names,omitempty"`
	DiscussionLanguages        Optional[[]int]    `json:"discussion_languages" url:"discussion_languages,omitempty"`
	SlurFilterRegex            Optional[string]   `json:"slur_filter_regex" url:"slur_filter_regex,omitempty"`
	ActorNameMaxLength         Optional[int]      `json:"actor_name_max_length" url:"actor_name_max_length,omitempty"`
	RateLimitMessage           Optional[int]      `json:"rate_limit_message" url:"rate_limit_message,omitempty"`
	RateLimitMessagePerSecond  Optional[int]      `json:"rate_limit_message_per_second" url:"rate_limit_message_per_second,omitempty"`
	RateLimitPost              Optional[int]      `json:"rate_limit_post" url:"rate_limit_post,omitempty"`
	RateLimitPostPerSecond     Optional[int]      `json:"rate_limit_post_per_second" url:"rate_limit_post_per_second,omitempty"`
	RateLimitRegister          Optional[int]      `json:"rate_limit_register" url:"rate_limit_register,omitempty"`
	RateLimitRegisterPerSecond Optional[int]      `json:"rate_limit_register_per_second" url:"rate_limit_register_per_second,omitempty"`
	RateLimitImage             Optional[int]      `json:"rate_limit_image" url:"rate_limit_image,omitempty"`
	RateLimitImagePerSecond    Optional[int]      `json:"rate_limit_image_per_second" url:"rate_limit_image_per_second,omitempty"`
	RateLimitComment           Optional[int]      `json:"rate_limit_comment" url:"rate_limit_comment,omitempty"`
	RateLimitCommentPerSecond  Optional[int]      `json:"rate_limit_comment_per_second" url:"rate_limit_comment_per_second,omitempty"`
	RateLimitSearch            Optional[int]      `json:"rate_limit_search" url:"rate_limit_search,omitempty"`
	RateLimitSearchPerSecond   Optional[int]      `json:"rate_limit_search_per_second" url:"rate_limit_search_per_second,omitempty"`
	FederationEnabled          Optional[bool]     `json:"federation_enabled" url:"federation_enabled,omitempty"`
	FederationDebug            Optional[bool]     `json:"federation_debug" url:"federation_debug,omitempty"`
	FederationWorkerCount      Optional[int]      `json:"federation_worker_count" url:"federation_worker_count,omitempty"`
	CaptchaEnabled             Optional[bool]     `json:"captcha_enabled" url:"captcha_enabled,omitempty"`
	CaptchaDifficulty          Optional[string]   `json:"captcha_difficulty" url:"captcha_difficulty,omitempty"`
	AllowedInstances           Optional[[]string] `json:"allowed_instances" url:"allowed_instances,omitempty"`
	BlockedInstances           Optional[[]string] `json:"blocked_instances" url:"blocked_instances,omitempty"`
	Taglines                   Optional[[]string] `json:"taglines" url:"taglines,omitempty"`
	Auth                       string             `json:"auth" url:"auth,omitempty"`
}

type GetSite struct {
	Auth Optional[string] `json:"auth" url:"auth,omitempty"`
}

type SiteResponse struct {
	SiteView SiteView `json:"site_view" url:"site_view,omitempty"`
	LemmyResponse
}

type GetSiteResponse struct {
	SiteView            SiteView                     `json:"site_view" url:"site_view,omitempty"`
	Admins              []PersonViewSafe             `json:"admins" url:"admins,omitempty"`
	Online              int                          `json:"online" url:"online,omitempty"`
	Version             string                       `json:"version" url:"version,omitempty"`
	MyUser              Optional[MyUserInfo]         `json:"my_user" url:"my_user,omitempty"`
	FederatedInstances  Optional[FederatedInstances] `json:"federated_instances" url:"federated_instances,omitempty"`
	AllLanguages        []Language                   `json:"all_languages" url:"all_languages,omitempty"`
	DiscussionLanguages []int                        `json:"discussion_languages" url:"discussion_languages,omitempty"`
	Taglines            Optional[[]Tagline]          `json:"taglines" url:"taglines,omitempty"`
	LemmyResponse
}

type MyUserInfo struct {
	LocalUserView       LocalUserSettingsView    `json:"local_user_view" url:"local_user_view,omitempty"`
	Follows             []CommunityFollowerView  `json:"follows" url:"follows,omitempty"`
	Moderates           []CommunityModeratorView `json:"moderates" url:"moderates,omitempty"`
	CommunityBlocks     []CommunityBlockView     `json:"community_blocks" url:"community_blocks,omitempty"`
	PersonBlocks        []PersonBlockView        `json:"person_blocks" url:"person_blocks,omitempty"`
	DiscussionLanguages []Language               `json:"discussion_languages" url:"discussion_languages,omitempty"`
}

type LeaveAdmin struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}

type FederatedInstances struct {
	Linked  []string           `json:"linked" url:"linked,omitempty"`
	Allowed Optional[[]string] `json:"allowed" url:"allowed,omitempty"`
	Blocked Optional[[]string] `json:"blocked" url:"blocked,omitempty"`
}

type PurgePerson struct {
	PersonID int              `json:"person_id" url:"person_id,omitempty"`
	Reason   Optional[string] `json:"reason" url:"reason,omitempty"`
	Auth     string           `json:"auth" url:"auth,omitempty"`
}

type PurgeCommunity struct {
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Auth        string           `json:"auth" url:"auth,omitempty"`
}

type PurgePost struct {
	PostID int              `json:"post_id" url:"post_id,omitempty"`
	Reason Optional[string] `json:"reason" url:"reason,omitempty"`
	Auth   string           `json:"auth" url:"auth,omitempty"`
}

type PurgeComment struct {
	CommentID int              `json:"comment_id" url:"comment_id,omitempty"`
	Reason    Optional[string] `json:"reason" url:"reason,omitempty"`
	Auth      string           `json:"auth" url:"auth,omitempty"`
}

type PurgeItemResponse struct {
	Success bool `json:"success" url:"success,omitempty"`
	LemmyResponse
}

type ListRegistrationApplications struct {
	UnreadOnly Optional[bool] `json:"unread_only" url:"unread_only,omitempty"`
	Page       Optional[int]  `json:"page" url:"page,omitempty"`
	Limit      Optional[int]  `json:"limit" url:"limit,omitempty"`
	Auth       string         `json:"auth" url:"auth,omitempty"`
}

type ListRegistrationApplicationsResponse struct {
	RegistrationApplications []RegistrationApplicationView `json:"registration_applications" url:"registration_applications,omitempty"`
	LemmyResponse
}

type ApproveRegistrationApplication struct {
	ID         int              `json:"id" url:"id,omitempty"`
	Approve    bool             `json:"approve" url:"approve,omitempty"`
	DenyReason Optional[string] `json:"deny_reason" url:"deny_reason,omitempty"`
	Auth       string           `json:"auth" url:"auth,omitempty"`
}

type RegistrationApplicationResponse struct {
	RegistrationApplication RegistrationApplicationView `json:"registration_application" url:"registration_application,omitempty"`
	LemmyResponse
}

type GetUnreadRegistrationApplicationCount struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}

type GetUnreadRegistrationApplicationCountResponse struct {
	RegistrationApplications int `json:"registration_applications" url:"registration_applications,omitempty"`
	LemmyResponse
}
