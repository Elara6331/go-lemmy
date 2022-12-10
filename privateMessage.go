package lemmy

import (
	"context"
	"net/http"

	"go.arsenm.dev/go-lemmy/types"
)

func (c *Client) CreatePrivateMessage(ctx context.Context, d types.CreatePrivateMessage) (*types.PrivateMessageResponse, error) {
	ar := &types.PrivateMessageResponse{}
	res, err := c.req(ctx, http.MethodPost, "/private_message", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) CreatePrivateMessageReport(ctx context.Context, d types.CreatePrivateMessageReport) (*types.PrivateMessageReportResponse, error) {
	ar := &types.PrivateMessageReportResponse{}
	res, err := c.req(ctx, http.MethodPost, "/private_message/report", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) DeletePrivateMessage(ctx context.Context, d types.DeletePrivateMessage) (*types.PrivateMessageResponse, error) {
	ar := &types.PrivateMessageResponse{}
	res, err := c.req(ctx, http.MethodPost, "/private_message/delete", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) EditPrivateMessage(ctx context.Context, d types.EditPrivateMessage) (*types.PrivateMessageResponse, error) {
	ar := &types.PrivateMessageResponse{}
	res, err := c.req(ctx, http.MethodPut, "/private_message", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PrivateMessages(ctx context.Context, d types.GetPrivateMessages) (*types.PrivateMessagesResponse, error) {
	ar := &types.PrivateMessagesResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/private_message/list", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PrivateMessageReports(ctx context.Context, d types.ListPrivateMessageReports) (*types.ListPrivateMessageReportsResponse, error) {
	ar := &types.ListPrivateMessageReportsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/private_message/report/list", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) MarkPrivateMessageAsRead(ctx context.Context, d types.MarkPrivateMessageAsRead) (*types.PrivateMessageResponse, error) {
	ar := &types.PrivateMessageResponse{}
	res, err := c.req(ctx, http.MethodPost, "/private_message/mark_as_read", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) ResolvePrivateMessageReport(ctx context.Context, d types.ResolvePrivateMessageReport) (*types.PrivateMessageReportResponse, error) {
	ar := &types.PrivateMessageReportResponse{}
	res, err := c.req(ctx, http.MethodPut, "/private_message/report/resolve", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
