package lemmy

import (
	"context"
	"net/http"

	"go.arsenm.dev/go-lemmy/types"
)

func (c *Client) Search(ctx context.Context, d types.Search) (*types.SearchResponse, error) {
	ar := &types.SearchResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/search", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
