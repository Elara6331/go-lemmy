package types

import "time"

type PasswordResetRequest struct {
	ID             int32     `json:"id,omitempty" url:"id,omitempty"`
	TokenEncrypted string    `json:"token_encrypted,omitempty" url:"token_encrypted,omitempty"`
	Published      time.Time `json:"published,omitempty" url:"published,omitempty"`
	LocalUserID    int       `json:"local_user_id,omitempty" url:"local_user_id,omitempty"`
}
type PasswordResetRequestForm struct {
	LocalUserID    int    `json:"local_user_id,omitempty" url:"local_user_id,omitempty"`
	TokenEncrypted string `json:"token_encrypted,omitempty" url:"token_encrypted,omitempty"`
}
