package models

type Message struct {
	Id         int64  `json:"id"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
	UserId int64 `json:"user_id"`
	WorkspaceId int64 `json:"workspace_id"`
	ChannelId   int64  `json:"channel_id"`
}
