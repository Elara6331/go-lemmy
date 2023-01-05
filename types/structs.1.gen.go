package types

type CommentReportView struct {
	CommentReport              CommentReport              `json:"comment_report,omitempty" url:"comment_report,omitempty"`
	Comment                    Comment                    `json:"comment,omitempty" url:"comment,omitempty"`
	Post                       Post                       `json:"post,omitempty" url:"post,omitempty"`
	Community                  CommunitySafe              `json:"community,omitempty" url:"community,omitempty"`
	Creator                    PersonSafe                 `json:"creator,omitempty" url:"creator,omitempty"`
	CommentCreator             PersonSafeAlias1           `json:"comment_creator,omitempty" url:"comment_creator,omitempty"`
	Counts                     CommentAggregates          `json:"counts,omitempty" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool                       `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	MyVote                     Optional[int16]            `json:"my_vote,omitempty" url:"my_vote,omitempty"`
	Resolver                   Optional[PersonSafeAlias2] `json:"resolver,omitempty" url:"resolver,omitempty"`
}
type CommentView struct {
	Comment                    Comment                    `json:"comment,omitempty" url:"comment,omitempty"`
	Creator                    PersonSafe                 `json:"creator,omitempty" url:"creator,omitempty"`
	Recipient                  Optional[PersonSafeAlias1] `json:"recipient,omitempty" url:"recipient,omitempty"`
	Post                       Post                       `json:"post,omitempty" url:"post,omitempty"`
	Community                  CommunitySafe              `json:"community,omitempty" url:"community,omitempty"`
	Counts                     CommentAggregates          `json:"counts,omitempty" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool                       `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	Subscribed                 bool                       `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Saved                      bool                       `json:"saved,omitempty" url:"saved,omitempty"`
	CreatorBlocked             bool                       `json:"creator_blocked,omitempty" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int16]            `json:"my_vote,omitempty" url:"my_vote,omitempty"`
}
type LocalUserView struct {
	LocalUser LocalUser        `json:"local_user,omitempty" url:"local_user,omitempty"`
	Person    Person           `json:"person,omitempty" url:"person,omitempty"`
	Counts    PersonAggregates `json:"counts,omitempty" url:"counts,omitempty"`
}
type LocalUserSettingsView struct {
	LocalUser LocalUserSettings `json:"local_user,omitempty" url:"local_user,omitempty"`
	Person    PersonSafe        `json:"person,omitempty" url:"person,omitempty"`
	Counts    PersonAggregates  `json:"counts,omitempty" url:"counts,omitempty"`
}
type PostReportView struct {
	PostReport                 PostReport                 `json:"post_report,omitempty" url:"post_report,omitempty"`
	Post                       Post                       `json:"post,omitempty" url:"post,omitempty"`
	Community                  CommunitySafe              `json:"community,omitempty" url:"community,omitempty"`
	Creator                    PersonSafe                 `json:"creator,omitempty" url:"creator,omitempty"`
	PostCreator                PersonSafeAlias1           `json:"post_creator,omitempty" url:"post_creator,omitempty"`
	CreatorBannedFromCommunity bool                       `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	MyVote                     Optional[int16]            `json:"my_vote,omitempty" url:"my_vote,omitempty"`
	Counts                     PostAggregates             `json:"counts,omitempty" url:"counts,omitempty"`
	Resolver                   Optional[PersonSafeAlias2] `json:"resolver,omitempty" url:"resolver,omitempty"`
}
type PostView struct {
	Post                       Post            `json:"post,omitempty" url:"post,omitempty"`
	Creator                    PersonSafe      `json:"creator,omitempty" url:"creator,omitempty"`
	Community                  CommunitySafe   `json:"community,omitempty" url:"community,omitempty"`
	CreatorBannedFromCommunity bool            `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	Counts                     PostAggregates  `json:"counts,omitempty" url:"counts,omitempty"`
	Subscribed                 bool            `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Saved                      bool            `json:"saved,omitempty" url:"saved,omitempty"`
	Read                       bool            `json:"read,omitempty" url:"read,omitempty"`
	CreatorBlocked             bool            `json:"creator_blocked,omitempty" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int16] `json:"my_vote,omitempty" url:"my_vote,omitempty"`
}
type PrivateMessageView struct {
	PrivateMessage PrivateMessage   `json:"private_message,omitempty" url:"private_message,omitempty"`
	Creator        PersonSafe       `json:"creator,omitempty" url:"creator,omitempty"`
	Recipient      PersonSafeAlias1 `json:"recipient,omitempty" url:"recipient,omitempty"`
}
type RegistrationApplicationView struct {
	RegistrationApplication RegistrationApplication    `json:"registration_application,omitempty" url:"registration_application,omitempty"`
	CreatorLocalUser        LocalUserSettings          `json:"creator_local_user,omitempty" url:"creator_local_user,omitempty"`
	Creator                 PersonSafe                 `json:"creator,omitempty" url:"creator,omitempty"`
	Admin                   Optional[PersonSafeAlias1] `json:"admin,omitempty" url:"admin,omitempty"`
}
type SiteView struct {
	Site   Site           `json:"site,omitempty" url:"site,omitempty"`
	Counts SiteAggregates `json:"counts,omitempty" url:"counts,omitempty"`
}
