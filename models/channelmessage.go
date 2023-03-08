package models

type ChannelMessage struct {
	Id int64 `json:"id"`
	MessageIds []int64 `json:"messages"`
}
