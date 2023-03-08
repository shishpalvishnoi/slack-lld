package models

type UserChannel struct {
	Id int64 `json:"id"`
	UserId int64 `json:"user"`
	ChannelId int64 `json:"channel_id"`
	RoleType string `json:"role_type"`
}
