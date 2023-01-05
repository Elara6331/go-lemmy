package lemmy

import (
	"context"
	types "go.arsenm.dev/go-lemmy/types"
)

func (c *Client) Site(ctx context.Context, data types.GetSite) (*types.GetSiteResponse, error) {
	resData := &types.GetSiteResponse{}
	res, err := c.getReq(ctx, "GET", "/site", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreateSite(ctx context.Context, data types.CreateSite) (*types.SiteResponse, error) {
	resData := &types.SiteResponse{}
	res, err := c.req(ctx, "POST", "/site", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) EditSite(ctx context.Context, data types.EditSite) (*types.SiteResponse, error) {
	resData := &types.SiteResponse{}
	res, err := c.req(ctx, "PUT", "/site", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) SiteConfig(ctx context.Context, data types.GetSiteConfig) (*types.GetSiteConfigResponse, error) {
	resData := &types.GetSiteConfigResponse{}
	res, err := c.getReq(ctx, "GET", "/site/config", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) SaveSiteConfig(ctx context.Context, data types.SaveSiteConfig) (*types.GetSiteConfigResponse, error) {
	resData := &types.GetSiteConfigResponse{}
	res, err := c.req(ctx, "PUT", "/site/config", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Modlog(ctx context.Context, data types.GetModlog) (*types.GetModlogResponse, error) {
	resData := &types.GetModlogResponse{}
	res, err := c.getReq(ctx, "GET", "/modlog", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Search(ctx context.Context, data types.Search) (*types.SearchResponse, error) {
	resData := &types.SearchResponse{}
	res, err := c.getReq(ctx, "GET", "/search", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ResolveObject(ctx context.Context, data types.ResolveObject) (*types.ResolveObjectResponse, error) {
	resData := &types.ResolveObjectResponse{}
	res, err := c.getReq(ctx, "GET", "/resolve_object", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreateCommunity(ctx context.Context, data types.CreateCommunity) (*types.CommunityResponse, error) {
	resData := &types.CommunityResponse{}
	res, err := c.req(ctx, "POST", "/community", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Community(ctx context.Context, data types.GetCommunity) (*types.GetCommunityResponse, error) {
	resData := &types.GetCommunityResponse{}
	res, err := c.getReq(ctx, "GET", "/community", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) EditCommunity(ctx context.Context, data types.EditCommunity) (*types.CommunityResponse, error) {
	resData := &types.CommunityResponse{}
	res, err := c.req(ctx, "PUT", "/community", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) HideCommunity(ctx context.Context, data types.HideCommunity) (*types.CommunityResponse, error) {
	resData := &types.CommunityResponse{}
	res, err := c.req(ctx, "PUT", "/community/hide", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ListCommunities(ctx context.Context, data types.ListCommunities) (*types.ListCommunitiesResponse, error) {
	resData := &types.ListCommunitiesResponse{}
	res, err := c.getReq(ctx, "GET", "/community/list", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) FollowCommunity(ctx context.Context, data types.FollowCommunity) (*types.CommunityResponse, error) {
	resData := &types.CommunityResponse{}
	res, err := c.req(ctx, "POST", "/community/follow", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) BlockCommunity(ctx context.Context, data types.BlockCommunity) (*types.BlockCommunityResponse, error) {
	resData := &types.BlockCommunityResponse{}
	res, err := c.req(ctx, "POST", "/community/block", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) DeleteCommunity(ctx context.Context, data types.DeleteCommunity) (*types.CommunityResponse, error) {
	resData := &types.CommunityResponse{}
	res, err := c.req(ctx, "POST", "/community/delete", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) RemoveCommunity(ctx context.Context, data types.RemoveCommunity) (*types.CommunityResponse, error) {
	resData := &types.CommunityResponse{}
	res, err := c.req(ctx, "POST", "/community/remove", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) TransferCommunity(ctx context.Context, data types.TransferCommunity) (*types.GetCommunityResponse, error) {
	resData := &types.GetCommunityResponse{}
	res, err := c.req(ctx, "POST", "/community/transfer", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) BanFromCommunity(ctx context.Context, data types.BanFromCommunity) (*types.BanFromCommunityResponse, error) {
	resData := &types.BanFromCommunityResponse{}
	res, err := c.req(ctx, "POST", "/community/ban_user", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) AddModToCommunity(ctx context.Context, data types.AddModToCommunity) (*types.AddModToCommunityResponse, error) {
	resData := &types.AddModToCommunityResponse{}
	res, err := c.req(ctx, "POST", "/community/mod", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CommunityJoin(ctx context.Context, data types.CommunityJoin) (*types.CommunityJoinResponse, error) {
	resData := &types.CommunityJoinResponse{}
	res, err := c.req(ctx, "POST", "/community/join", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ModJoin(ctx context.Context, data types.ModJoin) (*types.ModJoinResponse, error) {
	resData := &types.ModJoinResponse{}
	res, err := c.req(ctx, "POST", "/community/mod/join", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreatePost(ctx context.Context, data types.CreatePost) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "POST", "/post", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Post(ctx context.Context, data types.GetPost) (*types.GetPostResponse, error) {
	resData := &types.GetPostResponse{}
	res, err := c.getReq(ctx, "GET", "/post", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) EditPost(ctx context.Context, data types.EditPost) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "PUT", "/post", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) DeletePost(ctx context.Context, data types.DeletePost) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "POST", "/post/delete", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) RemovePost(ctx context.Context, data types.RemovePost) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "POST", "/post/remove", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) MarkPostAsRead(ctx context.Context, data types.MarkPostAsRead) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "POST", "/post/mark_as_read", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) LockPost(ctx context.Context, data types.LockPost) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "POST", "/post/lock", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) StickyPost(ctx context.Context, data types.StickyPost) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "POST", "/post/sticky", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Posts(ctx context.Context, data types.GetPosts) (*types.GetPostsResponse, error) {
	resData := &types.GetPostsResponse{}
	res, err := c.getReq(ctx, "GET", "/post/list", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreatePostLike(ctx context.Context, data types.CreatePostLike) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "POST", "/post/like", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) SavePost(ctx context.Context, data types.SavePost) (*types.PostResponse, error) {
	resData := &types.PostResponse{}
	res, err := c.req(ctx, "PUT", "/post/save", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) PostJoin(ctx context.Context, data types.PostJoin) (*types.PostJoinResponse, error) {
	resData := &types.PostJoinResponse{}
	res, err := c.req(ctx, "POST", "/post/join", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreatePostReport(ctx context.Context, data types.CreatePostReport) (*types.PostReportResponse, error) {
	resData := &types.PostReportResponse{}
	res, err := c.req(ctx, "POST", "/post/report", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ResolvePostReport(ctx context.Context, data types.ResolvePostReport) (*types.PostReportResponse, error) {
	resData := &types.PostReportResponse{}
	res, err := c.req(ctx, "PUT", "/post/report/resolve", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ListPostReports(ctx context.Context, data types.ListPostReports) (*types.ListPostReportsResponse, error) {
	resData := &types.ListPostReportsResponse{}
	res, err := c.getReq(ctx, "GET", "/post/report/list", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) SiteMetadata(ctx context.Context, data types.GetSiteMetadata) (*types.GetSiteMetadataResponse, error) {
	resData := &types.GetSiteMetadataResponse{}
	res, err := c.getReq(ctx, "GET", "/post/site_metadata", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreateComment(ctx context.Context, data types.CreateComment) (*types.CommentResponse, error) {
	resData := &types.CommentResponse{}
	res, err := c.req(ctx, "POST", "/comment", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Comment(ctx context.Context, data types.GetComment) (*types.CommentResponse, error) {
	resData := &types.CommentResponse{}
	res, err := c.getReq(ctx, "GET", "/comment", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) EditComment(ctx context.Context, data types.EditComment) (*types.CommentResponse, error) {
	resData := &types.CommentResponse{}
	res, err := c.req(ctx, "PUT", "/comment", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) DeleteComment(ctx context.Context, data types.DeleteComment) (*types.CommentResponse, error) {
	resData := &types.CommentResponse{}
	res, err := c.req(ctx, "POST", "/comment/delete", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) RemoveComment(ctx context.Context, data types.RemoveComment) (*types.CommentResponse, error) {
	resData := &types.CommentResponse{}
	res, err := c.req(ctx, "POST", "/comment/remove", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) MarkCommentAsRead(ctx context.Context, data types.MarkCommentAsRead) (*types.CommentResponse, error) {
	resData := &types.CommentResponse{}
	res, err := c.req(ctx, "POST", "/comment/mark_as_read", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreateCommentLike(ctx context.Context, data types.CreateCommentLike) (*types.CommentResponse, error) {
	resData := &types.CommentResponse{}
	res, err := c.req(ctx, "POST", "/comment/like", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) SaveComment(ctx context.Context, data types.SaveComment) (*types.CommentResponse, error) {
	resData := &types.CommentResponse{}
	res, err := c.req(ctx, "PUT", "/comment/save", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Comments(ctx context.Context, data types.GetComments) (*types.GetCommentsResponse, error) {
	resData := &types.GetCommentsResponse{}
	res, err := c.getReq(ctx, "GET", "/comment/list", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreateCommentReport(ctx context.Context, data types.CreateCommentReport) (*types.CommentReportResponse, error) {
	resData := &types.CommentReportResponse{}
	res, err := c.req(ctx, "POST", "/comment/report", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ResolveCommentReport(ctx context.Context, data types.ResolveCommentReport) (*types.CommentReportResponse, error) {
	resData := &types.CommentReportResponse{}
	res, err := c.req(ctx, "PUT", "/comment/report/resolve", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ListCommentReports(ctx context.Context, data types.ListCommentReports) (*types.ListCommentReportsResponse, error) {
	resData := &types.ListCommentReportsResponse{}
	res, err := c.getReq(ctx, "GET", "/comment/report/list", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) PrivateMessages(ctx context.Context, data types.GetPrivateMessages) (*types.PrivateMessagesResponse, error) {
	resData := &types.PrivateMessagesResponse{}
	res, err := c.getReq(ctx, "GET", "/private_message/list", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) CreatePrivateMessage(ctx context.Context, data types.CreatePrivateMessage) (*types.PrivateMessageResponse, error) {
	resData := &types.PrivateMessageResponse{}
	res, err := c.req(ctx, "POST", "/private_message", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) EditPrivateMessage(ctx context.Context, data types.EditPrivateMessage) (*types.PrivateMessageResponse, error) {
	resData := &types.PrivateMessageResponse{}
	res, err := c.req(ctx, "PUT", "/private_message", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) DeletePrivateMessage(ctx context.Context, data types.DeletePrivateMessage) (*types.PrivateMessageResponse, error) {
	resData := &types.PrivateMessageResponse{}
	res, err := c.req(ctx, "POST", "/private_message/delete", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) MarkPrivateMessageAsRead(ctx context.Context, data types.MarkPrivateMessageAsRead) (*types.PrivateMessageResponse, error) {
	resData := &types.PrivateMessageResponse{}
	res, err := c.req(ctx, "POST", "/private_message/mark_as_read", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Register(ctx context.Context, data types.Register) (*types.LoginResponse, error) {
	resData := &types.LoginResponse{}
	res, err := c.req(ctx, "POST", "/user/register", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Captcha(ctx context.Context, data types.GetCaptcha) (*types.GetCaptchaResponse, error) {
	resData := &types.GetCaptchaResponse{}
	res, err := c.getReq(ctx, "GET", "/user/get_captcha", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) PersonDetails(ctx context.Context, data types.GetPersonDetails) (*types.GetPersonDetailsResponse, error) {
	resData := &types.GetPersonDetailsResponse{}
	res, err := c.getReq(ctx, "GET", "/user", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) PersonMentions(ctx context.Context, data types.GetPersonMentions) (*types.GetPersonMentionsResponse, error) {
	resData := &types.GetPersonMentionsResponse{}
	res, err := c.getReq(ctx, "GET", "/user/mention", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) MarkPersonMentionAsRead(ctx context.Context, data types.MarkPersonMentionAsRead) (*types.PersonMentionResponse, error) {
	resData := &types.PersonMentionResponse{}
	res, err := c.req(ctx, "POST", "/user/mention/mark_as_read", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Replies(ctx context.Context, data types.GetReplies) (*types.GetRepliesResponse, error) {
	resData := &types.GetRepliesResponse{}
	res, err := c.getReq(ctx, "GET", "/user/replies", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) UserJoin(ctx context.Context, data types.UserJoin) (*types.UserJoinResponse, error) {
	resData := &types.UserJoinResponse{}
	res, err := c.req(ctx, "POST", "/user/join", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) BanPerson(ctx context.Context, data types.BanPerson) (*types.BanPersonResponse, error) {
	resData := &types.BanPersonResponse{}
	res, err := c.req(ctx, "POST", "/user/ban", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) BannedPersons(ctx context.Context, data types.GetBannedPersons) (*types.BannedPersonsResponse, error) {
	resData := &types.BannedPersonsResponse{}
	res, err := c.getReq(ctx, "GET", "/user/banned", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) BlockPerson(ctx context.Context, data types.BlockPerson) (*types.BlockPersonResponse, error) {
	resData := &types.BlockPersonResponse{}
	res, err := c.req(ctx, "POST", "/user/block", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) Login(ctx context.Context, data types.Login) (*types.LoginResponse, error) {
	resData := &types.LoginResponse{}
	res, err := c.req(ctx, "POST", "/user/login", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) DeleteAccount(ctx context.Context, data types.DeleteAccount) (*types.DeleteAccountResponse, error) {
	resData := &types.DeleteAccountResponse{}
	res, err := c.req(ctx, "POST", "/user/delete_account", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) PasswordReset(ctx context.Context, data types.PasswordReset) (*types.PasswordResetResponse, error) {
	resData := &types.PasswordResetResponse{}
	res, err := c.req(ctx, "POST", "/user/password_reset", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) PasswordChangeAfterReset(ctx context.Context, data types.PasswordChangeAfterReset) (*types.LoginResponse, error) {
	resData := &types.LoginResponse{}
	res, err := c.req(ctx, "POST", "/user/password_change", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) MarkAllAsRead(ctx context.Context, data types.MarkAllAsRead) (*types.GetRepliesResponse, error) {
	resData := &types.GetRepliesResponse{}
	res, err := c.req(ctx, "POST", "/user/mark_all_as_read", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) SaveUserSettings(ctx context.Context, data types.SaveUserSettings) (*types.LoginResponse, error) {
	resData := &types.LoginResponse{}
	res, err := c.req(ctx, "PUT", "/user/save_user_settings", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ChangePassword(ctx context.Context, data types.ChangePassword) (*types.LoginResponse, error) {
	resData := &types.LoginResponse{}
	res, err := c.req(ctx, "PUT", "/user/change_password", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ReportCount(ctx context.Context, data types.GetReportCount) (*types.GetReportCountResponse, error) {
	resData := &types.GetReportCountResponse{}
	res, err := c.getReq(ctx, "GET", "/user/report_count", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) UnreadCount(ctx context.Context, data types.GetUnreadCount) (*types.GetUnreadCountResponse, error) {
	resData := &types.GetUnreadCountResponse{}
	res, err := c.getReq(ctx, "GET", "/user/unread_count", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) VerifyEmail(ctx context.Context, data types.VerifyEmail) (*types.VerifyEmailResponse, error) {
	resData := &types.VerifyEmailResponse{}
	res, err := c.req(ctx, "POST", "/user/verify_email", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) LeaveAdmin(ctx context.Context, data types.LeaveAdmin) (*types.GetSiteResponse, error) {
	resData := &types.GetSiteResponse{}
	res, err := c.req(ctx, "POST", "/user/leave_admin", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) AddAdmin(ctx context.Context, data types.AddAdmin) (*types.AddAdminResponse, error) {
	resData := &types.AddAdminResponse{}
	res, err := c.req(ctx, "POST", "/admin/add", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) UnreadRegistrationApplicationCount(ctx context.Context, data types.GetUnreadRegistrationApplicationCount) (*types.GetUnreadRegistrationApplicationCountResponse, error) {
	resData := &types.GetUnreadRegistrationApplicationCountResponse{}
	res, err := c.getReq(ctx, "GET", "/admin/registration_application/count", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ListRegistrationApplications(ctx context.Context, data types.ListRegistrationApplications) (*types.ListRegistrationApplicationsResponse, error) {
	resData := &types.ListRegistrationApplicationsResponse{}
	res, err := c.getReq(ctx, "GET", "/admin/registration_application/list", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
func (c *Client) ApproveRegistrationApplication(ctx context.Context, data types.ApproveRegistrationApplication) (*types.RegistrationApplicationResponse, error) {
	resData := &types.RegistrationApplicationResponse{}
	res, err := c.req(ctx, "PUT", "/admin/registration_application/approve", data, &resData)
	if err != nil {
		return nil, err
	}
	err = resError(res, resData.LemmyResponse)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
