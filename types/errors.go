package types

import (
	"fmt"
	"net/http"
)

// HTTPError represents an error caused by a non-200 HTTP status code
type HTTPError struct {
	Code int
}

func (he HTTPError) Error() string {
	return fmt.Sprintf("%d %s", he.Code, http.StatusText(he.Code))
}

// LemmyError represents an error returned by the Lemmy API
type LemmyError struct {
	ErrStr string
	Code   int
}

func (le LemmyError) Error() string {
	return fmt.Sprintf("%d %s: %s", le.Code, http.StatusText(le.Code), le.ErrStr)
}
