package types

import "time"

type CommunityBlock struct {
	ID          int       `json:"id,omitempty" url:"id,omitempty"`
	PersonID    int       `json:"person_id,omitempty" url:"person_id,omitempty"`
	CommunityID int       `json:"community_id,omitempty" url:"community_id,omitempty"`
	Published   time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type CommunityBlockForm struct {
	PersonID    int `json:"person_id,omitempty" url:"person_id,omitempty"`
	CommunityID int `json:"community_id,omitempty" url:"community_id,omitempty"`
}
