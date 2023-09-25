package types

// EmptyResponse is a response without any fields.
// It embeds LemmyResponse to capture any errors.
type EmptyResponse struct {
	LemmyResponse
}

// EmptyData is a request without any fields. It contains
// an Auth field so that the auth token is sent to the server.
type EmptyData struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}

// LemmyResponse is embedded in all response structs
// to capture any errors sent by the Lemmy server.
type LemmyResponse struct {
	Error Optional[string] `json:"error" url:"error,omitempty"`
}
