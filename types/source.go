package types

import "time"

type LocalUserSettings struct {
	ID                  int              `json:"id,omitempty" url:"id,omitempty"`
	PersonID            int              `json:"person_id,omitempty" url:"person_id,omitempty"`
	Email               Optional[string] `json:"email,omitempty" url:"email,omitempty"`
	ShowNSFW            bool             `json:"show_nsfw,omitempty" url:"show_nsfw,omitempty"`
	Theme               string           `json:"theme,omitempty" url:"theme,omitempty"`
	DefaultSortType     int              `json:"default_sort_type,omitempty" url:"default_sort_type,omitempty"`
	DefaultListingType  int              `json:"default_listing_type,omitempty" url:"default_listing_type,omitempty"`
	InterfaceLanguage   string           `json:"interface_language,omitempty" url:"interface_language,omitempty"`
	ShowAvatars         bool             `json:"show_avatars,omitempty" url:"show_avatars,omitempty"`
	SendNotifications   bool             `json:"send_notifications_to_email,omitempty" url:"send_notifications_to_email,omitempty"`
	ValidatorTime       string           `json:"validator_time,omitempty" url:"validator_time,omitempty"`
	ShowBotAccounts     bool             `json:"show_bot_accounts,omitempty" url:"show_bot_accounts,omitempty"`
	ShowScores          bool             `json:"show_scores,omitempty" url:"show_scores,omitempty"`
	ShowReadPosts       bool             `json:"show_read_posts,omitempty" url:"show_read_posts,omitempty"`
	ShowNewPostNotifs   bool             `json:"show_new_post_notifs,omitempty" url:"show_new_post_notifs,omitempty"`
	EmailVerified       bool             `json:"email_verified,omitempty" url:"email_verified,omitempty"`
	AcceptedApplication bool             `json:"accepted_application,omitempty" url:"accepted_application,omitempty"`
}

type PersonSafe struct {
	ID             int              `json:"id,omitempty" url:"id,omitempty"`
	Name           string           `json:"name,omitempty" url:"name,omitempty"`
	DisplayName    Optional[string] `json:"display_name,omitempty" url:"display_name,omitempty"`
	Avatar         Optional[string] `json:"avatar,omitempty" url:"avatar,omitempty"`
	Banned         bool             `json:"banned,omitempty" url:"banned,omitempty"`
	Published      string           `json:"published,omitempty" url:"published,omitempty"`
	Updated        Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
	ActorID        string           `json:"actor_id,omitempty" url:"actor_id,omitempty"`
	Bio            Optional[string] `json:"bio,omitempty" url:"bio,omitempty"`
	Local          bool             `json:"local,omitempty" url:"local,omitempty"`
	Banner         Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	Deleted        bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	InboxURL       string           `json:"inbox_url,omitempty" url:"inbox_url,omitempty"`
	SharedInboxURL Optional[string] `json:"shared_inbox_url,omitempty" url:"shared_inbox_url,omitempty"`
	MatrixUserID   Optional[string] `json:"matrix_user_id,omitempty" url:"matrix_user_id,omitempty"`
	Admin          bool             `json:"admin,omitempty" url:"admin,omitempty"`
	BotAccount     bool             `json:"bot_account,omitempty" url:"bot_account,omitempty"`
	BanExpires     Optional[string] `json:"ban_expires,omitempty" url:"ban_expires,omitempty"`
	InstanceID     int              `json:"instance_id,omitempty" url:"instance_id,omitempty"`
}

type Site struct {
	ID              int              `json:"id,omitempty" url:"id,omitempty"`
	Name            string           `json:"name,omitempty" url:"name,omitempty"`
	Sidebar         Optional[string] `json:"sidebar,omitempty" url:"sidebar,omitempty"`
	Published       string           `json:"published,omitempty" url:"published,omitempty"`
	Updated         Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
	Icon            Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner          Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	Description     Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	ActorID         string           `json:"actor_id,omitempty" url:"actor_id,omitempty"`
	LastRefreshedAt string           `json:"last_refreshed_at,omitempty" url:"last_refreshed_at,omitempty"`
	InboxURL        string           `json:"inbox_url,omitempty" url:"inbox_url,omitempty"`
	PrivateKey      Optional[string] `json:"private_key,omitempty" url:"private_key,omitempty"`
	PublicKey       string           `json:"public_key,omitempty" url:"public_key,omitempty"`
	InstanceID      int              `json:"instance_id,omitempty" url:"instance_id,omitempty"`
}

type LocalSite struct {
	ID                         int              `json:"id,omitempty" url:"id,omitempty"`
	SiteID                     int              `json:"site_id,omitempty" url:"site_id,omitempty"`
	SiteSetup                  bool             `json:"site_setup,omitempty" url:"site_setup,omitempty"`
	EnableDownvotes            bool             `json:"enable_downvotes,omitempty" url:"enable_downvotes,omitempty"`
	OpenRegistration           bool             `json:"open_registration,omitempty" url:"open_registration,omitempty"`
	EnableNSFW                 bool             `json:"enable_nsfw,omitempty" url:"enable_nsfw,omitempty"`
	AdminOnlyCommunityCreation bool             `json:"community_creation_admin_only,omitempty" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   bool             `json:"require_email_verification,omitempty" url:"require_email_verification,omitempty"`
	RequireApplication         bool             `json:"require_application,omitempty" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string] `json:"application_question,omitempty" url:"application_question,omitempty"`
	PrivateInstance            bool             `json:"private_instance,omitempty" url:"private_instance,omitempty"`
	DefaultTheme               string           `json:"default_theme,omitempty" url:"default_theme,omitempty"`
	DefaultPostListingType     string           `json:"default_post_listing_type,omitempty" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string] `json:"legal_information,omitempty" url:"legal_information,omitempty"`
	HideModlogModNames         bool             `json:"hide_modlog_mod_names,omitempty" url:"hide_modlog_mod_names,omitempty"`
	ApplicationEmailAdmins     bool             `json:"application_email_admins,omitempty" url:"application_email_admins,omitempty"`
	SlurFilterRegex            Optional[string] `json:"slur_filter_regex,omitempty" url:"slur_filter_regex,omitempty"`
	ActorNameMaxLength         int              `json:"actor_name_max_length,omitempty" url:"actor_name_max_length,omitempty"`
	FederationEnabled          bool             `json:"federation_enabled,omitempty" url:"federation_enabled,omitempty"`
	FederationDebug            bool             `json:"federation_debug,omitempty" url:"federation_debug,omitempty"`
	FederationStrictAllowlist  bool             `json:"federation_strict_allowlist,omitempty" url:"federation_strict_allowlist,omitempty"`
	FederationRetryLimit       int              `json:"federation_http_fetch_retry_limit,omitempty" url:"federation_http_fetch_retry_limit,omitempty"`
	FederationWorkerCount      int              `json:"federation_worker_count,omitempty" url:"federation_worker_count,omitempty"`
	CaptchaEnabled             bool             `json:"captcha_enabled,omitempty" url:"captcha_enabled,omitempty"`
	CaptchaDifficulty          string           `json:"captcha_difficulty,omitempty" url:"captcha_difficulty,omitempty"`
	Published                  string           `json:"published,omitempty" url:"published,omitempty"`
	Updated                    Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
}

type LocalSiteRateLimit struct {
	ID                int              `json:"id,omitempty" url:"id,omitempty"`
	LocalSiteID       int              `json:"local_site_id,omitempty" url:"local_site_id,omitempty"`
	Message           int              `json:"message,omitempty" url:"message,omitempty"`
	MessagePerSecond  int              `json:"message_per_second,omitempty" url:"message_per_second,omitempty"`
	Post              int              `json:"post,omitempty" url:"post,omitempty"`
	PostPerSecond     int              `json:"post_per_second,omitempty" url:"post_per_second,omitempty"`
	Register          int              `json:"register,omitempty" url:"register,omitempty"`
	RegisterPerSecond int              `json:"register_per_second,omitempty" url:"register_per_second,omitempty"`
	Image             int              `json:"image,omitempty" url:"image,omitempty"`
	ImagePerSecond    int              `json:"image_per_second,omitempty" url:"image_per_second,omitempty"`
	Comment           int              `json:"comment,omitempty" url:"comment,omitempty"`
	CommentPerSecond  int              `json:"comment_per_second,omitempty" url:"comment_per_second,omitempty"`
	Search            int              `json:"search,omitempty" url:"search,omitempty"`
	SearchPerSecond   int              `json:"search_per_second,omitempty" url:"search_per_second,omitempty"`
	Published         string           `json:"published,omitempty" url:"published,omitempty"`
	Updated           Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
}

type PrivateMessage struct {
	ID          int              `json:"id,omitempty" url:"id,omitempty"`
	CreatorID   int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	RecipientID int              `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	Content     string           `json:"content,omitempty" url:"content,omitempty"`
	Deleted     bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	Read        bool             `json:"read,omitempty" url:"read,omitempty"`
	Published   string           `json:"published,omitempty" url:"published,omitempty"`
	Updated     Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
	ApID        string           `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local       bool             `json:"local,omitempty" url:"local,omitempty"`
}

type PostReport struct {
	ID               int              `json:"id,omitempty" url:"id,omitempty"`
	CreatorID        int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	PostID           int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	OriginalPostName string           `json:"original_post_name,omitempty" url:"original_post_name,omitempty"`
	OriginalPostURL  Optional[string] `json:"original_post_url,omitempty" url:"original_post_url,omitempty"`
	OriginalPostBody Optional[string] `json:"original_post_body,omitempty" url:"original_post_body,omitempty"`
	Reason           string           `json:"reason,omitempty" url:"reason,omitempty"`
	Resolved         bool             `json:"resolved,omitempty" url:"resolved,omitempty"`
	ResolverID       Optional[int]    `json:"resolver_id,omitempty" url:"resolver_id,omitempty"`
	Published        string           `json:"published,omitempty" url:"published,omitempty"`
	Updated          Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
}

type Post struct {
	ID               int              `json:"id,omitempty" url:"id,omitempty"`
	Name             string           `json:"name,omitempty" url:"name,omitempty"`
	URL              Optional[string] `json:"url,omitempty" url:"url,omitempty"`
	Body             Optional[string] `json:"body,omitempty" url:"body,omitempty"`
	CreatorID        int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	CommunityID      int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed          bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Locked           bool             `json:"locked,omitempty" url:"locked,omitempty"`
	Published        string           `json:"published,omitempty" url:"published,omitempty"`
	Updated          Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted          bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	NSFW             bool             `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	Stickied         bool             `json:"stickied,omitempty" url:"stickied,omitempty"`
	EmbedTitle       Optional[string] `json:"embed_title,omitempty" url:"embed_title,omitempty"`
	EmbedDescription Optional[string] `json:"embed_description,omitempty" url:"embed_description,omitempty"`
	EmbedVideoURL    Optional[string] `json:"embed_video_url,omitempty" url:"embed_video_url,omitempty"`
	ThumbnailURL     Optional[string] `json:"thumbnail_url,omitempty" url:"thumbnail_url,omitempty"`
	ApID             string           `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local            bool             `json:"local,omitempty" url:"local,omitempty"`
	LanguageID       int              `json:"language_id,omitempty" url:"language_id,omitempty"`
}

type PasswordResetRequest struct {
	ID             int    `json:"id,omitempty" url:"id,omitempty"`
	LocalUserID    int    `json:"local_user_id,omitempty" url:"local_user_id,omitempty"`
	TokenEncrypted string `json:"token_encrypted,omitempty" url:"token_encrypted,omitempty"`
	Published      string `json:"published,omitempty" url:"published,omitempty"`
}

type ModRemovePost struct {
	ID          int              `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	When        string           `json:"when_,omitempty" url:"when_,omitempty"`
}

type ModLockPost struct {
	ID          int            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id,omitempty" url:"post_id,omitempty"`
	Locked      Optional[bool] `json:"locked,omitempty" url:"locked,omitempty"`
	When        string         `json:"when_,omitempty" url:"when_,omitempty"`
}

// ModStickyPost represents a post stickied by a moderator.
type ModStickyPost struct {
	ID          int            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id,omitempty" url:"post_id,omitempty"`
	Stickied    Optional[bool] `json:"stickied,omitempty" url:"stickied,omitempty"`
	When        string         `json:"when_,omitempty" url:"when_,omitempty"`
}

// ModRemoveComment represents a comment removal by a moderator.
type ModRemoveComment struct {
	ID          int              `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	CommentID   int              `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	When        string           `json:"when_,omitempty" url:"when_,omitempty"`
}

// ModRemoveCommunity represents a community removal by a moderator.
type ModRemoveCommunity struct {
	ID          int              `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	Expires     Optional[string] `json:"expires,omitempty" url:"expires,omitempty"`
	When        string           `json:"when_,omitempty" url:"when_,omitempty"`
}

// ModBanFromCommunity represents a user being banned from a community by a moderator.
type ModBanFromCommunity struct {
	ID            int              `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned,omitempty" url:"banned,omitempty"`
	Expires       Optional[string] `json:"expires,omitempty" url:"expires,omitempty"`
	When          string           `json:"when_,omitempty" url:"when_,omitempty"`
}

// ModBan represents a user being banned by a moderator.
type ModBan struct {
	ID            int              `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned,omitempty" url:"banned,omitempty"`
	Expires       Optional[string] `json:"expires,omitempty" url:"expires,omitempty"`
	When          string           `json:"when_,omitempty" url:"when_,omitempty"`
}

// ModAddCommunity represents a user being added as a moderator of a community.
type ModAddCommunity struct {
	ID            int            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
	When          string         `json:"when_,omitempty" url:"when_,omitempty"`
}

// ModTransferCommunity represents a community being transferred to another moderator.
type ModTransferCommunity struct {
	ID            int            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
	When          string         `json:"when_,omitempty" url:"when_,omitempty"`
}

// ModAdd represents a user being added as a moderator.
type ModAdd struct {
	ID            int            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
	When          string         `json:"when_,omitempty" url:"when_,omitempty"`
}

type AdminPurgePerson struct {
	ID            int              `json:"id,omitempty" url:"id,omitempty"`
	AdminPersonID int              `json:"admin_person_id,omitempty" url:"admin_person_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	When          string           `json:"when_,omitempty" url:"when_,omitempty"`
}

type AdminPurgeCommunity struct {
	ID            int              `json:"id,omitempty" url:"id,omitempty"`
	AdminPersonID int              `json:"admin_person_id,omitempty" url:"admin_person_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	When          string           `json:"when_,omitempty" url:"when_,omitempty"`
}

type AdminPurgePost struct {
	ID            int              `json:"id,omitempty" url:"id,omitempty"`
	AdminPersonID int              `json:"admin_person_id,omitempty" url:"admin_person_id,omitempty"`
	CommunityID   int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	When          string           `json:"when_,omitempty" url:"when_,omitempty"`
}

type AdminPurgeComment struct {
	ID            int              `json:"id,omitempty" url:"id,omitempty"`
	AdminPersonID int              `json:"admin_person_id,omitempty" url:"admin_person_id,omitempty"`
	PostID        int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	When          string           `json:"when_,omitempty" url:"when_,omitempty"`
}

type CommunitySafe struct {
	ID                      int              `json:"id,omitempty" url:"id,omitempty"`
	Name                    string           `json:"name,omitempty" url:"name,omitempty"`
	Title                   string           `json:"title,omitempty" url:"title,omitempty"`
	Description             Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	Removed                 bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Published               string           `json:"published,omitempty" url:"published,omitempty"`
	Updated                 Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted                 bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	NSFW                    bool             `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	ActorID                 string           `json:"actor_id,omitempty" url:"actor_id,omitempty"`
	Local                   bool             `json:"local,omitempty" url:"local,omitempty"`
	Icon                    Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	Hidden                  bool             `json:"hidden,omitempty" url:"hidden,omitempty"`
	PostingRestrictedToMods bool             `json:"posting_restricted_to_mods,omitempty" url:"posting_restricted_to_mods,omitempty"`
	InstanceID              int              `json:"instance_id,omitempty" url:"instance_id,omitempty"`
}

type CommentReport struct {
	ID                  int              `json:"id,omitempty" url:"id,omitempty"`
	CreatorID           int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	CommentID           int              `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	OriginalCommentText string           `json:"original_comment_text,omitempty" url:"original_comment_text,omitempty"`
	Reason              string           `json:"reason,omitempty" url:"reason,omitempty"`
	Resolved            bool             `json:"resolved,omitempty" url:"resolved,omitempty"`
	ResolverID          Optional[int]    `json:"resolver_id,omitempty" url:"resolver_id,omitempty"`
	Published           string           `json:"published,omitempty" url:"published,omitempty"`
	Updated             Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
}

type Comment struct {
	ID            int              `json:"id,omitempty" url:"id,omitempty"`
	CreatorID     int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	PostID        int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Content       string           `json:"content,omitempty" url:"content,omitempty"`
	Removed       bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Published     string           `json:"published,omitempty" url:"published,omitempty"`
	Updated       Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted       bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	APID          string           `json:"ap_id,omitempty" url:"ap_id,omitempty"`
	Local         bool             `json:"local,omitempty" url:"local,omitempty"`
	Path          string           `json:"path,omitempty" url:"path,omitempty"`
	Distinguished bool             `json:"distinguished,omitempty" url:"distinguished,omitempty"`
	LanguageID    int              `json:"language_id,omitempty" url:"language_id,omitempty"`
}

type PersonMention struct {
	ID          int    `json:"id,omitempty" url:"id,omitempty"`
	RecipientID int    `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	CommentID   int    `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Read        bool   `json:"read,omitempty" url:"read,omitempty"`
	Published   string `json:"published,omitempty" url:"published,omitempty"`
}

type Language struct {
	ID   int    `json:"id,omitempty" url:"id,omitempty"`
	Code string `json:"code,omitempty" url:"code,omitempty"`
	Name string `json:"name,omitempty" url:"name,omitempty"`
}

type PrivateMessageReport struct {
	ID               int              `json:"id,omitempty" url:"id,omitempty"`
	CreatorID        int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	PrivateMessageID int              `json:"private_message_id,omitempty" url:"private_message_id,omitempty"`
	OriginalPMText   string           `json:"original_pm_text,omitempty" url:"original_pm_text,omitempty"`
	Reason           string           `json:"reason,omitempty" url:"reason,omitempty"`
	Resolved         bool             `json:"resolved,omitempty" url:"resolved,omitempty"`
	ResolverID       Optional[int]    `json:"resolver_id,omitempty" url:"resolver_id,omitempty"`
	Published        string           `json:"published,omitempty" url:"published,omitempty"`
	Updated          Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
}

type RegistrationApplication struct {
	ID          int              `json:"id,omitempty" url:"id,omitempty"`
	LocalUserID int              `json:"local_user_id,omitempty" url:"local_user_id,omitempty"`
	Answer      string           `json:"answer,omitempty" url:"answer,omitempty"`
	AdminID     Optional[int]    `json:"admin_id,omitempty" url:"admin_id,omitempty"`
	DenyReason  Optional[string] `json:"deny_reason,omitempty" url:"deny_reason,omitempty"`
	Published   string           `json:"published,omitempty" url:"published,omitempty"`
}

type ModHideCommunityView struct {
	ModHideCommunity ModHideCommunity     `json:"mod_hide_community,omitempty" url:"mod_hide_community,omitempty"`
	Admin            Optional[PersonSafe] `json:"admin,omitempty" url:"admin,omitempty"`
	Community        CommunitySafe        `json:"community,omitempty" url:"community,omitempty"`
}

type ModHideCommunity struct {
	ID          int              `json:"id,omitempty" url:"id,omitempty"`
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Hidden      Optional[bool]   `json:"hidden,omitempty" url:"hidden,omitempty"`
	When        time.Time        `json:"when_,omitempty" url:"when_,omitempty"`
}

type CommentReply struct {
	ID          int    `json:"id,omitempty" url:"id,omitempty"`
	RecipientID int    `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	CommentID   int    `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Read        bool   `json:"read,omitempty" url:"read,omitempty"`
	Published   string `json:"published,omitempty" url:"published,omitempty"`
}
