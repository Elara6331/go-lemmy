package types

import "time"

type LocalUserSettings struct {
	ID                  int              `json:"id" url:"id,omitempty"`
	PersonID            int              `json:"person_id" url:"person_id,omitempty"`
	Email               Optional[string] `json:"email" url:"email,omitempty"`
	ShowNSFW            bool             `json:"show_nsfw" url:"show_nsfw,omitempty"`
	Theme               string           `json:"theme" url:"theme,omitempty"`
	DefaultSortType     int              `json:"default_sort_type" url:"default_sort_type,omitempty"`
	DefaultListingType  int              `json:"default_listing_type" url:"default_listing_type,omitempty"`
	InterfaceLanguage   string           `json:"interface_language" url:"interface_language,omitempty"`
	ShowAvatars         bool             `json:"show_avatars" url:"show_avatars,omitempty"`
	SendNotifications   bool             `json:"send_notifications_to_email" url:"send_notifications_to_email,omitempty"`
	ValidatorTime       string           `json:"validator_time" url:"validator_time,omitempty"`
	ShowBotAccounts     bool             `json:"show_bot_accounts" url:"show_bot_accounts,omitempty"`
	ShowScores          bool             `json:"show_scores" url:"show_scores,omitempty"`
	ShowReadPosts       bool             `json:"show_read_posts" url:"show_read_posts,omitempty"`
	ShowNewPostNotifs   bool             `json:"show_new_post_notifs" url:"show_new_post_notifs,omitempty"`
	EmailVerified       bool             `json:"email_verified" url:"email_verified,omitempty"`
	AcceptedApplication bool             `json:"accepted_application" url:"accepted_application,omitempty"`
}

type PersonSafe struct {
	ID             int              `json:"id" url:"id,omitempty"`
	Name           string           `json:"name" url:"name,omitempty"`
	DisplayName    Optional[string] `json:"display_name" url:"display_name,omitempty"`
	Avatar         Optional[string] `json:"avatar" url:"avatar,omitempty"`
	Banned         bool             `json:"banned" url:"banned,omitempty"`
	Published      string           `json:"published" url:"published,omitempty"`
	Updated        Optional[string] `json:"updated" url:"updated,omitempty"`
	ActorID        string           `json:"actor_id" url:"actor_id,omitempty"`
	Bio            Optional[string] `json:"bio" url:"bio,omitempty"`
	Local          bool             `json:"local" url:"local,omitempty"`
	Banner         Optional[string] `json:"banner" url:"banner,omitempty"`
	Deleted        bool             `json:"deleted" url:"deleted,omitempty"`
	InboxURL       string           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL Optional[string] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	MatrixUserID   Optional[string] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	Admin          bool             `json:"admin" url:"admin,omitempty"`
	BotAccount     bool             `json:"bot_account" url:"bot_account,omitempty"`
	BanExpires     Optional[string] `json:"ban_expires" url:"ban_expires,omitempty"`
	InstanceID     int              `json:"instance_id" url:"instance_id,omitempty"`
}

type Site struct {
	ID              int              `json:"id" url:"id,omitempty"`
	Name            string           `json:"name" url:"name,omitempty"`
	Sidebar         Optional[string] `json:"sidebar" url:"sidebar,omitempty"`
	Published       string           `json:"published" url:"published,omitempty"`
	Updated         Optional[string] `json:"updated" url:"updated,omitempty"`
	Icon            Optional[string] `json:"icon" url:"icon,omitempty"`
	Banner          Optional[string] `json:"banner" url:"banner,omitempty"`
	Description     Optional[string] `json:"description" url:"description,omitempty"`
	ActorID         string           `json:"actor_id" url:"actor_id,omitempty"`
	LastRefreshedAt string           `json:"last_refreshed_at" url:"last_refreshed_at,omitempty"`
	InboxURL        string           `json:"inbox_url" url:"inbox_url,omitempty"`
	PrivateKey      Optional[string] `json:"private_key" url:"private_key,omitempty"`
	PublicKey       string           `json:"public_key" url:"public_key,omitempty"`
	InstanceID      int              `json:"instance_id" url:"instance_id,omitempty"`
}

type LocalSite struct {
	ID                         int              `json:"id" url:"id,omitempty"`
	SiteID                     int              `json:"site_id" url:"site_id,omitempty"`
	SiteSetup                  bool             `json:"site_setup" url:"site_setup,omitempty"`
	EnableDownvotes            bool             `json:"enable_downvotes" url:"enable_downvotes,omitempty"`
	OpenRegistration           bool             `json:"open_registration" url:"open_registration,omitempty"`
	EnableNSFW                 bool             `json:"enable_nsfw" url:"enable_nsfw,omitempty"`
	AdminOnlyCommunityCreation bool             `json:"community_creation_admin_only" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   bool             `json:"require_email_verification" url:"require_email_verification,omitempty"`
	RequireApplication         bool             `json:"require_application" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string] `json:"application_question" url:"application_question,omitempty"`
	PrivateInstance            bool             `json:"private_instance" url:"private_instance,omitempty"`
	DefaultTheme               string           `json:"default_theme" url:"default_theme,omitempty"`
	DefaultPostListingType     string           `json:"default_post_listing_type" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string] `json:"legal_information" url:"legal_information,omitempty"`
	HideModlogModNames         bool             `json:"hide_modlog_mod_names" url:"hide_modlog_mod_names,omitempty"`
	ApplicationEmailAdmins     bool             `json:"application_email_admins" url:"application_email_admins,omitempty"`
	SlurFilterRegex            Optional[string] `json:"slur_filter_regex" url:"slur_filter_regex,omitempty"`
	ActorNameMaxLength         int              `json:"actor_name_max_length" url:"actor_name_max_length,omitempty"`
	FederationEnabled          bool             `json:"federation_enabled" url:"federation_enabled,omitempty"`
	FederationDebug            bool             `json:"federation_debug" url:"federation_debug,omitempty"`
	FederationStrictAllowlist  bool             `json:"federation_strict_allowlist" url:"federation_strict_allowlist,omitempty"`
	FederationRetryLimit       int              `json:"federation_http_fetch_retry_limit" url:"federation_http_fetch_retry_limit,omitempty"`
	FederationWorkerCount      int              `json:"federation_worker_count" url:"federation_worker_count,omitempty"`
	CaptchaEnabled             bool             `json:"captcha_enabled" url:"captcha_enabled,omitempty"`
	CaptchaDifficulty          string           `json:"captcha_difficulty" url:"captcha_difficulty,omitempty"`
	Published                  string           `json:"published" url:"published,omitempty"`
	Updated                    Optional[string] `json:"updated" url:"updated,omitempty"`
}

type LocalSiteRateLimit struct {
	ID                int              `json:"id" url:"id,omitempty"`
	LocalSiteID       int              `json:"local_site_id" url:"local_site_id,omitempty"`
	Message           int              `json:"message" url:"message,omitempty"`
	MessagePerSecond  int              `json:"message_per_second" url:"message_per_second,omitempty"`
	Post              int              `json:"post" url:"post,omitempty"`
	PostPerSecond     int              `json:"post_per_second" url:"post_per_second,omitempty"`
	Register          int              `json:"register" url:"register,omitempty"`
	RegisterPerSecond int              `json:"register_per_second" url:"register_per_second,omitempty"`
	Image             int              `json:"image" url:"image,omitempty"`
	ImagePerSecond    int              `json:"image_per_second" url:"image_per_second,omitempty"`
	Comment           int              `json:"comment" url:"comment,omitempty"`
	CommentPerSecond  int              `json:"comment_per_second" url:"comment_per_second,omitempty"`
	Search            int              `json:"search" url:"search,omitempty"`
	SearchPerSecond   int              `json:"search_per_second" url:"search_per_second,omitempty"`
	Published         string           `json:"published" url:"published,omitempty"`
	Updated           Optional[string] `json:"updated" url:"updated,omitempty"`
}

type PrivateMessage struct {
	ID          int              `json:"id" url:"id,omitempty"`
	CreatorID   int              `json:"creator_id" url:"creator_id,omitempty"`
	RecipientID int              `json:"recipient_id" url:"recipient_id,omitempty"`
	Content     string           `json:"content" url:"content,omitempty"`
	Deleted     bool             `json:"deleted" url:"deleted,omitempty"`
	Read        bool             `json:"read" url:"read,omitempty"`
	Published   string           `json:"published" url:"published,omitempty"`
	Updated     Optional[string] `json:"updated" url:"updated,omitempty"`
	ApID        string           `json:"ap_id" url:"ap_id,omitempty"`
	Local       bool             `json:"local" url:"local,omitempty"`
}

type PostReport struct {
	ID               int              `json:"id" url:"id,omitempty"`
	CreatorID        int              `json:"creator_id" url:"creator_id,omitempty"`
	PostID           int              `json:"post_id" url:"post_id,omitempty"`
	OriginalPostName string           `json:"original_post_name" url:"original_post_name,omitempty"`
	OriginalPostURL  Optional[string] `json:"original_post_url" url:"original_post_url,omitempty"`
	OriginalPostBody Optional[string] `json:"original_post_body" url:"original_post_body,omitempty"`
	Reason           string           `json:"reason" url:"reason,omitempty"`
	Resolved         bool             `json:"resolved" url:"resolved,omitempty"`
	ResolverID       Optional[int]    `json:"resolver_id" url:"resolver_id,omitempty"`
	Published        string           `json:"published" url:"published,omitempty"`
	Updated          Optional[string] `json:"updated" url:"updated,omitempty"`
}

type Post struct {
	ID               int              `json:"id" url:"id,omitempty"`
	Name             string           `json:"name" url:"name,omitempty"`
	URL              Optional[string] `json:"url" url:"url,omitempty"`
	Body             Optional[string] `json:"body" url:"body,omitempty"`
	CreatorID        int              `json:"creator_id" url:"creator_id,omitempty"`
	CommunityID      int              `json:"community_id" url:"community_id,omitempty"`
	Removed          bool             `json:"removed" url:"removed,omitempty"`
	Locked           bool             `json:"locked" url:"locked,omitempty"`
	Published        string           `json:"published" url:"published,omitempty"`
	Updated          Optional[string] `json:"updated" url:"updated,omitempty"`
	Deleted          bool             `json:"deleted" url:"deleted,omitempty"`
	NSFW             bool             `json:"nsfw" url:"nsfw,omitempty"`
	Stickied         bool             `json:"stickied" url:"stickied,omitempty"`
	EmbedTitle       Optional[string] `json:"embed_title" url:"embed_title,omitempty"`
	EmbedDescription Optional[string] `json:"embed_description" url:"embed_description,omitempty"`
	EmbedVideoURL    Optional[string] `json:"embed_video_url" url:"embed_video_url,omitempty"`
	ThumbnailURL     Optional[string] `json:"thumbnail_url" url:"thumbnail_url,omitempty"`
	ApID             string           `json:"ap_id" url:"ap_id,omitempty"`
	Local            bool             `json:"local" url:"local,omitempty"`
	LanguageID       int              `json:"language_id" url:"language_id,omitempty"`
}

type PasswordResetRequest struct {
	ID             int    `json:"id" url:"id,omitempty"`
	LocalUserID    int    `json:"local_user_id" url:"local_user_id,omitempty"`
	TokenEncrypted string `json:"token_encrypted" url:"token_encrypted,omitempty"`
	Published      string `json:"published" url:"published,omitempty"`
}

type ModRemovePost struct {
	ID          int              `json:"id" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int              `json:"post_id" url:"post_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
	When        string           `json:"when_" url:"when_,omitempty"`
}

type ModLockPost struct {
	ID          int            `json:"id" url:"id,omitempty"`
	ModPersonID int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id" url:"post_id,omitempty"`
	Locked      Optional[bool] `json:"locked" url:"locked,omitempty"`
	When        string         `json:"when_" url:"when_,omitempty"`
}

// ModStickyPost represents a post stickied by a moderator.
type ModStickyPost struct {
	ID          int            `json:"id" url:"id,omitempty"`
	ModPersonID int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id" url:"post_id,omitempty"`
	Stickied    Optional[bool] `json:"stickied" url:"stickied,omitempty"`
	When        string         `json:"when_" url:"when_,omitempty"`
}

// ModRemoveComment represents a comment removal by a moderator.
type ModRemoveComment struct {
	ID          int              `json:"id" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	CommentID   int              `json:"comment_id" url:"comment_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
	When        string           `json:"when_" url:"when_,omitempty"`
}

// ModRemoveCommunity represents a community removal by a moderator.
type ModRemoveCommunity struct {
	ID          int              `json:"id" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
	Expires     Optional[string] `json:"expires" url:"expires,omitempty"`
	When        string           `json:"when_" url:"when_,omitempty"`
}

// ModBanFromCommunity represents a user being banned from a community by a moderator.
type ModBanFromCommunity struct {
	ID            int              `json:"id" url:"id,omitempty"`
	ModPersonID   int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int              `json:"community_id" url:"community_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned" url:"banned,omitempty"`
	Expires       Optional[string] `json:"expires" url:"expires,omitempty"`
	When          string           `json:"when_" url:"when_,omitempty"`
}

// ModBan represents a user being banned by a moderator.
type ModBan struct {
	ID            int              `json:"id" url:"id,omitempty"`
	ModPersonID   int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id" url:"other_person_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned" url:"banned,omitempty"`
	Expires       Optional[string] `json:"expires" url:"expires,omitempty"`
	When          string           `json:"when_" url:"when_,omitempty"`
}

// ModAddCommunity represents a user being added as a moderator of a community.
type ModAddCommunity struct {
	ID            int            `json:"id" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
	When          string         `json:"when_" url:"when_,omitempty"`
}

// ModTransferCommunity represents a community being transferred to another moderator.
type ModTransferCommunity struct {
	ID            int            `json:"id" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
	When          string         `json:"when_" url:"when_,omitempty"`
}

// ModAdd represents a user being added as a moderator.
type ModAdd struct {
	ID            int            `json:"id" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
	When          string         `json:"when_" url:"when_,omitempty"`
}

type AdminPurgePerson struct {
	ID            int              `json:"id" url:"id,omitempty"`
	AdminPersonID int              `json:"admin_person_id" url:"admin_person_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	When          string           `json:"when_" url:"when_,omitempty"`
}

type AdminPurgeCommunity struct {
	ID            int              `json:"id" url:"id,omitempty"`
	AdminPersonID int              `json:"admin_person_id" url:"admin_person_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	When          string           `json:"when_" url:"when_,omitempty"`
}

type AdminPurgePost struct {
	ID            int              `json:"id" url:"id,omitempty"`
	AdminPersonID int              `json:"admin_person_id" url:"admin_person_id,omitempty"`
	CommunityID   int              `json:"community_id" url:"community_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	When          string           `json:"when_" url:"when_,omitempty"`
}

type AdminPurgeComment struct {
	ID            int              `json:"id" url:"id,omitempty"`
	AdminPersonID int              `json:"admin_person_id" url:"admin_person_id,omitempty"`
	PostID        int              `json:"post_id" url:"post_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	When          string           `json:"when_" url:"when_,omitempty"`
}

type CommunitySafe struct {
	ID                      int              `json:"id" url:"id,omitempty"`
	Name                    string           `json:"name" url:"name,omitempty"`
	Title                   string           `json:"title" url:"title,omitempty"`
	Description             Optional[string] `json:"description" url:"description,omitempty"`
	Removed                 bool             `json:"removed" url:"removed,omitempty"`
	Published               string           `json:"published" url:"published,omitempty"`
	Updated                 Optional[string] `json:"updated" url:"updated,omitempty"`
	Deleted                 bool             `json:"deleted" url:"deleted,omitempty"`
	NSFW                    bool             `json:"nsfw" url:"nsfw,omitempty"`
	ActorID                 string           `json:"actor_id" url:"actor_id,omitempty"`
	Local                   bool             `json:"local" url:"local,omitempty"`
	Icon                    Optional[string] `json:"icon" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner" url:"banner,omitempty"`
	Hidden                  bool             `json:"hidden" url:"hidden,omitempty"`
	PostingRestrictedToMods bool             `json:"posting_restricted_to_mods" url:"posting_restricted_to_mods,omitempty"`
	InstanceID              int              `json:"instance_id" url:"instance_id,omitempty"`
}

type CommentReport struct {
	ID                  int              `json:"id" url:"id,omitempty"`
	CreatorID           int              `json:"creator_id" url:"creator_id,omitempty"`
	CommentID           int              `json:"comment_id" url:"comment_id,omitempty"`
	OriginalCommentText string           `json:"original_comment_text" url:"original_comment_text,omitempty"`
	Reason              string           `json:"reason" url:"reason,omitempty"`
	Resolved            bool             `json:"resolved" url:"resolved,omitempty"`
	ResolverID          Optional[int]    `json:"resolver_id" url:"resolver_id,omitempty"`
	Published           string           `json:"published" url:"published,omitempty"`
	Updated             Optional[string] `json:"updated" url:"updated,omitempty"`
}

type Comment struct {
	ID            int              `json:"id" url:"id,omitempty"`
	CreatorID     int              `json:"creator_id" url:"creator_id,omitempty"`
	PostID        int              `json:"post_id" url:"post_id,omitempty"`
	Content       string           `json:"content" url:"content,omitempty"`
	Removed       bool             `json:"removed" url:"removed,omitempty"`
	Published     string           `json:"published" url:"published,omitempty"`
	Updated       Optional[string] `json:"updated" url:"updated,omitempty"`
	Deleted       bool             `json:"deleted" url:"deleted,omitempty"`
	APID          string           `json:"ap_id" url:"ap_id,omitempty"`
	Local         bool             `json:"local" url:"local,omitempty"`
	Path          string           `json:"path" url:"path,omitempty"`
	Distinguished bool             `json:"distinguished" url:"distinguished,omitempty"`
	LanguageID    int              `json:"language_id" url:"language_id,omitempty"`
}

type PersonMention struct {
	ID          int    `json:"id" url:"id,omitempty"`
	RecipientID int    `json:"recipient_id" url:"recipient_id,omitempty"`
	CommentID   int    `json:"comment_id" url:"comment_id,omitempty"`
	Read        bool   `json:"read" url:"read,omitempty"`
	Published   string `json:"published" url:"published,omitempty"`
}

type Language struct {
	ID   int    `json:"id" url:"id,omitempty"`
	Code string `json:"code" url:"code,omitempty"`
	Name string `json:"name" url:"name,omitempty"`
}

type PrivateMessageReport struct {
	ID               int              `json:"id" url:"id,omitempty"`
	CreatorID        int              `json:"creator_id" url:"creator_id,omitempty"`
	PrivateMessageID int              `json:"private_message_id" url:"private_message_id,omitempty"`
	OriginalPMText   string           `json:"original_pm_text" url:"original_pm_text,omitempty"`
	Reason           string           `json:"reason" url:"reason,omitempty"`
	Resolved         bool             `json:"resolved" url:"resolved,omitempty"`
	ResolverID       Optional[int]    `json:"resolver_id" url:"resolver_id,omitempty"`
	Published        string           `json:"published" url:"published,omitempty"`
	Updated          Optional[string] `json:"updated" url:"updated,omitempty"`
}

type RegistrationApplication struct {
	ID          int              `json:"id" url:"id,omitempty"`
	LocalUserID int              `json:"local_user_id" url:"local_user_id,omitempty"`
	Answer      string           `json:"answer" url:"answer,omitempty"`
	AdminID     Optional[int]    `json:"admin_id" url:"admin_id,omitempty"`
	DenyReason  Optional[string] `json:"deny_reason" url:"deny_reason,omitempty"`
	Published   string           `json:"published" url:"published,omitempty"`
}

type ModHideCommunityView struct {
	ModHideCommunity ModHideCommunity     `json:"mod_hide_community" url:"mod_hide_community,omitempty"`
	Admin            Optional[PersonSafe] `json:"admin" url:"admin,omitempty"`
	Community        CommunitySafe        `json:"community" url:"community,omitempty"`
}

type ModHideCommunity struct {
	ID          int              `json:"id" url:"id,omitempty"`
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Hidden      Optional[bool]   `json:"hidden" url:"hidden,omitempty"`
	When        time.Time        `json:"when_" url:"when_,omitempty"`
}

type CommentReply struct {
	ID          int    `json:"id" url:"id,omitempty"`
	RecipientID int    `json:"recipient_id" url:"recipient_id,omitempty"`
	CommentID   int    `json:"comment_id" url:"comment_id,omitempty"`
	Read        bool   `json:"read" url:"read,omitempty"`
	Published   string `json:"published" url:"published,omitempty"`
}
