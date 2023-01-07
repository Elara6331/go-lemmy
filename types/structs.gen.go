// Code generated by go.arsenm.dev/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type CommunityBlockView struct {
	Person    PersonSafe    `json:"person" url:"person,omitempty"`
	Community CommunitySafe `json:"community" url:"community,omitempty"`
}
type CommunityFollowerView struct {
	Community CommunitySafe `json:"community" url:"community,omitempty"`
	Follower  PersonSafe    `json:"follower" url:"follower,omitempty"`
}
type CommunityModeratorView struct {
	Community CommunitySafe `json:"community" url:"community,omitempty"`
	Moderator PersonSafe    `json:"moderator" url:"moderator,omitempty"`
}
type CommunityPersonBanView struct {
	Community CommunitySafe `json:"community" url:"community,omitempty"`
	Person    PersonSafe    `json:"person" url:"person,omitempty"`
}
type CommunityView struct {
	Community  CommunitySafe       `json:"community" url:"community,omitempty"`
	Subscribed bool                `json:"subscribed" url:"subscribed,omitempty"`
	Blocked    bool                `json:"blocked" url:"blocked,omitempty"`
	Counts     CommunityAggregates `json:"counts" url:"counts,omitempty"`
}
type PersonBlockView struct {
	Person PersonSafe       `json:"person" url:"person,omitempty"`
	Target PersonSafeAlias1 `json:"target" url:"target,omitempty"`
}
type PersonMentionView struct {
	PersonMention              PersonMention     `json:"person_mention" url:"person_mention,omitempty"`
	Comment                    Comment           `json:"comment" url:"comment,omitempty"`
	Creator                    PersonSafe        `json:"creator" url:"creator,omitempty"`
	Post                       Post              `json:"post" url:"post,omitempty"`
	Community                  CommunitySafe     `json:"community" url:"community,omitempty"`
	Recipient                  PersonSafeAlias1  `json:"recipient" url:"recipient,omitempty"`
	Counts                     CommentAggregates `json:"counts" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community" url:"creator_banned_from_community,omitempty"`
	Subscribed                 bool              `json:"subscribed" url:"subscribed,omitempty"`
	Saved                      bool              `json:"saved" url:"saved,omitempty"`
	CreatorBlocked             bool              `json:"creator_blocked" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int16]   `json:"my_vote" url:"my_vote,omitempty"`
}
type PersonViewSafe struct {
	Person PersonSafe       `json:"person" url:"person,omitempty"`
	Counts PersonAggregates `json:"counts" url:"counts,omitempty"`
}
