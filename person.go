package lemmy

import (
	"context"
	"net/http"

	"go.arsenm.dev/go-lemmy/types"
)

func (c *Client) BanPerson(ctx context.Context, d types.BanPerson) (*types.BanPersonResponse, error) {
	ar := &types.BanPersonResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/ban", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) BlockPerson(ctx context.Context, d types.BlockPerson) (*types.BlockPersonResponse, error) {
	ar := &types.BlockPersonResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/block", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) ChangePassword(ctx context.Context, d types.ChangePassword) (*types.LoginResponse, error) {
	ar := &types.LoginResponse{}
	res, err := c.req(ctx, http.MethodPut, "/user/change_password", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	c.token = ar.JWT.MustValue()

	return ar, nil
}

func (c *Client) DeleteAccount(ctx context.Context, d types.DeleteAccount) (*types.DeleteAccountResponse, error) {
	ar := &types.DeleteAccountResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/delete_account", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) BannedPersons(ctx context.Context, d types.GetBannedPersons) (*types.BannedPersonsResponse, error) {
	ar := &types.BannedPersonsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/user/banned", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) Captcha(ctx context.Context, d types.GetCaptcha) (*types.CaptchaResponse, error) {
	ar := &types.CaptchaResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/user/get_captcha", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PersonDetails(ctx context.Context, d types.GetPersonDetails) (*types.GetPersonDetailsResponse, error) {
	ar := &types.GetPersonDetailsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/user", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PersonMentions(ctx context.Context, d types.GetPersonMentions) (*types.GetPersonMentionsResponse, error) {
	ar := &types.GetPersonMentionsResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/user/mention", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) Replies(ctx context.Context, d types.GetReplies) (*types.GetRepliesResponse, error) {
	ar := &types.GetRepliesResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/user/replies", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) ReportCount(ctx context.Context, d types.GetReportCount) (*types.GetReportCountResponse, error) {
	ar := &types.GetReportCountResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/user/report_count", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) UnreadCount(ctx context.Context, d types.GetUnreadCount) (*types.GetUnreadCountResponse, error) {
	ar := &types.GetUnreadCountResponse{}
	res, err := c.getReq(ctx, http.MethodGet, "/user/unread_count", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) LeaveAdmin(ctx context.Context, d types.LeaveAdmin) (*types.GetSiteResponse, error) {
	ar := &types.GetSiteResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/leave_admin", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) MarkAllAsRead(ctx context.Context, d types.MarkAllAsRead) (*types.GetRepliesResponse, error) {
	ar := &types.GetRepliesResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/mark_all_as_read", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) MarkCommentReplyAsRead(ctx context.Context, d types.MarkCommentReplyAsRead) (*types.CommentResponse, error) {
	ar := &types.CommentResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/mark_as_read", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) MarkPersonMentionAsRead(ctx context.Context, d types.MarkPersonMentionAsRead) (*types.PersonMentionResponse, error) {
	ar := &types.PersonMentionResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/mention/mark_as_read", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) PasswordChange(ctx context.Context, d types.PasswordChange) (*types.LoginResponse, error) {
	ar := &types.LoginResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/password_change", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	c.token = ar.JWT.MustValue()

	return ar, nil
}

func (c *Client) PasswordReset(ctx context.Context, d types.PasswordReset) (*types.PasswordResetResponse, error) {
	ar := &types.PasswordResetResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/password_reset", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (c *Client) Register(ctx context.Context, d types.Register) (*types.LoginResponse, error) {
	ar := &types.LoginResponse{}
	res, err := c.req(ctx, http.MethodPost, "/user/register", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	c.token = ar.JWT.MustValue()

	return ar, nil
}

func (c *Client) SaveUserSettings(ctx context.Context, d types.SaveUserSettings) (*types.LoginResponse, error) {
	ar := &types.LoginResponse{}
	res, err := c.req(ctx, http.MethodPut, "/user/save_user_settings", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	c.token = ar.JWT.MustValue()

	return ar, nil
}

func (c *Client) VerifyEmail(ctx context.Context, d types.VerifyEmail) (*types.VerifyEmailResponse, error) {
	ar := &types.VerifyEmailResponse{}
	res, err := c.req(ctx, http.MethodPut, "/user/verify_email", d, &ar)
	if err != nil {
		return nil, err
	}

	err = resError(res, ar.LemmyResponse)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
