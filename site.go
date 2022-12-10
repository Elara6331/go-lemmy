package lemmy

import (
	"context"
	"net/http"

	"go.arsenm.dev/go-lemmy/types"
)

func (c *Client) CreateSite(ctx context.Context, d types.CreateSite) (*types.SiteResponse, error) {
	ar := &types.SiteResponse{}
	res, err := c.req(ctx, http.MethodPost, "/site", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) EditSite(ctx context.Context, d types.EditSite) (*types.SiteResponse, error) {
	ar := &types.SiteResponse{}
	res, err := c.req(ctx, http.MethodPut, "/site", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) Site(ctx context.Context, d types.GetSite) (*types.GetSiteResponse, error) {
	ar := &types.GetSiteResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/site", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
