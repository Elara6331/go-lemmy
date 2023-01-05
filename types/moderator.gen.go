package types

import "time"

type ModRemovePost struct {
	ID          int32            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	When        time.Time        `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModRemovePostForm struct {
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
}
type ModLockPost struct {
	ID          int32          `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id,omitempty" url:"post_id,omitempty"`
	Locked      Optional[bool] `json:"locked,omitempty" url:"locked,omitempty"`
	When        time.Time      `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModLockPostForm struct {
	ModPersonID int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id,omitempty" url:"post_id,omitempty"`
	Locked      Optional[bool] `json:"locked,omitempty" url:"locked,omitempty"`
}
type ModStickyPost struct {
	ID          int32          `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id,omitempty" url:"post_id,omitempty"`
	Stickied    Optional[bool] `json:"stickied,omitempty" url:"stickied,omitempty"`
	When        time.Time      `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModStickyPostForm struct {
	ModPersonID int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id,omitempty" url:"post_id,omitempty"`
	Stickied    Optional[bool] `json:"stickied,omitempty" url:"stickied,omitempty"`
}
type ModRemoveComment struct {
	ID          int32            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	CommentID   int              `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	When        time.Time        `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModRemoveCommentForm struct {
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	CommentID   int              `json:"comment_id,omitempty" url:"comment_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
}
type ModRemoveCommunity struct {
	ID          int32            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	Expires     time.Time        `json:"expires,omitempty" url:"expires,omitempty"`
	When        time.Time        `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModRemoveCommunityForm struct {
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed,omitempty" url:"removed,omitempty"`
	Expires     time.Time        `json:"expires,omitempty" url:"expires,omitempty"`
}
type ModBanFromCommunity struct {
	ID            int32            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned,omitempty" url:"banned,omitempty"`
	Expires       time.Time        `json:"expires,omitempty" url:"expires,omitempty"`
	When          time.Time        `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModBanFromCommunityForm struct {
	ModPersonID   int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned,omitempty" url:"banned,omitempty"`
	Expires       time.Time        `json:"expires,omitempty" url:"expires,omitempty"`
}
type ModBan struct {
	ID            int32            `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned,omitempty" url:"banned,omitempty"`
	Expires       time.Time        `json:"expires,omitempty" url:"expires,omitempty"`
	When          time.Time        `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModHideCommunityForm struct {
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	Hidden      Optional[bool]   `json:"hidden,omitempty" url:"hidden,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
}
type ModHideCommunity struct {
	ID          int32            `json:"id,omitempty" url:"id,omitempty"`
	CommunityID int              `json:"community_id,omitempty" url:"community_id,omitempty"`
	ModPersonID int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	Reason      Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Hidden      Optional[bool]   `json:"hidden,omitempty" url:"hidden,omitempty"`
	When        time.Time        `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModBanForm struct {
	ModPersonID   int              `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	Reason        Optional[string] `json:"reason,omitempty" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned,omitempty" url:"banned,omitempty"`
	Expires       time.Time        `json:"expires,omitempty" url:"expires,omitempty"`
}
type ModAddCommunity struct {
	ID            int32          `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
	When          time.Time      `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModAddCommunityForm struct {
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
}
type ModTransferCommunity struct {
	ID            int32          `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
	When          time.Time      `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModTransferCommunityForm struct {
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id,omitempty" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
}
type ModAdd struct {
	ID            int32          `json:"id,omitempty" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
	When          time.Time      `json:"when_,omitempty" url:"when_,omitempty"`
}
type ModAddForm struct {
	ModPersonID   int            `json:"mod_person_id,omitempty" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id,omitempty" url:"other_person_id,omitempty"`
	Removed       Optional[bool] `json:"removed,omitempty" url:"removed,omitempty"`
}
