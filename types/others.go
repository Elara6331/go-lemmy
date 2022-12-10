package types

/*type SiteMetadata struct {
	Title       Optional[string] `json:"title" url:"title,omitempty"`
	Description Optional[string] `json:"description" url:"description,omitempty"`
	Image       Optional[string] `json:"image" url:"image,omitempty"`
	HTML        Optional[string] `json:"html" url:"html,omitempty"`
}*/

type UserOperation int

const (
	LoginOp UserOperation = iota
	RegisterOp
	GetCaptchaOp
	CreateCommunityOp
	CreatePostOp
	ListCommunitiesOp
	GetPostOp
	GetCommunityOp
	CreateCommentOp
	EditCommentOp
	DeleteCommentOp
	RemoveCommentOp
	SaveCommentOp
	CreateCommentLikeOp
	GetPostsOp
	CreatePostLikeOp
	EditPostOp
	DeletePostOp
	RemovePostOp
	LockPostOp
	StickyPostOp
	MarkPostAsReadOp
	SavePostOp
	EditCommunityOp
	DeleteCommunityOp
	RemoveCommunityOp
	FollowCommunityOp
	GetPersonDetailsOp
	GetRepliesOp
	GetPersonMentionsOp
	MarkPersonMentionAsReadOp
	MarkCommentReplyAsReadOp
	GetModlogOp
	BanFromCommunityOp
	AddModToCommunityOp
	CreateSiteOp
	EditSiteOp
	GetSiteOp
	AddAdminOp
	GetUnreadRegistrationApplicationCountOp
	ListRegistrationApplicationsOp
	ApproveRegistrationApplicationOp
	BanPersonOp
	GetBannedPersonsOp
	SearchOp
	ResolveObjectOp
	MarkAllAsReadOp
	SaveUserSettingsOp
	TransferCommunityOp
	LeaveAdminOp
	DeleteAccountOp
	PasswordResetOp
	PasswordChangeOp
	CreatePrivateMessageOp
	EditPrivateMessageOp
	DeletePrivateMessageOp
	MarkPrivateMessageAsReadOp
	CreatePrivateMessageReportOp
	ResolvePrivateMessageReportOp
	ListPrivateMessageReportsOp
	GetPrivateMessagesOp
	UserJoinOp
	GetCommentsOp
	PostJoinOp
	CommunityJoinOp
	ChangePasswordOp
	GetSiteMetadataOp
	BlockCommunityOp
	BlockPersonOp
	PurgePersonOp
	PurgeCommunityOp
	PurgePostOp
	PurgeCommentOp
	CreateCommentReportOp
	ResolveCommentReportOp
	ListCommentReportsOp
	CreatePostReportOp
	ResolvePostReportOp
	ListPostReportsOp
	GetReportCountOp
	GetUnreadCountOp
	VerifyEmailOp
)

type SortType string

const (
	Active       SortType = "Active"
	Hot          SortType = "Hot"
	New          SortType = "New"
	Old          SortType = "Old"
	TopDay       SortType = "TopDay"
	TopWeek      SortType = "TopWeek"
	TopMonth     SortType = "TopMonth"
	TopYear      SortType = "TopYear"
	TopAll       SortType = "TopAll"
	MostComments SortType = "MostComments"
	NewComments  SortType = "NewComments"
)

type CommentSortType string

const (
	CommentSortHot CommentSortType = "Hot"
	CommentSortTop CommentSortType = "Top"
	CommentSortNew CommentSortType = "New"
	CommentSortOld CommentSortType = "Old"
)

type ListingType string

const (
	ListingAll        ListingType = "All"
	ListingLocal      ListingType = "Local"
	ListingSubscribed ListingType = "Subscribed"
	ListingCommunity  ListingType = "Community"
)

type SearchType string

const (
	SearchAll         SearchType = "All"
	SearchComments    SearchType = "Comments"
	SearchPosts       SearchType = "Posts"
	SearchCommunities SearchType = "Communities"
	SearchUsers       SearchType = "Users"
	SearchURL         SearchType = "URL"
)

type ModlogActionType string

const (
	ModlogAll                  ModlogActionType = "All"
	ModlogModRemovePost        ModlogActionType = "ModRemovePost"
	ModlogModLockPost          ModlogActionType = "ModLockPost"
	ModlogModStickyPost        ModlogActionType = "ModStickyPost"
	ModlogModRemoveComment     ModlogActionType = "ModRemoveComment"
	ModlogModRemoveCommunity   ModlogActionType = "ModRemoveCommunity"
	ModlogModBanFromCommunity  ModlogActionType = "ModBanFromCommunity"
	ModlogModAddCommunity      ModlogActionType = "ModAddCommunity"
	ModlogModTransferCommunity ModlogActionType = "ModTransferCommunity"
	ModlogModAdd               ModlogActionType = "ModAdd"
	ModlogModBan               ModlogActionType = "ModBan"
	ModlogModHideCommunity     ModlogActionType = "ModHideCommunity"
	ModlogAdminPurgePerson     ModlogActionType = "AdminPurgePerson"
	ModlogAdminPurgeCommunity  ModlogActionType = "AdminPurgeCommunity"
	ModlogAdminPurgePost       ModlogActionType = "AdminPurgePost"
	ModlogAdminPurgeComment    ModlogActionType = "AdminPurgeComment"
)

type SubscribedType string

const (
	Subscribed    SubscribedType = "Subscribed"
	NotSubscribed SubscribedType = "NotSubscribed"
	Pending       SubscribedType = "Pending"
)
