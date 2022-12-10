package types

type SearchResponse struct {
	Type        string           `json:"type,omitempty" url:"type,omitempty"`
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
	ModPersonID   Optional[int]              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	CommunityID   Optional[int]              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Page          Optional[int64]            `json:"page,omitempty" url:"page,omitempty"`
	Limit         Optional[int64]            `json:"limit,omitempty" url:"limit,omitempty"`
	Auth          Optional[string]           `json:"auth,omitempty" url:"auth,omitempty"`
	Type          Optional[ModlogActionType] `json:"type,omitempty" url:"type,omitempty"`
	OtherPersonID Optional[int]              `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
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
	AdminPurgedPersons     []AdminPurgePersonView     `json:"admin_purged_persons,omitempty" url:"admin_purged_persons,omitempty"`
	AdminPurgedCommunities []AdminPurgeCommunityView  `json:"admin_purged_communities,omitempty" url:"admin_purged_communities,omitempty"`
	AdminPurgedPosts       []AdminPurgePostView       `json:"admin_purged_posts,omitempty" url:"admin_purged_posts,omitempty"`
	AdminPurgedComments    []AdminPurgeCommentView    `json:"admin_purged_comments,omitempty" url:"admin_purged_comments,omitempty"`
	HiddenCommunities      []ModHideCommunityView     `json:"hidden_communities,omitempty" url:"hidden_communities,omitempty"`
	LemmyResponse
}

type CreateSite struct {
	Name                       string             `json:"name,omitempty" url:"name,omitempty"`
	Sidebar                    Optional[string]   `json:"sidebar,omitempty" url:"sidebar,omitempty"`
	Description                Optional[string]   `json:"description,omitempty" url:"description,omitempty"`
	Icon                       Optional[string]   `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                     Optional[string]   `json:"banner,omitempty" url:"banner,omitempty"`
	EnableDownvotes            Optional[bool]     `json:"enable_downvotes,omitempty" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]     `json:"open_registration,omitempty" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]     `json:"enable_nsfw,omitempty" url:"enable_nsfw,omitempty"`
	CommunityCreationAdminOnly Optional[bool]     `json:"community_creation_admin_only,omitempty" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]     `json:"require_email_verification,omitempty" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]     `json:"require_application,omitempty" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string]   `json:"application_question,omitempty" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]     `json:"private_instance,omitempty" url:"private_instance,omitempty"`
	DefaultTheme               Optional[string]   `json:"default_theme,omitempty" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string]   `json:"default_post_listing_type,omitempty" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string]   `json:"legal_information,omitempty" url:"legal_information,omitempty"`
	ApplicationEmailAdmins     Optional[bool]     `json:"application_email_admins,omitempty" url:"application_email_admins,omitempty"`
	HideModlogModNames         Optional[bool]     `json:"hide_modlog_mod_names,omitempty" url:"hide_modlog_mod_names,omitempty"`
	DiscussionLanguages        Optional[[]int]    `json:"discussion_languages,omitempty" url:"discussion_languages,omitempty"`
	SlurFilterRegex            Optional[string]   `json:"slur_filter_regex,omitempty" url:"slur_filter_regex,omitempty"`
	ActorNameMaxLength         Optional[int]      `json:"actor_name_max_length,omitempty" url:"actor_name_max_length,omitempty"`
	RateLimitMessage           Optional[int]      `json:"rate_limit_message,omitempty" url:"rate_limit_message,omitempty"`
	RateLimitMessagePerSecond  Optional[int]      `json:"rate_limit_message_per_second,omitempty" url:"rate_limit_message_per_second,omitempty"`
	RateLimitPost              Optional[int]      `json:"rate_limit_post,omitempty" url:"rate_limit_post,omitempty"`
	RateLimitPostPerSecond     Optional[int]      `json:"rate_limit_post_per_second,omitempty" url:"rate_limit_post_per_second,omitempty"`
	RateLimitRegister          Optional[int]      `json:"rate_limit_register,omitempty" url:"rate_limit_register,omitempty"`
	RateLimitRegisterPerSecond Optional[int]      `json:"rate_limit_register_per_second,omitempty" url:"rate_limit_register_per_second,omitempty"`
	RateLimitImage             Optional[int]      `json:"rate_limit_image,omitempty" url:"rate_limit_image,omitempty"`
	RateLimitImagePerSecond    Optional[int]      `json:"rate_limit_image_per_second,omitempty" url:"rate_limit_image_per_second,omitempty"`
	RateLimitComment           Optional[int]      `json:"rate_limit_comment,omitempty" url:"rate_limit_comment,omitempty"`
	RateLimitCommentPerSecond  Optional[int]      `json:"rate_limit_comment_per_second,omitempty" url:"rate_limit_comment_per_second,omitempty"`
	RateLimitSearch            Optional[int]      `json:"rate_limit_search,omitempty" url:"rate_limit_search,omitempty"`
	RateLimitSearchPerSecond   Optional[int]      `json:"rate_limit_search_per_second,omitempty" url:"rate_limit_search_per_second,omitempty"`
	FederationEnabled          Optional[bool]     `json:"federation_enabled,omitempty" url:"federation_enabled,omitempty"`
	FederationDebug            Optional[bool]     `json:"federation_debug,omitempty" url:"federation_debug,omitempty"`
	FederationWorkerCount      Optional[int]      `json:"federation_worker_count,omitempty" url:"federation_worker_count,omitempty"`
	CaptchaEnabled             Optional[bool]     `json:"captcha_enabled,omitempty" url:"captcha_enabled,omitempty"`
	CaptchaDifficulty          Optional[string]   `json:"captcha_difficulty,omitempty" url:"captcha_difficulty,omitempty"`
	AllowedInstances           Optional[[]string] `json:"allowed_instances,omitempty" url:"allowed_instances,omitempty"`
	BlockedInstances           Optional[[]string] `json:"blocked_instances,omitempty" url:"blocked_instances,omitempty"`
	Auth                       string             `json:"auth,omitempty" url:"auth,omitempty"`
}

type EditSite struct {
	Name                       Optional[string]   `json:"name,omitempty" url:"name,omitempty"`
	Sidebar                    Optional[string]   `json:"sidebar,omitempty" url:"sidebar,omitempty"`
	Description                Optional[string]   `json:"description,omitempty" url:"description,omitempty"`
	Icon                       Optional[string]   `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                     Optional[string]   `json:"banner,omitempty" url:"banner,omitempty"`
	EnableDownvotes            Optional[bool]     `json:"enable_downvotes,omitempty" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]     `json:"open_registration,omitempty" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]     `json:"enable_nsfw,omitempty" url:"enable_nsfw,omitempty"`
	CommunityCreationAdminOnly Optional[bool]     `json:"community_creation_admin_only,omitempty" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]     `json:"require_email_verification,omitempty" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]     `json:"require_application,omitempty" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string]   `json:"application_question,omitempty" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]     `json:"private_instance,omitempty" url:"private_instance,omitempty"`
	DefaultTheme               Optional[string]   `json:"default_theme,omitempty" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string]   `json:"default_post_listing_type,omitempty" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string]   `json:"legal_information,omitempty" url:"legal_information,omitempty"`
	ApplicationEmailAdmins     Optional[bool]     `json:"application_email_admins,omitempty" url:"application_email_admins,omitempty"`
	HideModlogModNames         Optional[bool]     `json:"hide_modlog_mod_names,omitempty" url:"hide_modlog_mod_names,omitempty"`
	DiscussionLanguages        Optional[[]int]    `json:"discussion_languages,omitempty" url:"discussion_languages,omitempty"`
	SlurFilterRegex            Optional[string]   `json:"slur_filter_regex,omitempty" url:"slur_filter_regex,omitempty"`
	ActorNameMaxLength         Optional[int]      `json:"actor_name_max_length,omitempty" url:"actor_name_max_length,omitempty"`
	RateLimitMessage           Optional[int]      `json:"rate_limit_message,omitempty" url:"rate_limit_message,omitempty"`
	RateLimitMessagePerSecond  Optional[int]      `json:"rate_limit_message_per_second,omitempty" url:"rate_limit_message_per_second,omitempty"`
	RateLimitPost              Optional[int]      `json:"rate_limit_post,omitempty" url:"rate_limit_post,omitempty"`
	RateLimitPostPerSecond     Optional[int]      `json:"rate_limit_post_per_second,omitempty" url:"rate_limit_post_per_second,omitempty"`
	RateLimitRegister          Optional[int]      `json:"rate_limit_register,omitempty" url:"rate_limit_register,omitempty"`
	RateLimitRegisterPerSecond Optional[int]      `json:"rate_limit_register_per_second,omitempty" url:"rate_limit_register_per_second,omitempty"`
	RateLimitImage             Optional[int]      `json:"rate_limit_image,omitempty" url:"rate_limit_image,omitempty"`
	RateLimitImagePerSecond    Optional[int]      `json:"rate_limit_image_per_second,omitempty" url:"rate_limit_image_per_second,omitempty"`
	RateLimitComment           Optional[int]      `json:"rate_limit_comment,omitempty" url:"rate_limit_comment,omitempty"`
	RateLimitCommentPerSecond  Optional[int]      `json:"rate_limit_comment_per_second,omitempty" url:"rate_limit_comment_per_second,omitempty"`
	RateLimitSearch            Optional[int]      `json:"rate_limit_search,omitempty" url:"rate_limit_search,omitempty"`
	RateLimitSearchPerSecond   Optional[int]      `json:"rate_limit_search_per_second,omitempty" url:"rate_limit_search_per_second,omitempty"`
	FederationEnabled          Optional[bool]     `json:"federation_enabled,omitempty" url:"federation_enabled,omitempty"`
	FederationDebug            Optional[bool]     `json:"federation_debug,omitempty" url:"federation_debug,omitempty"`
	FederationWorkerCount      Optional[int]      `json:"federation_worker_count,omitempty" url:"federation_worker_count,omitempty"`
	CaptchaEnabled             Optional[bool]     `json:"captcha_enabled,omitempty" url:"captcha_enabled,omitempty"`
	CaptchaDifficulty          Optional[string]   `json:"captcha_difficulty,omitempty" url:"captcha_difficulty,omitempty"`
	AllowedInstances           Optional[[]string] `json:"allowed_instances,omitempty" url:"allowed_instances,omitempty"`
	BlockedInstances           Optional[[]string] `json:"blocked_instances,omitempty" url:"blocked_instances,omitempty"`
	Taglines                   Optional[[]string] `json:"taglines,omitempty" url:"taglines,omitempty"`
	Auth                       string             `json:"auth,omitempty" url:"auth,omitempty"`
}

type GetSite struct {
	Auth Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type SiteResponse struct {
	SiteView SiteView `json:"site_view,omitempty" url:"site_view,omitempty"`
	LemmyResponse
}

type GetSiteResponse struct {
	SiteView            SiteView                     `json:"site_view,omitempty" url:"site_view,omitempty"`
	Admins              []PersonViewSafe             `json:"admins,omitempty" url:"admins,omitempty"`
	Online              int                          `json:"online,omitempty" url:"online,omitempty"`
	Version             string                       `json:"version,omitempty" url:"version,omitempty"`
	MyUser              Optional[MyUserInfo]         `json:"my_user,omitempty" url:"my_user,omitempty"`
	FederatedInstances  Optional[FederatedInstances] `json:"federated_instances,omitempty" url:"federated_instances,omitempty"`
	AllLanguages        []Language                   `json:"all_languages,omitempty" url:"all_languages,omitempty"`
	DiscussionLanguages []int                        `json:"discussion_languages,omitempty" url:"discussion_languages,omitempty"`
	Taglines            Optional[[]Tagline]          `json:"taglines,omitempty" url:"taglines,omitempty"`
	LemmyResponse
}

type MyUserInfo struct {
	LocalUserView       LocalUserSettingsView    `json:"local_user_view,omitempty" url:"local_user_view,omitempty"`
	Follows             []CommunityFollowerView  `json:"follows,omitempty" url:"follows,omitempty"`
	Moderates           []CommunityModeratorView `json:"moderates,omitempty" url:"moderates,omitempty"`
	CommunityBlocks     []CommunityBlockView     `json:"community_blocks,omitempty" url:"community_blocks,omitempty"`
	PersonBlocks        []PersonBlockView        `json:"person_blocks,omitempty" url:"person_blocks,omitempty"`
	DiscussionLanguages []Language               `json:"discussion_languages,omitempty" url:"discussion_languages,omitempty"`
}

type LeaveAdmin struct {
	Auth string `json:"auth,omitempty" url:"auth,omitempty"`
}

type FederatedInstances struct {
	Linked  []string           `json:"linked,omitempty" url:"linked,omitempty"`
	Allowed Optional[[]string] `json:"allowed,omitempty" url:"allowed,omitempty"`
	Blocked Optional[[]string] `json:"blocked,omitempty" url:"blocked,omitempty"`
}

type PurgePerson struct {
	PersonID int              `json:"person_id,omitempty" url:"person_id,omitempty"`
	Reason   Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Auth     string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type PurgeCommunity struct {
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Auth        string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type PurgePost struct {
	PostID int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Reason Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Auth   string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type PurgeComment struct {
	CommentID int              `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Reason    Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Auth      string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type PurgeItemResponse struct {
	Success bool `json:"success,omitempty" url:"success,omitempty"`
	LemmyResponse
}

type ListRegistrationApplications struct {
	UnreadOnly Optional[bool] `json:"unread_only,omitempty" url:"unread_only,omitempty"`
	Page       Optional[int]  `json:"page,omitempty" url:"page,omitempty"`
	Limit      Optional[int]  `json:"limit,omitempty" url:"limit,omitempty"`
	Auth       string         `json:"auth,omitempty" url:"auth,omitempty"`
}

type ListRegistrationApplicationsResponse struct {
	RegistrationApplications []RegistrationApplicationView `json:"registration_applications,omitempty" url:"registration_applications,omitempty"`
	LemmyResponse
}

type ApproveRegistrationApplication struct {
	ID         int              `json:"id,omitempty" url:"id,omitempty"`
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
	RegistrationApplications int `json:"registration_applications,omitempty" url:"registration_applications,omitempty"`
	LemmyResponse
}
