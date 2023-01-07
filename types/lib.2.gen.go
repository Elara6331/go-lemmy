// Code generated by go.arsenm.dev/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type UserOperation string

const (
	UserOperationLogin                                 UserOperation = "Login"
	UserOperationGetCaptcha                            UserOperation = "GetCaptcha"
	UserOperationMarkCommentAsRead                     UserOperation = "MarkCommentAsRead"
	UserOperationSaveComment                           UserOperation = "SaveComment"
	UserOperationCreateCommentLike                     UserOperation = "CreateCommentLike"
	UserOperationCreateCommentReport                   UserOperation = "CreateCommentReport"
	UserOperationResolveCommentReport                  UserOperation = "ResolveCommentReport"
	UserOperationListCommentReports                    UserOperation = "ListCommentReports"
	UserOperationCreatePostLike                        UserOperation = "CreatePostLike"
	UserOperationLockPost                              UserOperation = "LockPost"
	UserOperationStickyPost                            UserOperation = "StickyPost"
	UserOperationMarkPostAsRead                        UserOperation = "MarkPostAsRead"
	UserOperationSavePost                              UserOperation = "SavePost"
	UserOperationCreatePostReport                      UserOperation = "CreatePostReport"
	UserOperationResolvePostReport                     UserOperation = "ResolvePostReport"
	UserOperationListPostReports                       UserOperation = "ListPostReports"
	UserOperationGetReportCount                        UserOperation = "GetReportCount"
	UserOperationGetUnreadCount                        UserOperation = "GetUnreadCount"
	UserOperationVerifyEmail                           UserOperation = "VerifyEmail"
	UserOperationFollowCommunity                       UserOperation = "FollowCommunity"
	UserOperationGetReplies                            UserOperation = "GetReplies"
	UserOperationGetPersonMentions                     UserOperation = "GetPersonMentions"
	UserOperationMarkPersonMentionAsRead               UserOperation = "MarkPersonMentionAsRead"
	UserOperationGetModlog                             UserOperation = "GetModlog"
	UserOperationBanFromCommunity                      UserOperation = "BanFromCommunity"
	UserOperationAddModToCommunity                     UserOperation = "AddModToCommunity"
	UserOperationAddAdmin                              UserOperation = "AddAdmin"
	UserOperationGetUnreadRegistrationApplicationCount UserOperation = "GetUnreadRegistrationApplicationCount"
	UserOperationListRegistrationApplications          UserOperation = "ListRegistrationApplications"
	UserOperationApproveRegistrationApplication        UserOperation = "ApproveRegistrationApplication"
	UserOperationBanPerson                             UserOperation = "BanPerson"
	UserOperationGetBannedPersons                      UserOperation = "GetBannedPersons"
	UserOperationSearch                                UserOperation = "Search"
	UserOperationResolveObject                         UserOperation = "ResolveObject"
	UserOperationMarkAllAsRead                         UserOperation = "MarkAllAsRead"
	UserOperationSaveUserSettings                      UserOperation = "SaveUserSettings"
	UserOperationTransferCommunity                     UserOperation = "TransferCommunity"
	UserOperationLeaveAdmin                            UserOperation = "LeaveAdmin"
	UserOperationPasswordReset                         UserOperation = "PasswordReset"
	UserOperationPasswordChange                        UserOperation = "PasswordChange"
	UserOperationMarkPrivateMessageAsRead              UserOperation = "MarkPrivateMessageAsRead"
	UserOperationUserJoin                              UserOperation = "UserJoin"
	UserOperationGetSiteConfig                         UserOperation = "GetSiteConfig"
	UserOperationSaveSiteConfig                        UserOperation = "SaveSiteConfig"
	UserOperationPostJoin                              UserOperation = "PostJoin"
	UserOperationCommunityJoin                         UserOperation = "CommunityJoin"
	UserOperationModJoin                               UserOperation = "ModJoin"
	UserOperationChangePassword                        UserOperation = "ChangePassword"
	UserOperationGetSiteMetadata                       UserOperation = "GetSiteMetadata"
	UserOperationBlockCommunity                        UserOperation = "BlockCommunity"
	UserOperationBlockPerson                           UserOperation = "BlockPerson"
)

type UserOperationCrud string

const (
	UserOperationCrudCreateSite           UserOperationCrud = "CreateSite"
	UserOperationCrudGetSite              UserOperationCrud = "GetSite"
	UserOperationCrudEditSite             UserOperationCrud = "EditSite"
	UserOperationCrudCreateCommunity      UserOperationCrud = "CreateCommunity"
	UserOperationCrudListCommunities      UserOperationCrud = "ListCommunities"
	UserOperationCrudGetCommunity         UserOperationCrud = "GetCommunity"
	UserOperationCrudEditCommunity        UserOperationCrud = "EditCommunity"
	UserOperationCrudDeleteCommunity      UserOperationCrud = "DeleteCommunity"
	UserOperationCrudRemoveCommunity      UserOperationCrud = "RemoveCommunity"
	UserOperationCrudCreatePost           UserOperationCrud = "CreatePost"
	UserOperationCrudGetPost              UserOperationCrud = "GetPost"
	UserOperationCrudGetPosts             UserOperationCrud = "GetPosts"
	UserOperationCrudEditPost             UserOperationCrud = "EditPost"
	UserOperationCrudDeletePost           UserOperationCrud = "DeletePost"
	UserOperationCrudRemovePost           UserOperationCrud = "RemovePost"
	UserOperationCrudCreateComment        UserOperationCrud = "CreateComment"
	UserOperationCrudGetComment           UserOperationCrud = "GetComment"
	UserOperationCrudGetComments          UserOperationCrud = "GetComments"
	UserOperationCrudEditComment          UserOperationCrud = "EditComment"
	UserOperationCrudDeleteComment        UserOperationCrud = "DeleteComment"
	UserOperationCrudRemoveComment        UserOperationCrud = "RemoveComment"
	UserOperationCrudRegister             UserOperationCrud = "Register"
	UserOperationCrudGetPersonDetails     UserOperationCrud = "GetPersonDetails"
	UserOperationCrudDeleteAccount        UserOperationCrud = "DeleteAccount"
	UserOperationCrudCreatePrivateMessage UserOperationCrud = "CreatePrivateMessage"
	UserOperationCrudGetPrivateMessages   UserOperationCrud = "GetPrivateMessages"
	UserOperationCrudEditPrivateMessage   UserOperationCrud = "EditPrivateMessage"
	UserOperationCrudDeletePrivateMessage UserOperationCrud = "DeletePrivateMessage"
)
