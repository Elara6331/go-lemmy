package types

type UserJoin struct {
	Auth string `json:"auth,omitempty" url:"auth,omitempty"`
}
type UserJoinResponse struct {
	Joined bool `json:"joined,omitempty" url:"joined,omitempty"`
	LemmyResponse
}
type CommunityJoin struct {
	CommunityID int `json:"community_id,omitempty" url:"community_id,omitempty"`
}
type CommunityJoinResponse struct {
	Joined bool `json:"joined,omitempty" url:"joined,omitempty"`
	LemmyResponse
}
type ModJoin struct {
	CommunityID int `json:"community_id,omitempty" url:"community_id,omitempty"`
}
type ModJoinResponse struct {
	Joined bool `json:"joined,omitempty" url:"joined,omitempty"`
	LemmyResponse
}
type PostJoin struct {
	PostID int `json:"post_id,omitempty" url:"post_id,omitempty"`
}
type PostJoinResponse struct {
	Joined bool `json:"joined,omitempty" url:"joined,omitempty"`
	LemmyResponse
}
