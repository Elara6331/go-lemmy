package types

import "time"

type EmailVerification struct {
	ID               int32     `json:"id,omitempty" url:"id,omitempty"`
	LocalUserID      int       `json:"local_user_id,omitempty" url:"local_user_id,omitempty"`
	Email            string    `json:"email,omitempty" url:"email,omitempty"`
	VerificationCode string    `json:"verification_code,omitempty" url:"verification_code,omitempty"`
	Published        time.Time `json:"published,omitempty" url:"published,omitempty"`
}
type EmailVerificationForm struct {
	LocalUserID       int    `json:"local_user_id,omitempty" url:"local_user_id,omitempty"`
	Email             string `json:"email,omitempty" url:"email,omitempty"`
	VerificationToken string `json:"verification_token,omitempty" url:"verification_token,omitempty"`
}
