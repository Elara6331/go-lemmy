package types

import "time"

type Site struct {
	ID                         int32            `json:"id,omitempty" url:"id,omitempty"`
	Name                       string           `json:"name,omitempty" url:"name,omitempty"`
	Sidebar                    Optional[string] `json:"sidebar,omitempty" url:"sidebar,omitempty"`
	Published                  time.Time        `json:"published,omitempty" url:"published,omitempty"`
	Updated                    time.Time        `json:"updated,omitempty" url:"updated,omitempty"`
	EnableDownvotes            bool             `json:"enable_downvotes,omitempty" url:"enable_downvotes,omitempty"`
	OpenRegistration           bool             `json:"open_registration,omitempty" url:"open_registration,omitempty"`
	EnableNSFW                 bool             `json:"enable_nsfw,omitempty" url:"enable_nsfw,omitempty"`
	Icon                       Optional[string] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                     Optional[string] `json:"banner,omitempty" url:"banner,omitempty"`
	Description                Optional[string] `json:"description,omitempty" url:"description,omitempty"`
	CommunityCreationAdminOnly bool             `json:"community_creation_admin_only,omitempty" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   bool             `json:"require_email_verification,omitempty" url:"require_email_verification,omitempty"`
	RequireApplication         bool             `json:"require_application,omitempty" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string] `json:"application_question,omitempty" url:"application_question,omitempty"`
	PrivateInstance            bool             `json:"private_instance,omitempty" url:"private_instance,omitempty"`
	ActorID                    string           `json:"actor_id,omitempty" url:"actor_id,omitempty"`
	LastRefreshedAt            time.Time        `json:"last_refreshed_at,omitempty" url:"last_refreshed_at,omitempty"`
	InboxURL                   string           `json:"inbox_url,omitempty" url:"inbox_url,omitempty"`
	PrivateKey                 Optional[string] `json:"private_key,omitempty" url:"private_key,omitempty"`
	PublicKey                  string           `json:"public_key,omitempty" url:"public_key,omitempty"`
	DefaultTheme               string           `json:"default_theme,omitempty" url:"default_theme,omitempty"`
	DefaultPostListingType     string           `json:"default_post_listing_type,omitempty" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string] `json:"legal_information,omitempty" url:"legal_information,omitempty"`
}
type SiteForm struct {
	Name                       string                     `json:"name,omitempty" url:"name,omitempty"`
	Sidebar                    Optional[Optional[string]] `json:"sidebar,omitempty" url:"sidebar,omitempty"`
	Updated                    time.Time                  `json:"updated,omitempty" url:"updated,omitempty"`
	EnableDownvotes            Optional[bool]             `json:"enable_downvotes,omitempty" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]             `json:"open_registration,omitempty" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]             `json:"enable_nsfw,omitempty" url:"enable_nsfw,omitempty"`
	Icon                       Optional[Optional[string]] `json:"icon,omitempty" url:"icon,omitempty"`
	Banner                     Optional[Optional[string]] `json:"banner,omitempty" url:"banner,omitempty"`
	Description                Optional[Optional[string]] `json:"description,omitempty" url:"description,omitempty"`
	CommunityCreationAdminOnly Optional[bool]             `json:"community_creation_admin_only,omitempty" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]             `json:"require_email_verification,omitempty" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]             `json:"require_application,omitempty" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[Optional[string]] `json:"application_question,omitempty" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]             `json:"private_instance,omitempty" url:"private_instance,omitempty"`
	ActorID                    Optional[string]           `json:"actor_id,omitempty" url:"actor_id,omitempty"`
	LastRefreshedAt            time.Time                  `json:"last_refreshed_at,omitempty" url:"last_refreshed_at,omitempty"`
	InboxURL                   Optional[string]           `json:"inbox_url,omitempty" url:"inbox_url,omitempty"`
	PrivateKey                 Optional[Optional[string]] `json:"private_key,omitempty" url:"private_key,omitempty"`
	PublicKey                  Optional[string]           `json:"public_key,omitempty" url:"public_key,omitempty"`
	DefaultTheme               Optional[string]           `json:"default_theme,omitempty" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string]           `json:"default_post_listing_type,omitempty" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string]           `json:"legal_information,omitempty" url:"legal_information,omitempty"`
}
