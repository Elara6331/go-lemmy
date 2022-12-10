package lemmy

import (
	"context"
	"net/http"

	"go.arsenm.dev/go-lemmy/types"
)

func (c *Client) AddModToCommunity(ctx context.Context, d types.AddModToCommunity) (*types.AddModToCommunityResponse, error) {
	ar := &types.AddModToCommunityResponse{}
	res, err := c.req(ctx, http.MethodPost, "/community/mod", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) BanFromCommunity(ctx context.Context, d types.BanFromCommunity) (*types.BanFromCommunityResponse, error) {
	ar := &types.BanFromCommunityResponse{}
	res, err := c.req(ctx, http.MethodPost, "/community/ban_user", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) BlockCommunity(ctx context.Context, d types.BlockCommunity) (*types.BlockCommunityResponse, error) {
	ar := &types.BlockCommunityResponse{}
	res, err := c.req(ctx, http.MethodPost, "/community/block", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) CreateCommunity(ctx context.Context, d types.CreateCommunity) (*types.CommunityResponse, error) {
	ar := &types.CommunityResponse{}
	res, err := c.req(ctx, http.MethodPost, "/community", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) EditCommunity(ctx context.Context, d types.EditCommunity) (*types.CommunityResponse, error) {
	ar := &types.CommunityResponse{}
	res, err := c.req(ctx, http.MethodPut, "/community", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) FollowCommunity(ctx context.Context, d types.FollowCommunity) (*types.CommunityResponse, error) {
	ar := &types.CommunityResponse{}
	res, err := c.req(ctx, http.MethodPost, "/community/follow", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) DeleteCommunity(ctx context.Context, d types.DeleteCommunity) (*types.CommunityResponse, error) {
	ar := &types.CommunityResponse{}
	res, err := c.req(ctx, http.MethodPost, "/community/delete", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) RemoveCommunity(ctx context.Context, d types.RemoveCommunity) (*types.CommunityResponse, error) {
	ar := &types.CommunityResponse{}
	res, err := c.req(ctx, http.MethodPost, "/community/remove", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) Community(ctx context.Context, d types.GetCommunity) (*types.GetCommunityResponse, error) {
	ar := &types.GetCommunityResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/community", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) Communities(ctx context.Context, d types.ListCommunities) (*types.ListCommunitiesResponse, error) {
	ar := &types.ListCommunitiesResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/community/list", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) TransferCommunity(ctx context.Context, d types.TransferCommunity) (*types.GetCommunityResponse, error) {
	ar := &types.GetCommunityResponse{}
	res, err := c.req(ctx, http.MethodPost, "/community/transfer", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
