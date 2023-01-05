package types

type Secret struct {
	ID        int32  `json:"id" url:"id,omitempty"`
	JwtSecret string `json:"jwt_secret" url:"jwt_secret,omitempty"`
}
