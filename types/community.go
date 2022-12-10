package types

type GetCommunity struct {
	ID   Optional[int]    `json:"id,omitempty" url:"id,omitempty"`
	Name Optional[string] `json:"name,omitempty" url:"name,omitempty"`
	Auth Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type GetCommunityResponse struct {
	CommunityView       CommunityView            `json:"community_view,omitempty" url:"community_view,omitempty"`
	Site                Optional[Site]           `json:"site,omitempty" url:"site,omitempty"`
	Moderators          []CommunityModeratorView `json:"moderators,omitempty" url:"moderators,omitempty"`
	Online              uint                     `json:"online,omitempty" url:"online,omitempty"`
	DiscussionLanguages []int                    `json:"discussion_languages,omitempty" url:"discussion_languages,omitempty"`
	DefaultPostLanguage Optional[int]            `json:"default_post_language,omitempty" url:"default_post_language,omitempty"`
	LemmyResponse
}

type CreateCommunity struct {
	Name                    string           `json:"name,omitempty" url:"name,omitempty"`
	Title                   string           `json:"title,omitempty" url:"title,omitempty"`
	Description             Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	Icon                    Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	NSFW                    Optional[bool]   `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	PostingRestrictedToMods Optional[bool]   `json:"posting_restricted_to_mods,omitempty" url:"posting_restricted_to_mods,omitempty"`
	Auth                    string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type CommunityResponse struct {
	CommunityView CommunityView `json:"community_view,omitempty" url:"community_view,omitempty"`
	LemmyResponse
}

type ListCommunities struct {
	Type  Optional[ListingType] `json:"type,omitempty" url:"type,omitempty"`
	Sort  Optional[SortType]    `json:"sort,omitempty" url:"sort,omitempty"`
	Page  Optional[int64]       `json:"page,omitempty" url:"page,omitempty"`
	Limit Optional[int64]       `json:"limit,omitempty" url:"limit,omitempty"`
	Auth  Optional[string]      `json:"auth,omitempty" url:"auth,omitempty"`
}

type ListCommunitiesResponse struct {
	Communities []CommunityView `json:"communities,omitempty" url:"communities,omitempty"`
	LemmyResponse
}

type BanFromCommunity struct {
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int              `json:"person_id,omitempty" url:"person_id,omitempty"`
	Ban         bool             `json:"ban,omitempty" url:"ban,omitempty"`
	RemoveData  Optional[bool]   `json:"remove_data,omitempty" url:"remove_data,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Expires     Optional[int64]  `json:"expires,omitempty" url:"expires,omitempty"`
	Auth        string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type BanFromCommunityResponse struct {
	PersonView PersonViewSafe `json:"person_view,omitempty" url:"person_view,omitempty"`
	Banned     bool           `json:"banned,omitempty" url:"banned,omitempty"`
	LemmyResponse
}

type AddModToCommunity struct {
	CommunityID int    `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int    `json:"person_id,omitempty" url:"person_id,omitempty"`
	Added       bool   `json:"added,omitempty" url:"added,omitempty"`
	Auth        string `json:"auth,omitempty" url:"auth,omitempty"`
}

type AddModToCommunityResponse struct {
	Moderators []CommunityModeratorView `json:"moderators,omitempty" url:"moderators,omitempty"`
	LemmyResponse
}

type EditCommunity struct {
	CommunityID             int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Title                   Optional[string] `json:"title,omitempty" url:"title,omitempty"`
	Description             Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	Icon                    Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	NSFW                    Optional[bool]   `json:"nsfw,omitempty" url:"nsfw,omitempty"`
	PostingRestrictedToMods Optional[bool]   `json:"posting_restricted_to_mods,omitempty" url:"posting_restricted_to_mods,omitempty"`
	DiscussionLanguages     Optional[[]int]  `json:"discussion_languages,omitempty" url:"discussion_languages,omitempty"`
	Auth                    string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type HideCommunity struct {
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Hidden      bool             `json:"hidden,omitempty" url:"hidden,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Auth        string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type DeleteCommunity struct {
	CommunityID int    `json:"community_id,omitempty" url:"community_id,omitempty"`
	Deleted     bool   `json:"deleted,omitempty" url:"deleted,omitempty"`
	Auth        string `json:"auth,omitempty" url:"auth,omitempty"`
}

type RemoveCommunity struct {
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed     bool             `json:"removed,omitempty" url:"removed,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Expires     Optional[int64]  `json:"expires,omitempty" url:"expires,omitempty"`
	Auth        string           `json:"auth,omitempty" url:"auth,omitempty"`
}

type FollowCommunity struct {
	CommunityID int    `json:"community_id,omitempty" url:"community_id,omitempty"`
	Follow      bool   `json:"follow,omitempty" url:"follow,omitempty"`
	Auth        string `json:"auth,omitempty" url:"auth,omitempty"`
}

type BlockCommunity struct {
	CommunityID int    `json:"community_id,omitempty" url:"community_id,omitempty"`
	Block       bool   `json:"block,omitempty" url:"block,omitempty"`
	Auth        string `json:"auth,omitempty" url:"auth,omitempty"`
}

type BlockCommunityResponse struct {
	CommunityView CommunityView `json:"community_view,omitempty" url:"community_view,omitempty"`
	Blocked       bool          `json:"blocked,omitempty" url:"blocked,omitempty"`
	LemmyResponse
}

type TransferCommunity struct {
	CommunityID int    `json:"community_id,omitempty" url:"community_id,omitempty"`
	PersonID    int    `json:"person_id,omitempty" url:"person_id,omitempty"`
	Auth        string `json:"auth,omitempty" url:"auth,omitempty"`
}
