package types

type CreatePrivateMessage struct {
	Content     string           `json:"content,omitempty" url:"content,omitempty"`
	RecipientID int              `json:"recipient_id,omitempty" url:"recipient_id,omitempty"`
	Auth        Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type EditPrivateMessage struct {
	PrivateMessageID int              `json:"private_message_id,omitempty" url:"private_message_id,omitempty"`
	Content          string           `json:"content,omitempty" url:"content,omitempty"`
	Auth             Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type DeletePrivateMessage struct {
	PrivateMessageID int              `json:"private_message_id,omitempty" url:"private_message_id,omitempty"`
	Deleted          bool             `json:"deleted,omitempty" url:"deleted,omitempty"`
	Auth             Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type MarkPrivateMessageAsRead struct {
	PrivateMessageID int              `json:"private_message_id,omitempty" url:"private_message_id,omitempty"`
	Read             bool             `json:"read,omitempty" url:"read,omitempty"`
	Auth             Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type GetPrivateMessages struct {
	UnreadOnly Optional[bool]   `json:"unread_only,omitempty" url:"unread_only,omitempty"`
	Page       Optional[int64]  `json:"page,omitempty" url:"page,omitempty"`
	Limit      Optional[int64]  `json:"limit,omitempty" url:"limit,omitempty"`
	Auth       Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type PrivateMessagesResponse struct {
	PrivateMessages []PrivateMessageView `json:"private_messages,omitempty" url:"private_messages,omitempty"`
	LemmyResponse
}

type PrivateMessageResponse struct {
	PrivateMessageView PrivateMessageView `json:"private_message_view,omitempty" url:"private_message_view,omitempty"`
	LemmyResponse
}

type CreatePrivateMessageReport struct {
	PrivateMessageID int              `json:"private_message_id,omitempty" url:"private_message_id,omitempty"`
	Reason           string           `json:"reason,omitempty" url:"reason,omitempty"`
	Auth             Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type PrivateMessageReportResponse struct {
	PrivateMessageReportView PrivateMessageReportView `json:"private_message_report_view,omitempty" url:"private_message_report_view,omitempty"`
	LemmyResponse
}

type ResolvePrivateMessageReport struct {
	ReportID int              `json:"report_id,omitempty" url:"report_id,omitempty"`
	Resolved bool             `json:"resolved,omitempty" url:"resolved,omitempty"`
	Auth     Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type ListPrivateMessageReports struct {
	Page           Optional[int64]  `json:"page,omitempty" url:"page,omitempty"`
	Limit          Optional[int64]  `json:"limit,omitempty" url:"limit,omitempty"`
	UnresolvedOnly Optional[bool]   `json:"unresolved_only,omitempty" url:"unresolved_only,omitempty"`
	Auth           Optional[string] `json:"auth,omitempty" url:"auth,omitempty"`
}

type ListPrivateMessageReportsResponse struct {
	PrivateMessageReports []PrivateMessageReportView `json:"private_message_reports,omitempty" url:"private_message_reports,omitempty"`
	LemmyResponse
}
