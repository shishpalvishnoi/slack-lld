package models

type RegUserReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
}

type CreateWorkspaceReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// created_by
}

type CreateChannelReq struct {
	WorkspaceId int64  `json:"workspace_id"`
	Name        string `json:"name"`
	CreatedBy   string `json:"created_by"`
}

type AddUserToWorkspaceReq struct {
	UserId      int64  `json:"user_id"`
	WorkSpaceId int64  `json:"work_space_id"`
	RoleType    string `json:"role_type"`
}

type AddUserToChannelReq struct {
	UserId    int64  `json:"user"`
	ChannelId int64  `json:"channel_id"`
	RoleType  string `json:"role_type"`
}

type SendMessageToChReq struct {
	Content     string `json:"content"`
	Attachment  string `json:"attachment"`
	UserId      int64  `json:"user_id"`
	WorkspaceId int64  `json:"workspace_id"`
	ChannelId   int64  `json:"channel_id"`
}

type SendMessageToUserReq struct {
	Content        string `json:"content"`
	Attachment     string `json:"attachment"`
	UserId         int64  `json:"user_id"`
	WorkspaceId    int64  `json:"workspace_id"`
	ReceiverUserId int64  `json:"receiver_user_id"`
}
