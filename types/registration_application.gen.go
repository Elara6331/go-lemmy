package types

type RegistrationApplication struct {
	ID          int32            `json:"id" url:"id,omitempty"`
	LocalUserID int              `json:"local_user_id" url:"local_user_id,omitempty"`
	Answer      string           `json:"answer" url:"answer,omitempty"`
	AdminID     Optional[int]    `json:"admin_id" url:"admin_id,omitempty"`
	DenyReason  Optional[string] `json:"deny_reason" url:"deny_reason,omitempty"`
	Published   LemmyTime        `json:"published" url:"published,omitempty"`
}
type RegistrationApplicationForm struct {
	LocalUserID Optional[int]              `json:"local_user_id" url:"local_user_id,omitempty"`
	Answer      Optional[string]           `json:"answer" url:"answer,omitempty"`
	AdminID     Optional[int]              `json:"admin_id" url:"admin_id,omitempty"`
	DenyReason  Optional[Optional[string]] `json:"deny_reason" url:"deny_reason,omitempty"`
}
