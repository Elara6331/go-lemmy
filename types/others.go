package types

type UserOperation string

const (
	UserOpLogin                                 UserOperation = "Login"
	UserOpRegister                              UserOperation = "Register"
	UserOpGetCaptcha                            UserOperation = "GetCaptcha"
	UserOpCreateCommunity                       UserOperation = "CreateCommunity"
	UserOpCreatePost                            UserOperation = "CreatePost"
	UserOpListCommunities                       UserOperation = "ListCommunities"
	UserOpGetPost                               UserOperation = "GetPost"
	UserOpGetCommunity                          UserOperation = "GetCommunity"
	UserOpCreateComment                         UserOperation = "CreateComment"
	UserOpEditComment                           UserOperation = "EditComment"
	UserOpDeleteComment                         UserOperation = "DeleteComment"
	UserOpRemoveComment                         UserOperation = "RemoveComment"
	UserOpSaveComment                           UserOperation = "SaveComment"
	UserOpCreateCommentLike                     UserOperation = "CreateCommentLike"
	UserOpGetPosts                              UserOperation = "GetPosts"
	UserOpCreatePostLike                        UserOperation = "CreatePostLike"
	UserOpEditPost                              UserOperation = "EditPost"
	UserOpDeletePost                            UserOperation = "DeletePost"
	UserOpRemovePost                            UserOperation = "RemovePost"
	UserOpLockPost                              UserOperation = "LockPost"
	UserOpFeaturePost                           UserOperation = "FeaturePost"
	UserOpMarkPostAsRead                        UserOperation = "MarkPostAsRead"
	UserOpSavePost                              UserOperation = "SavePost"
	UserOpEditCommunity                         UserOperation = "EditCommunity"
	UserOpDeleteCommunity                       UserOperation = "DeleteCommunity"
	UserOpRemoveCommunity                       UserOperation = "RemoveCommunity"
	UserOpFollowCommunity                       UserOperation = "FollowCommunity"
	UserOpGetPersonDetails                      UserOperation = "GetPersonDetails"
	UserOpGetReplies                            UserOperation = "GetReplies"
	UserOpGetPersonMentions                     UserOperation = "GetPersonMentions"
	UserOpMarkPersonMentionAsRead               UserOperation = "MarkPersonMentionAsRead"
	UserOpMarkCommentReplyAsRead                UserOperation = "MarkCommentReplyAsRead"
	UserOpGetModlog                             UserOperation = "GetModlog"
	UserOpBanFromCommunity                      UserOperation = "BanFromCommunity"
	UserOpAddModToCommunity                     UserOperation = "AddModToCommunity"
	UserOpCreateSite                            UserOperation = "CreateSite"
	UserOpEditSite                              UserOperation = "EditSite"
	UserOpGetSite                               UserOperation = "GetSite"
	UserOpAddAdmin                              UserOperation = "AddAdmin"
	UserOpGetUnreadRegistrationApplicationCount UserOperation = "GetUnreadRegistrationApplicationCount"
	UserOpListRegistrationApplications          UserOperation = "ListRegistrationApplications"
	UserOpApproveRegistrationApplication        UserOperation = "ApproveRegistrationApplication"
	UserOpBanPerson                             UserOperation = "BanPerson"
	UserOpGetBannedPersons                      UserOperation = "GetBannedPersons"
	UserOpSearch                                UserOperation = "Search"
	UserOpResolveObject                         UserOperation = "ResolveObject"
	UserOpMarkAllAsRead                         UserOperation = "MarkAllAsRead"
	UserOpSaveUserSettings                      UserOperation = "SaveUserSettings"
	UserOpTransferCommunity                     UserOperation = "TransferCommunity"
	UserOpLeaveAdmin                            UserOperation = "LeaveAdmin"
	UserOpDeleteAccount                         UserOperation = "DeleteAccount"
	UserOpPasswordReset                         UserOperation = "PasswordReset"
	UserOpPasswordChange                        UserOperation = "PasswordChange"
	UserOpCreatePrivateMessage                  UserOperation = "CreatePrivateMessage"
	UserOpEditPrivateMessage                    UserOperation = "EditPrivateMessage"
	UserOpDeletePrivateMessage                  UserOperation = "DeletePrivateMessage"
	UserOpMarkPrivateMessageAsRead              UserOperation = "MarkPrivateMessageAsRead"
	UserOpCreatePrivateMessageReport            UserOperation = "CreatePrivateMessageReport"
	UserOpResolvePrivateMessageReport           UserOperation = "ResolvePrivateMessageReport"
	UserOpListPrivateMessageReports             UserOperation = "ListPrivateMessageReports"
	UserOpGetPrivateMessages                    UserOperation = "GetPrivateMessages"
	UserOpUserJoin                              UserOperation = "UserJoin"
	UserOpGetComments                           UserOperation = "GetComments"
	UserOpPostJoin                              UserOperation = "PostJoin"
	UserOpCommunityJoin                         UserOperation = "CommunityJoin"
	UserOpChangePassword                        UserOperation = "ChangePassword"
	UserOpGetSiteMetadata                       UserOperation = "GetSiteMetadata"
	UserOpBlockCommunity                        UserOperation = "BlockCommunity"
	UserOpBlockPerson                           UserOperation = "BlockPerson"
	UserOpPurgePerson                           UserOperation = "PurgePerson"
	UserOpPurgeCommunity                        UserOperation = "PurgeCommunity"
	UserOpPurgePost                             UserOperation = "PurgePost"
	UserOpPurgeComment                          UserOperation = "PurgeComment"
	UserOpCreateCommentReport                   UserOperation = "CreateCommentReport"
	UserOpResolveCommentReport                  UserOperation = "ResolveCommentReport"
	UserOpListCommentReports                    UserOperation = "ListCommentReports"
	UserOpCreatePostReport                      UserOperation = "CreatePostReport"
	UserOpResolvePostReport                     UserOperation = "ResolvePostReport"
	UserOpListPostReports                       UserOperation = "ListPostReports"
	UserOpGetReportCount                        UserOperation = "GetReportCount"
	UserOpGetUnreadCount                        UserOperation = "GetUnreadCount"
	UserOpVerifyEmail                           UserOperation = "VerifyEmail"
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
