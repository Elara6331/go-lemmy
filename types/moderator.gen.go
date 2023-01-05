package types

type ModRemovePost struct {
	ID          int32            `json:"id" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int              `json:"post_id" url:"post_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
	When        LemmyTime        `json:"when_" url:"when_,omitempty"`
}
type ModRemovePostForm struct {
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int              `json:"post_id" url:"post_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
}
type ModLockPost struct {
	ID          int32          `json:"id" url:"id,omitempty"`
	ModPersonID int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id" url:"post_id,omitempty"`
	Locked      Optional[bool] `json:"locked" url:"locked,omitempty"`
	When        LemmyTime      `json:"when_" url:"when_,omitempty"`
}
type ModLockPostForm struct {
	ModPersonID int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id" url:"post_id,omitempty"`
	Locked      Optional[bool] `json:"locked" url:"locked,omitempty"`
}
type ModStickyPost struct {
	ID          int32          `json:"id" url:"id,omitempty"`
	ModPersonID int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id" url:"post_id,omitempty"`
	Stickied    Optional[bool] `json:"stickied" url:"stickied,omitempty"`
	When        LemmyTime      `json:"when_" url:"when_,omitempty"`
}
type ModStickyPostForm struct {
	ModPersonID int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	PostID      int            `json:"post_id" url:"post_id,omitempty"`
	Stickied    Optional[bool] `json:"stickied" url:"stickied,omitempty"`
}
type ModRemoveComment struct {
	ID          int32            `json:"id" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	CommentID   int              `json:"comment_id" url:"comment_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
	When        LemmyTime        `json:"when_" url:"when_,omitempty"`
}
type ModRemoveCommentForm struct {
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	CommentID   int              `json:"comment_id" url:"comment_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
}
type ModRemoveCommunity struct {
	ID          int32            `json:"id" url:"id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
	Expires     LemmyTime        `json:"expires" url:"expires,omitempty"`
	When        LemmyTime        `json:"when_" url:"when_,omitempty"`
}
type ModRemoveCommunityForm struct {
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Removed     Optional[bool]   `json:"removed" url:"removed,omitempty"`
	Expires     LemmyTime        `json:"expires" url:"expires,omitempty"`
}
type ModBanFromCommunity struct {
	ID            int32            `json:"id" url:"id,omitempty"`
	ModPersonID   int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int              `json:"community_id" url:"community_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned" url:"banned,omitempty"`
	Expires       LemmyTime        `json:"expires" url:"expires,omitempty"`
	When          LemmyTime        `json:"when_" url:"when_,omitempty"`
}
type ModBanFromCommunityForm struct {
	ModPersonID   int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int              `json:"community_id" url:"community_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned" url:"banned,omitempty"`
	Expires       LemmyTime        `json:"expires" url:"expires,omitempty"`
}
type ModBan struct {
	ID            int32            `json:"id" url:"id,omitempty"`
	ModPersonID   int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id" url:"other_person_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned" url:"banned,omitempty"`
	Expires       LemmyTime        `json:"expires" url:"expires,omitempty"`
	When          LemmyTime        `json:"when_" url:"when_,omitempty"`
}
type ModHideCommunityForm struct {
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	Hidden      Optional[bool]   `json:"hidden" url:"hidden,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
}
type ModHideCommunity struct {
	ID          int32            `json:"id" url:"id,omitempty"`
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	ModPersonID int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Hidden      Optional[bool]   `json:"hidden" url:"hidden,omitempty"`
	When        LemmyTime        `json:"when_" url:"when_,omitempty"`
}
type ModBanForm struct {
	ModPersonID   int              `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int              `json:"other_person_id" url:"other_person_id,omitempty"`
	Reason        Optional[string] `json:"reason" url:"reason,omitempty"`
	Banned        Optional[bool]   `json:"banned" url:"banned,omitempty"`
	Expires       LemmyTime        `json:"expires" url:"expires,omitempty"`
}
type ModAddCommunity struct {
	ID            int32          `json:"id" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
	When          LemmyTime      `json:"when_" url:"when_,omitempty"`
}
type ModAddCommunityForm struct {
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
}
type ModTransferCommunity struct {
	ID            int32          `json:"id" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
	When          LemmyTime      `json:"when_" url:"when_,omitempty"`
}
type ModTransferCommunityForm struct {
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	CommunityID   int            `json:"community_id" url:"community_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
}
type ModAdd struct {
	ID            int32          `json:"id" url:"id,omitempty"`
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
	When          LemmyTime      `json:"when_" url:"when_,omitempty"`
}
type ModAddForm struct {
	ModPersonID   int            `json:"mod_person_id" url:"mod_person_id,omitempty"`
	OtherPersonID int            `json:"other_person_id" url:"other_person_id,omitempty"`
	Removed       Optional[bool] `json:"removed" url:"removed,omitempty"`
}
