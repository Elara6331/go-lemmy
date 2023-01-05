package types

type CommunityBlockView struct {
	Person    PersonSafe    `json:"person,omitempty" url:"person,omitempty"`
	Community CommunitySafe `json:"community,omitempty" url:"community,omitempty"`
}
type CommunityFollowerView struct {
	Community CommunitySafe `json:"community,omitempty" url:"community,omitempty"`
	Follower  PersonSafe    `json:"follower,omitempty" url:"follower,omitempty"`
}
type CommunityModeratorView struct {
	Community CommunitySafe `json:"community,omitempty" url:"community,omitempty"`
	Moderator PersonSafe    `json:"moderator,omitempty" url:"moderator,omitempty"`
}
type CommunityPersonBanView struct {
	Community CommunitySafe `json:"community,omitempty" url:"community,omitempty"`
	Person    PersonSafe    `json:"person,omitempty" url:"person,omitempty"`
}
type CommunityView struct {
	Community  CommunitySafe       `json:"community,omitempty" url:"community,omitempty"`
	Subscribed bool                `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Blocked    bool                `json:"blocked,omitempty" url:"blocked,omitempty"`
	Counts     CommunityAggregates `json:"counts,omitempty" url:"counts,omitempty"`
}
type PersonBlockView struct {
	Person PersonSafe       `json:"person,omitempty" url:"person,omitempty"`
	Target PersonSafeAlias1 `json:"target,omitempty" url:"target,omitempty"`
}
type PersonMentionView struct {
	PersonMention              PersonMention     `json:"person_mention,omitempty" url:"person_mention,omitempty"`
	Comment                    Comment           `json:"comment,omitempty" url:"comment,omitempty"`
	Creator                    PersonSafe        `json:"creator,omitempty" url:"creator,omitempty"`
	Post                       Post              `json:"post,omitempty" url:"post,omitempty"`
	Community                  CommunitySafe     `json:"community,omitempty" url:"community,omitempty"`
	Recipient                  PersonSafeAlias1  `json:"recipient,omitempty" url:"recipient,omitempty"`
	Counts                     CommentAggregates `json:"counts,omitempty" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	Subscribed                 bool              `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Saved                      bool              `json:"saved,omitempty" url:"saved,omitempty"`
	CreatorBlocked             bool              `json:"creator_blocked,omitempty" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int16]   `json:"my_vote,omitempty" url:"my_vote,omitempty"`
}
type PersonViewSafe struct {
	Person PersonSafe       `json:"person,omitempty" url:"person,omitempty"`
	Counts PersonAggregates `json:"counts,omitempty" url:"counts,omitempty"`
}
