package types

import "time"

type RegistrationApplication struct {
	ID          int32            `json:"id,omitempty" url:"id,omitempty"`
	LocalUserID int              `json:"local_user_id,omitempty" url:"local_user_id,omitempty"`
	Answer      string           `json:"answer,omitempty" url:"answer,omitempty"`
	AdminID     Optional[int]    `json:"admin_id,omitempty" url:"admin_id,omitempty"`
	DenyReason  Optional[string] `json:"deny_reason,omitempty" url:"deny_reason,omitempty"`
	Published   time.Time        `json:"published,omitempty" url:"published,omitempty"`
}
type RegistrationApplicationForm struct {
	LocalUserID Optional[int]              `json:"local_user_id,omitempty" url:"local_user_id,omitempty"`
	Answer      Optional[string]           `json:"answer,omitempty" url:"answer,omitempty"`
	AdminID     Optional[int]              `json:"admin_id,omitempty" url:"admin_id,omitempty"`
	DenyReason  Optional[Optional[string]] `json:"deny_reason,omitempty" url:"deny_reason,omitempty"`
}
