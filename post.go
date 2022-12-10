package lemmy

import (
	"context"
	"net/http"

	"go.arsenm.dev/go-lemmy/types"
)

func (c *Client) CreatePost(ctx context.Context, d types.CreatePost) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPost, "/post", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) EditPost(ctx context.Context, d types.EditPost) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPut, "/post", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) Post(ctx context.Context, d types.GetPost) (*types.GetPostResponse, error) {
	ar := &types.GetPostResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/post", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) CreatePostReport(ctx context.Context, d types.CreatePostReport) (*types.PostReportResponse, error) {
	ar := &types.PostReportResponse{}
	res, err := c.req(ctx, http.MethodPost, "/post/report", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) DeletePost(ctx context.Context, d types.DeletePost) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPost, "/post/delete", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) RemovePost(ctx context.Context, d types.RemovePost) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPost, "/post/remove", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) Posts(ctx context.Context, d types.GetPosts) (*types.GetPostsResponse, error) {
	ar := &types.GetPostsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/post/list", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) SiteMetadata(ctx context.Context, d types.GetSiteMetadata) (*types.GetSiteMetadataResponse, error) {
	ar := &types.GetSiteMetadataResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/post/site_metadata", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) LikePost(ctx context.Context, d types.CreatePostLike) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPost, "/post/like", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PostReports(ctx context.Context, d types.ListPostReports) (*types.ListPostReportsResponse, error) {
	ar := &types.ListPostReportsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/post/report/list", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) LockPost(ctx context.Context, d types.LockPost) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPost, "/post/lock", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) MarkPostAsRead(ctx context.Context, d types.MarkPostAsRead) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPost, "/post/mark_as_read", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) ResolvePostReport(ctx context.Context, d types.ResolvePostReport) (*types.PostReportResponse, error) {
	ar := &types.PostReportResponse{}
	res, err := c.req(ctx, http.MethodPut, "/post/report/resolve", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) SavePost(ctx context.Context, d types.SavePost) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPut, "/post/save", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) StickyPost(ctx context.Context, d types.StickyPost) (*types.PostResponse, error) {
	ar := &types.PostResponse{}
	res, err := c.req(ctx, http.MethodPost, "/post/sticky", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
