//  Source: lemmy/crates/db_schema/src/source/actor_language.rs
// Code generated by go.arsenm.dev/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type LocalUserLanguage struct {
	ID          int `json:"id" url:"id,omitempty"`
	LocalUserID int `json:"local_user_id" url:"local_user_id,omitempty"`
	LanguageID  int `json:"language_id" url:"language_id,omitempty"`
}
type LocalUserLanguageForm struct {
	LocalUserID int `json:"local_user_id" url:"local_user_id,omitempty"`
	LanguageID  int `json:"language_id" url:"language_id,omitempty"`
}
type CommunityLanguage struct {
	ID          int `json:"id" url:"id,omitempty"`
	CommunityID int `json:"community_id" url:"community_id,omitempty"`
	LanguageID  int `json:"language_id" url:"language_id,omitempty"`
}
type CommunityLanguageForm struct {
	CommunityID int `json:"community_id" url:"community_id,omitempty"`
	LanguageID  int `json:"language_id" url:"language_id,omitempty"`
}
type SiteLanguage struct {
	ID         int `json:"id" url:"id,omitempty"`
	SiteID     int `json:"site_id" url:"site_id,omitempty"`
	LanguageID int `json:"language_id" url:"language_id,omitempty"`
}
type SiteLanguageForm struct {
	SiteID     int `json:"site_id" url:"site_id,omitempty"`
	LanguageID int `json:"language_id" url:"language_id,omitempty"`
}