package types

import "time"

type PostReport struct {
	ID               int              `json:"id,omitempty" url:"id,omitempty"`
	CreatorID        int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	PostID           int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	OriginalPostName string           `json:"original_post_name,omitempty" url:"original_post_name,omitempty"`
	OriginalPostURL  Optional[string] `json:"original_post_url,omitempty" url:"original_post_url,omitempty"`
	OriginalPostBody Optional[string] `json:"original_post_body,omitempty" url:"original_post_body,omitempty"`
	Reason           string           `json:"reason,omitempty" url:"reason,omitempty"`
	Resolved         bool             `json:"resolved,omitempty" url:"resolved,omitempty"`
	ResolverID       Optional[int]    `json:"resolver_id,omitempty" url:"resolver_id,omitempty"`
	Published        time.Time        `json:"published,omitempty" url:"published,omitempty"`
	Updated          time.Time        `json:"updated,omitempty" url:"updated,omitempty"`
}
type PostReportForm struct {
	CreatorID        int              `json:"creator_id,omitempty" url:"creator_id,omitempty"`
	PostID           int              `json:"post_id,omitempty" url:"post_id,omitempty"`
	OriginalPostName string           `json:"original_post_name,omitempty" url:"original_post_name,omitempty"`
	OriginalPostURL  Optional[string] `json:"original_post_url,omitempty" url:"original_post_url,omitempty"`
	OriginalPostBody Optional[string] `json:"original_post_body,omitempty" url:"original_post_body,omitempty"`
	Reason           string           `json:"reason,omitempty" url:"reason,omitempty"`
}
