package types

import "time"

type Community struct {
	ID                      int              `json:"id,omitempty" url:"id,omitempty"`
	Name                    string           `json:"name,omitempty" url:"name,omitempty"`
	Title                   string           `json:"title,omitempty" url:"title,omitempty"`
	Description             Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	Removed                 bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Published               time.Time        `json:"published,omitempty" url:"published,omitempty"`
	Updated                 time.Time        `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted                 bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	NSFW                    bool             `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	ActorID                 string           `json:"actor_id,omitempty" url:"actor_id,omitempty"`
	Local                   bool             `json:"local,omitempty" url:"local,omitempty"`
	PrivateKey              Optional[string] `json:"private_key,omitempty" url:"private_key,omitempty"`
	PublicKey               string           `json:"public_key,omitempty" url:"public_key,omitempty"`
	LastRefreshedAt         time.Time        `json:"last_refreshed_at,omitempty" url:"last_refreshed_at,omitempty"`
	Icon                    Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	FollowersURL            string           `json:"followers_url,omitempty" url:"followers_url,omitempty"`
	InboxURL                string           `json:"inbox_url,omitempty" url:"inbox_url,omitempty"`
	SharedInboxURL          Optional[string] `json:"shared_inbox_url,omitempty" url:"shared_inbox_url,omitempty"`
	Hidden                  bool             `json:"hidden,omitempty" url:"hidden,omitempty"`
	PostingRestrictedToMods bool             `json:"posting_restricted_to_mods,omitempty" url:"posting_restricted_to_mods,omitempty"`
}
type CommunitySafe struct {
	ID                      int              `json:"id,omitempty" url:"id,omitempty"`
	Name                    string           `json:"name,omitempty" url:"name,omitempty"`
	Title                   string           `json:"title,omitempty" url:"title,omitempty"`
	Description             Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	Removed                 bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Published               time.Time        `json:"published,omitempty" url:"published,omitempty"`
	Updated                 time.Time        `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted                 bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	NSFW                    bool             `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	ActorID                 string           `json:"actor_id,omitempty" url:"actor_id,omitempty"`
	Local                   bool             `json:"local,omitempty" url:"local,omitempty"`
	Icon                    Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	Hidden                  bool             `json:"hidden,omitempty" url:"hidden,omitempty"`
	PostingRestrictedToMods bool             `json:"posting_restricted_to_mods,omitempty" url:"posting_restricted_to_mods,omitempty"`
}
type CommunityForm struct {
	Name                    string                     `json:"name,omitempty" url:"name,omitempty"`
	Title                   string                     `json:"title,omitempty" url:"title,omitempty"`
	Description             Optional[string]           `json:"description,omitempty" url:"description,omitempty"`
	Removed                 Optional[bool]             `json:"removed,omitempty" url:"removed,omitempty"`
	Published               time.Time                  `json:"published,omitempty" url:"published,omitempty"`
	Updated                 time.Time                  `json:"updated,omitempty" url:"updated,omitempty"`
	Deleted                 Optional[bool]             `json:"deleted,omitempty" url:"deleted,omitempty"`
	NSFW                    Optional[bool]             `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	ActorID                 Optional[string]           `json:"actor_id,omitempty" url:"actor_id,omitempty"`
	Local                   Optional[bool]             `json:"local,omitempty" url:"local,omitempty"`
	PrivateKey              Optional[Optional[string]] `json:"private_key,omitempty" url:"private_key,omitempty"`
	PublicKey               Optional[string]           `json:"public_key,omitempty" url:"public_key,omitempty"`
	LastRefreshedAt         time.Time                  `json:"last_refreshed_at,omitempty" url:"last_refreshed_at,omitempty"`
	Icon                    Optional[Optional[string]] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                  Optional[Optional[string]] `json:"banner,omitempty" url:"banner,omitempty"`
	FollowersURL            Optional[string]           `json:"followers_url,omitempty" url:"followers_url,omitempty"`
	InboxURL                Optional[string]           `json:"inbox_url,omitempty" url:"inbox_url,omitempty"`
	SharedInboxURL          Optional[Optional[string]] `json:"shared_inbox_url,omitempty" url:"shared_inbox_url,omitempty"`
	Hidden                  Optional[bool]             `json:"hidden,omitempty" url:"hidden,omitempty"`
	PostingRestrictedToMods Optional[bool]             `json:"posting_restricted_to_mods,omitempty" url:"posting_restricted_to_mods,omitempty"`
}
type CommunityModerator struct {
	ID          int32     `json:"id,omitempty" url:"id,omitempty"`
	CommunityID int       `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	Published   time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type CommunityModeratorForm struct {
	CommunityID int `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int `json:"person_id,omitempty" url:"person_id,omitempty"`
}
type CommunityPersonBan struct {
	ID          int32     `json:"id,omitempty" url:"id,omitempty"`
	CommunityID int       `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	Published   time.Time `json:"published,omitempty" url:"published,omitempty"`
	Expires     time.Time `json:"expires,omitempty" url:"expires,omitempty"`
}
type CommunityPersonBanForm struct {
	CommunityID int       `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	Expires     time.Time `json:"expires,omitempty" url:"expires,omitempty"`
}
type CommunityFollower struct {
	ID          int32          `json:"id,omitempty" url:"id,omitempty"`
	CommunityID int            `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int            `json:"person_id,omitempty" url:"person_id,omitempty"`
	Published   time.Time      `json:"published,omitempty" url:"published,omitempty"`
	Pending     Optional[bool] `json:"pending,omitempty" url:"pending,omitempty"`
}
type CommunityFollowerForm struct {
	CommunityID int  `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int  `json:"person_id,omitempty" url:"person_id,omitempty"`
	Pending     bool `json:"pending,omitempty" url:"pending,omitempty"`
}
