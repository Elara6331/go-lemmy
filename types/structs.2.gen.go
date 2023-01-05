package types

type ModAddCommunityView struct {
	ModAddCommunity ModAddCommunity  `json:"mod_add_community,omitempty" url:"mod_add_community,omitempty"`
	Moderator       PersonSafe       `json:"moderator,omitempty" url:"moderator,omitempty"`
	Community       CommunitySafe    `json:"community,omitempty" url:"community,omitempty"`
	ModdedPerson    PersonSafeAlias1 `json:"modded_person,omitempty" url:"modded_person,omitempty"`
}
type ModAddView struct {
	ModAdd       ModAdd           `json:"mod_add,omitempty" url:"mod_add,omitempty"`
	Moderator    PersonSafe       `json:"moderator,omitempty" url:"moderator,omitempty"`
	ModdedPerson PersonSafeAlias1 `json:"modded_person,omitempty" url:"modded_person,omitempty"`
}
type ModBanFromCommunityView struct {
	ModBanFromCommunity ModBanFromCommunity `json:"mod_ban_from_community,omitempty" url:"mod_ban_from_community,omitempty"`
	Moderator           PersonSafe          `json:"moderator,omitempty" url:"moderator,omitempty"`
	Community           CommunitySafe       `json:"community,omitempty" url:"community,omitempty"`
	BannedPerson        PersonSafeAlias1    `json:"banned_person,omitempty" url:"banned_person,omitempty"`
}
type ModBanView struct {
	ModBan       ModBan           `json:"mod_ban,omitempty" url:"mod_ban,omitempty"`
	Moderator    PersonSafe       `json:"moderator,omitempty" url:"moderator,omitempty"`
	BannedPerson PersonSafeAlias1 `json:"banned_person,omitempty" url:"banned_person,omitempty"`
}
type ModHideCommunityView struct {
	ModHideCommunity ModHideCommunity `json:"mod_hide_community,omitempty" url:"mod_hide_community,omitempty"`
	Admin            PersonSafe       `json:"admin,omitempty" url:"admin,omitempty"`
	Community        CommunitySafe    `json:"community,omitempty" url:"community,omitempty"`
}
type ModLockPostView struct {
	ModLockPost ModLockPost   `json:"mod_lock_post,omitempty" url:"mod_lock_post,omitempty"`
	Moderator   PersonSafe    `json:"moderator,omitempty" url:"moderator,omitempty"`
	Post        Post          `json:"post,omitempty" url:"post,omitempty"`
	Community   CommunitySafe `json:"community,omitempty" url:"community,omitempty"`
}
type ModRemoveCommentView struct {
	ModRemoveComment ModRemoveComment `json:"mod_remove_comment,omitempty" url:"mod_remove_comment,omitempty"`
	Moderator        PersonSafe       `json:"moderator,omitempty" url:"moderator,omitempty"`
	Comment          Comment          `json:"comment,omitempty" url:"comment,omitempty"`
	Commenter        PersonSafeAlias1 `json:"commenter,omitempty" url:"commenter,omitempty"`
	Post             Post             `json:"post,omitempty" url:"post,omitempty"`
	Community        CommunitySafe    `json:"community,omitempty" url:"community,omitempty"`
}
type ModRemoveCommunityView struct {
	ModRemoveCommunity ModRemoveCommunity `json:"mod_remove_community,omitempty" url:"mod_remove_community,omitempty"`
	Moderator          PersonSafe         `json:"moderator,omitempty" url:"moderator,omitempty"`
	Community          CommunitySafe      `json:"community,omitempty" url:"community,omitempty"`
}
type ModRemovePostView struct {
	ModRemovePost ModRemovePost `json:"mod_remove_post,omitempty" url:"mod_remove_post,omitempty"`
	Moderator     PersonSafe    `json:"moderator,omitempty" url:"moderator,omitempty"`
	Post          Post          `json:"post,omitempty" url:"post,omitempty"`
	Community     CommunitySafe `json:"community,omitempty" url:"community,omitempty"`
}
type ModStickyPostView struct {
	ModStickyPost ModStickyPost `json:"mod_sticky_post,omitempty" url:"mod_sticky_post,omitempty"`
	Moderator     PersonSafe    `json:"moderator,omitempty" url:"moderator,omitempty"`
	Post          Post          `json:"post,omitempty" url:"post,omitempty"`
	Community     CommunitySafe `json:"community,omitempty" url:"community,omitempty"`
}
type ModTransferCommunityView struct {
	ModTransferCommunity ModTransferCommunity `json:"mod_transfer_community,omitempty" url:"mod_transfer_community,omitempty"`
	Moderator            PersonSafe           `json:"moderator,omitempty" url:"moderator,omitempty"`
	Community            CommunitySafe        `json:"community,omitempty" url:"community,omitempty"`
	ModdedPerson         PersonSafeAlias1     `json:"modded_person,omitempty" url:"modded_person,omitempty"`
}
