package types

type Secret struct {
	ID        int32  `json:"id,omitempty" url:"id,omitempty"`
	JwtSecret string `json:"jwt_secret,omitempty" url:"jwt_secret,omitempty"`
}
