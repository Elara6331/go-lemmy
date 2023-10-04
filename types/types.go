package types

// EmptyResponse is a response without any fields.
// It embeds LemmyResponse to capture any errors.
type EmptyResponse struct {
	LemmyResponse
}

// LemmyResponse is embedded in all response structs
// to capture any errors sent by the Lemmy server.
type LemmyResponse struct {
	Error Optional[string] `json:"error" url:"error,omitempty"`
}
