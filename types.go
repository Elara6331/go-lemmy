package lemmy

// emptyResponse is a response without any fields.
// It embeds LemmyResponse to capture any errors.
type emptyResponse struct {
	lemmyResponse
}

// lemmyResponse is embedded in all response structs
// to capture any errors sent by the Lemmy server.
type lemmyResponse struct {
	Error Optional[string] `json:"error" url:"error,omitempty"`
}
