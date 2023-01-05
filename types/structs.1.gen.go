package types

type CommentReportView struct {
	CommentReport              CommentReport              `json:"comment_report" url:"comment_report,omitempty"`
	Comment                    Comment                    `json:"comment" url:"comment,omitempty"`
	Post                       Post                       `json:"post" url:"post,omitempty"`
	Community                  CommunitySafe              `json:"community" url:"community,omitempty"`
	Creator                    PersonSafe                 `json:"creator" url:"creator,omitempty"`
	CommentCreator             PersonSafeAlias1           `json:"comment_creator" url:"comment_creator,omitempty"`
	Counts                     CommentAggregates          `json:"counts" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool                       `json:"creator_banned_from_community" url:"creator_banned_from_community,omitempty"`
	MyVote                     Optional[int16]            `json:"my_vote" url:"my_vote,omitempty"`
	Resolver                   Optional[PersonSafeAlias2] `json:"resolver" url:"resolver,omitempty"`
}
type CommentView struct {
	Comment                    Comment                    `json:"comment" url:"comment,omitempty"`
	Creator                    PersonSafe                 `json:"creator" url:"creator,omitempty"`
	Recipient                  Optional[PersonSafeAlias1] `json:"recipient" url:"recipient,omitempty"`
	Post                       Post                       `json:"post" url:"post,omitempty"`
	Community                  CommunitySafe              `json:"community" url:"community,omitempty"`
	Counts                     CommentAggregates          `json:"counts" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool                       `json:"creator_banned_from_community" url:"creator_banned_from_community,omitempty"`
	Subscribed                 bool                       `json:"subscribed" url:"subscribed,omitempty"`
	Saved                      bool                       `json:"saved" url:"saved,omitempty"`
	CreatorBlocked             bool                       `json:"creator_blocked" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int16]            `json:"my_vote" url:"my_vote,omitempty"`
}
type LocalUserView struct {
	LocalUser LocalUser        `json:"local_user" url:"local_user,omitempty"`
	Person    Person           `json:"person" url:"person,omitempty"`
	Counts    PersonAggregates `json:"counts" url:"counts,omitempty"`
}
type LocalUserSettingsView struct {
	LocalUser LocalUserSettings `json:"local_user" url:"local_user,omitempty"`
	Person    PersonSafe        `json:"person" url:"person,omitempty"`
	Counts    PersonAggregates  `json:"counts" url:"counts,omitempty"`
}
type PostReportView struct {
	PostReport                 PostReport                 `json:"post_report" url:"post_report,omitempty"`
	Post                       Post                       `json:"post" url:"post,omitempty"`
	Community                  CommunitySafe              `json:"community" url:"community,omitempty"`
	Creator                    PersonSafe                 `json:"creator" url:"creator,omitempty"`
	PostCreator                PersonSafeAlias1           `json:"post_creator" url:"post_creator,omitempty"`
	CreatorBannedFromCommunity bool                       `json:"creator_banned_from_community" url:"creator_banned_from_community,omitempty"`
	MyVote                     Optional[int16]            `json:"my_vote" url:"my_vote,omitempty"`
	Counts                     PostAggregates             `json:"counts" url:"counts,omitempty"`
	Resolver                   Optional[PersonSafeAlias2] `json:"resolver" url:"resolver,omitempty"`
}
type PostView struct {
	Post                       Post            `json:"post" url:"post,omitempty"`
	Creator                    PersonSafe      `json:"creator" url:"creator,omitempty"`
	Community                  CommunitySafe   `json:"community" url:"community,omitempty"`
	CreatorBannedFromCommunity bool            `json:"creator_banned_from_community" url:"creator_banned_from_community,omitempty"`
	Counts                     PostAggregates  `json:"counts" url:"counts,omitempty"`
	Subscribed                 bool            `json:"subscribed" url:"subscribed,omitempty"`
	Saved                      bool            `json:"saved" url:"saved,omitempty"`
	Read                       bool            `json:"read" url:"read,omitempty"`
	CreatorBlocked             bool            `json:"creator_blocked" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int16] `json:"my_vote" url:"my_vote,omitempty"`
}
type PrivateMessageView struct {
	PrivateMessage PrivateMessage   `json:"private_message" url:"private_message,omitempty"`
	Creator        PersonSafe       `json:"creator" url:"creator,omitempty"`
	Recipient      PersonSafeAlias1 `json:"recipient" url:"recipient,omitempty"`
}
type RegistrationApplicationView struct {
	RegistrationApplication RegistrationApplication    `json:"registration_application" url:"registration_application,omitempty"`
	CreatorLocalUser        LocalUserSettings          `json:"creator_local_user" url:"creator_local_user,omitempty"`
	Creator                 PersonSafe                 `json:"creator" url:"creator,omitempty"`
	Admin                   Optional[PersonSafeAlias1] `json:"admin" url:"admin,omitempty"`
}
type SiteView struct {
	Site   Site           `json:"site" url:"site,omitempty"`
	Counts SiteAggregates `json:"counts" url:"counts,omitempty"`
}
