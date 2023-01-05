package types

type Login struct {
	UsernameOrEmail string `json:"username_or_email,omitempty" url:"username_or_email,omitempty"`
	Password        string `json:"password,omitempty" url:"password,omitempty"`
}
type Register struct {
	Username       string           `json:"username,omitempty" url:"username,omitempty"`
	Password       string           `json:"password,omitempty" url:"password,omitempty"`
	PasswordVerify string           `json:"password_verify,omitempty" url:"password_verify,omitempty"`
	ShowNSFW       bool             `json:"show_nsfw,omitempty" url:"show_nsfw,omitempty"`
	Email          Optional[string] `json:"email,omitempty" url:"email,omitempty"`
	CaptchaUuid    Optional[string] `json:"captcha_uuid,omitempty" url:"captcha_uuid,omitempty"`
	CaptchaAnswer  Optional[string] `json:"captcha_answer,omitempty" url:"captcha_answer,omitempty"`
	Honeypot       Optional[string] `json:"honeypot,omitempty" url:"honeypot,omitempty"`
	Answer         Optional[string] `json:"answer,omitempty" url:"answer,omitempty"`
}
type GetCaptcha struct{}
type GetCaptchaResponse struct {
	Ok Optional[CaptchaResponse] `json:"ok,omitempty" url:"ok,omitempty"`
	LemmyResponse
}
type CaptchaResponse struct {
	Png  string `json:"png,omitempty" url:"png,omitempty"`
	Wav  string `json:"wav,omitempty" url:"wav,omitempty"`
	Uuid string `json:"uuid,omitempty" url:"uuid,omitempty"`
	LemmyResponse
}
type SaveUserSettings struct {
	ShowNSFW                 Optional[bool]   `json:"show_nsfw,omitempty" url:"show_nsfw,omitempty"`
	ShowScores               Optional[bool]   `json:"show_scores,omitempty" url:"show_scores,omitempty"`
	Theme                    Optional[string] `json:"theme,omitempty" url:"theme,omitempty"`
	DefaultSortType          Optional[int16]  `json:"default_sort_type,omitempty" url:"default_sort_type,omitempty"`
	DefaultListingType       Optional[int16]  `json:"default_listing_type,omitempty" url:"default_listing_type,omitempty"`
	Lang                     Optional[string] `json:"lang,omitempty" url:"lang,omitempty"`
	Avatar                   Optional[string] `json:"avatar,omitempty" url:"avatar,omitempty"`
	Banner                   Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	DisplayName              Optional[string] `json:"display_name,omitempty" url:"display_name,omitempty"`
	Email                    Optional[string] `json:"email,omitempty" url:"email,omitempty"`
	Bio                      Optional[string] `json:"bio,omitempty" url:"bio,omitempty"`
	MatrixUserID             Optional[string] `json:"matrix_user_id,omitempty" url:"matrix_user_id,omitempty"`
	ShowAvatars              Optional[bool]   `json:"show_avatars,omitempty" url:"show_avatars,omitempty"`
	SendNotificationsToEmail Optional[bool]   `json:"send_notifications_to_email,omitempty" url:"send_notifications_to_email,omitempty"`
	BotAccount               Optional[bool]   `json:"bot_account,omitempty" url:"bot_account,omitempty"`
	ShowBotAccounts          Optional[bool]   `json:"show_bot_accounts,omitempty" url:"show_bot_accounts,omitempty"`
	ShowReadPosts            Optional[bool]   `json:"show_read_posts,omitempty" url:"show_read_posts,omitempty"`
	ShowNewPostNotifs        Optional[bool]   `json:"show_new_post_notifs,omitempty" url:"show_new_post_notifs,omitempty"`
	Auth                     string           `json:"auth,omitempty" url:"auth,omitempty"`
}
type ChangePassword struct {
	NewPassword       string `json:"new_password,omitempty" url:"new_password,omitempty"`
	NewPasswordVerify string `json:"new_password_verify,omitempty" url:"new_password_verify,omitempty"`
	OldPassword       string `json:"old_password,omitempty" url:"old_password,omitempty"`
	Auth              string `json:"auth,omitempty" url:"auth,omitempty"`
}
type LoginResponse struct {
	Jwt                 Optional[string] `json:"jwt,omitempty" url:"jwt,omitempty"`
	RegistrationCreated bool             `json:"registration_created,omitempty" url:"registration_created,omitempty"`
	VerifyEmailSent     bool             `json:"verify_email_sent,omitempty" url:"verify_email_sent,omitempty"`
	LemmyResponse
}
type GetPersonDetails struct {
	PersonID    Optional[int]      `json:"person_id,omitempty" url:"person_id,omitempty"`
	Username    Optional[string]   `json:"username,omitempty" url:"username,omitempty"`
	Sort        Optional[SortType] `json:"sort,omitempty" url:"sort,omitempty"`
	Page        Optional[int64]    `json:"page,omitempty" url:"page,omitempty"`
	Limit       Optional[int64]    `json:"limit,omitempty" url:"limit,omitempty"`
	CommunityID Optional[int]      `json:"community_id,omitempty" url:"community_id,omitempty"`
	SavedOnly   Optional[bool]     `json:"saved_only,omitempty" url:"saved_only,omitempty"`
	Auth        Optional[string]   `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetPersonDetailsResponse struct {
	PersonView PersonViewSafe           `json:"person_view,omitempty" url:"person_view,omitempty"`
	Comments   []CommentView            `json:"comments,omitempty" url:"comments,omitempty"`
	Posts      []PostView               `json:"posts,omitempty" url:"posts,omitempty"`
	Moderates  []CommunityModeratorView `json:"moderates,omitempty" url:"moderates,omitempty"`
	LemmyResponse
}
type GetRepliesResponse struct {
	Replies []CommentView `json:"replies,omitempty" url:"replies,omitempty"`
	LemmyResponse
}
type GetPersonMentionsResponse struct {
	Mentions []PersonMentionView `json:"mentions,omitempty" url:"mentions,omitempty"`
	LemmyResponse
}
type MarkAllAsRead struct {
	Auth string `json:"auth,omitempty" url:"auth,omitempty"`
}
type AddAdmin struct {
	PersonID int    `json:"person_id,omitempty" url:"person_id,omitempty"`
	Added    bool   `json:"added,omitempty" url:"added,omitempty"`
	Auth     string `json:"auth,omitempty" url:"auth,omitempty"`
}
type AddAdminResponse struct {
	Admins []PersonViewSafe `json:"admins,omitempty" url:"admins,omitempty"`
	LemmyResponse
}
type BanPerson struct {
	PersonID   int              `json:"person_id,omitempty" url:"person_id,omitempty"`
	Ban        bool             `json:"ban,omitempty" url:"ban,omitempty"`
	RemoveData Optional[bool]   `json:"remove_data,omitempty" url:"remove_data,omitempty"`
	Reason     Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Expires    Optional[int64]  `json:"expires,omitempty" url:"expires,omitempty"`
	Auth       string           `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetBannedPersons struct {
	Auth string `json:"auth,omitempty" url:"auth,omitempty"`
}
type BannedPersonsResponse struct {
	Banned []PersonViewSafe `json:"banned,omitempty" url:"banned,omitempty"`
	LemmyResponse
}
type BanPersonResponse struct {
	PersonView PersonViewSafe `json:"person_view,omitempty" url:"person_view,omitempty"`
	Banned     bool           `json:"banned,omitempty" url:"banned,omitempty"`
	LemmyResponse
}
type BlockPerson struct {
	PersonID int    `json:"person_id,omitempty" url:"person_id,omitempty"`
	Block    bool   `json:"block,omitempty" url:"block,omitempty"`
	Auth     string `json:"auth,omitempty" url:"auth,omitempty"`
}
type BlockPersonResponse struct {
	PersonView PersonViewSafe `json:"person_view,omitempty" url:"person_view,omitempty"`
	Blocked    bool           `json:"blocked,omitempty" url:"blocked,omitempty"`
	LemmyResponse
}
type GetReplies struct {
	Sort       Optional[SortType] `json:"sort,omitempty" url:"sort,omitempty"`
	Page       Optional[int64]    `json:"page,omitempty" url:"page,omitempty"`
	Limit      Optional[int64]    `json:"limit,omitempty" url:"limit,omitempty"`
	UnreadOnly Optional[bool]     `json:"unread_only,omitempty" url:"unread_only,omitempty"`
	Auth       string             `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetPersonMentions struct {
	Sort       Optional[SortType] `json:"sort,omitempty" url:"sort,omitempty"`
	Page       Optional[int64]    `json:"page,omitempty" url:"page,omitempty"`
	Limit      Optional[int64]    `json:"limit,omitempty" url:"limit,omitempty"`
	UnreadOnly Optional[bool]     `json:"unread_only,omitempty" url:"unread_only,omitempty"`
	Auth       string             `json:"auth,omitempty" url:"auth,omitempty"`
}
type MarkPersonMentionAsRead struct {
	PersonMentionID int    `json:"person_mention_id,omitempty" url:"person_mention_id,omitempty"`
	Read            bool   `json:"read,omitempty" url:"read,omitempty"`
	Auth            string `json:"auth,omitempty" url:"auth,omitempty"`
}
type PersonMentionResponse struct {
	PersonMentionView PersonMentionView `json:"person_mention_view,omitempty" url:"person_mention_view,omitempty"`
	LemmyResponse
}
type DeleteAccount struct {
	Password string `json:"password,omitempty" url:"password,omitempty"`
	Auth     string `json:"auth,omitempty" url:"auth,omitempty"`
}
type DeleteAccountResponse struct {
	LemmyResponse
}
type PasswordReset struct {
	Email string `json:"email,omitempty" url:"email,omitempty"`
}
type PasswordResetResponse struct {
	LemmyResponse
}
type PasswordChangeAfterReset struct {
	Token          string `json:"token,omitempty" url:"token,omitempty"`
	Password       string `json:"password,omitempty" url:"password,omitempty"`
	PasswordVerify string `json:"password_verify,omitempty" url:"password_verify,omitempty"`
}
type CreatePrivateMessage struct {
	Content     string `json:"content,omitempty" url:"content,omitempty"`
	RecipientID int    `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	Auth        string `json:"auth,omitempty" url:"auth,omitempty"`
}
type EditPrivateMessage struct {
	PrivateMessageID int    `json:"private_message_id,omitempty" url:"private_message_id,omitempty"`
	Content          string `json:"content,omitempty" url:"content,omitempty"`
	Auth             string `json:"auth,omitempty" url:"auth,omitempty"`
}
type DeletePrivateMessage struct {
	PrivateMessageID int    `json:"private_message_id,omitempty" url:"private_message_id,omitempty"`
	Deleted          bool   `json:"deleted,omitempty" url:"deleted,omitempty"`
	Auth             string `json:"auth,omitempty" url:"auth,omitempty"`
}
type MarkPrivateMessageAsRead struct {
	PrivateMessageID int    `json:"private_message_id,omitempty" url:"private_message_id,omitempty"`
	Read             bool   `json:"read,omitempty" url:"read,omitempty"`
	Auth             string `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetPrivateMessages struct {
	UnreadOnly Optional[bool]  `json:"unread_only,omitempty" url:"unread_only,omitempty"`
	Page       Optional[int64] `json:"page,omitempty" url:"page,omitempty"`
	Limit      Optional[int64] `json:"limit,omitempty" url:"limit,omitempty"`
	Auth       string          `json:"auth,omitempty" url:"auth,omitempty"`
}
type PrivateMessagesResponse struct {
	PrivateMessages []PrivateMessageView `json:"private_messages,omitempty" url:"private_messages,omitempty"`
	LemmyResponse
}
type PrivateMessageResponse struct {
	PrivateMessageView PrivateMessageView `json:"private_message_view,omitempty" url:"private_message_view,omitempty"`
	LemmyResponse
}
type GetReportCount struct {
	CommunityID Optional[int] `json:"community_id,omitempty" url:"community_id,omitempty"`
	Auth        string        `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetReportCountResponse struct {
	CommunityID    Optional[int] `json:"community_id,omitempty" url:"community_id,omitempty"`
	CommentReports int64         `json:"comment_reports,omitempty" url:"comment_reports,omitempty"`
	PostReports    int64         `json:"post_reports,omitempty" url:"post_reports,omitempty"`
	LemmyResponse
}
type GetUnreadCount struct {
	Auth string `json:"auth,omitempty" url:"auth,omitempty"`
}
type GetUnreadCountResponse struct {
	Replies         int64 `json:"replies,omitempty" url:"replies,omitempty"`
	Mentions        int64 `json:"mentions,omitempty" url:"mentions,omitempty"`
	PrivateMessages int64 `json:"private_messages,omitempty" url:"private_messages,omitempty"`
	LemmyResponse
}
type VerifyEmail struct {
	Token string `json:"token,omitempty" url:"token,omitempty"`
}
type VerifyEmailResponse struct {
	LemmyResponse
}
