package lemmy

import (
	"context"
	"net/http"

	"go.arsenm.dev/go-lemmy/types"
)

func (c *Client) AddAdmin(ctx context.Context, d types.AddAdmin) (*types.AddAdminResponse, error) {
	ar := &types.AddAdminResponse{}
	res, err := c.req(ctx, http.MethodPost, "/admin/add", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) ApproveRegistrationApplication(ctx context.Context, d types.ApproveRegistrationApplication) (*types.RegistrationApplicationResponse, error) {
	ar := &types.RegistrationApplicationResponse{}
	res, err := c.req(ctx, http.MethodPut, "/admin/registration_application/approve", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) UnreadRegistrationApplicationCount(ctx context.Context, d types.GetUnreadRegistrationApplicationCount) (*types.GetUnreadRegistrationApplicationCountResponse, error) {
	ar := &types.GetUnreadRegistrationApplicationCountResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/admin/registration_application/count", nil, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) RegistrationApplications(ctx context.Context, d types.ListRegistrationApplications) (*types.ListRegistrationApplicationsResponse, error) {
	ar := &types.ListRegistrationApplicationsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/admin/registration_application/list", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PurgeComment(ctx context.Context, d types.PurgeComment) (*types.PurgeItemResponse, error) {
	ar := &types.PurgeItemResponse{}
	res, err := c.req(ctx, http.MethodPost, "/admin/purge/comment", nil, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PurgeCommunity(ctx context.Context, d types.PurgeCommunity) (*types.PurgeItemResponse, error) {
	ar := &types.PurgeItemResponse{}
	res, err := c.req(ctx, http.MethodPost, "/admin/purge/community", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PurgePerson(ctx context.Context, d types.PurgePerson) (*types.PurgeItemResponse, error) {
	ar := &types.PurgeItemResponse{}
	res, err := c.req(ctx, http.MethodPost, "/admin/purge/person", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PurgePost(ctx context.Context, d types.PurgePost) (*types.PurgeItemResponse, error) {
	ar := &types.PurgeItemResponse{}
	res, err := c.req(ctx, http.MethodPost, "/admin/purge/post", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
