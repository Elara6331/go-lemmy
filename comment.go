package lemmy

import (
	"context"
	"net/http"

	"go.arsenm.dev/go-lemmy/types"
)

func (c *Client) Comments(ctx context.Context, d types.GetComments) (*types.GetCommentsResponse, error) {
	ar := &types.GetCommentsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/comment/list", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) CreateComment(ctx context.Context, d types.CreateComment) (*types.CommentResponse, error) {
	ar := &types.CommentResponse{}
	res, err := c.req(ctx, http.MethodPost, "/comment", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) CreateCommentReport(ctx context.Context, d types.CreateCommentReport) (*types.CommentReportResponse, error) {
	ar := &types.CommentReportResponse{}
	res, err := c.req(ctx, http.MethodPost, "/comment/report", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) DeleteComment(ctx context.Context, d types.DeleteComment) (*types.CommentResponse, error) {
	ar := &types.CommentResponse{}
	res, err := c.req(ctx, http.MethodPost, "/comment/delete", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) RemoveComment(ctx context.Context, d types.RemoveComment) (*types.CommentResponse, error) {
	ar := &types.CommentResponse{}
	res, err := c.req(ctx, http.MethodPost, "/comment/remove", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) EditComment(ctx context.Context, d types.EditComment) (*types.CommentResponse, error) {
	ar := &types.CommentResponse{}
	res, err := c.req(ctx, http.MethodPut, "/comment", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) LikeComment(ctx context.Context, d types.CreateCommentLike) (*types.CommentResponse, error) {
	ar := &types.CommentResponse{}
	res, err := c.req(ctx, http.MethodPost, "/comment/like", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) ListCommentReports(ctx context.Context, d types.ListCommentReports) (*types.ListCommentReportsResponse, error) {
	ar := &types.ListCommentReportsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/comments/report/list", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) MarkCommentReplyRead(ctx context.Context, d types.MarkCommentReplyAsRead) (*types.CommentResponse, error) {
	ar := &types.CommentResponse{}
	res, err := c.req(ctx, http.MethodPost, "/comment/mark_as_read", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) SaveComment(ctx context.Context, d types.SaveComment) (*types.CommentResponse, error) {
	ar := &types.CommentResponse{}
	res, err := c.req(ctx, http.MethodPut, "/comment/save", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
