package types

type UserJoin struct {
	Auth string `json:"auth"`
}

type UserJoinResponse struct {
	Joined bool `json:"joined"`
	LemmyResponse
}

type CommunityJoin struct {
	CommunityID int `json:"community_id"`
}

type CommunityJoinResponse struct {
	Joined bool `json:"joined"`
	LemmyResponse
}

type ModJoin struct {
	CommunityID int `json:"community_id"`
}

type ModJoinResponse struct {
	Joined bool `json:"joined"`
	LemmyResponse
}

type PostJoin struct {
	PostID int `json:"post_id"`
}

type PostJoinResponse struct {
	Joined bool `json:"joined"`
	LemmyResponse
}
