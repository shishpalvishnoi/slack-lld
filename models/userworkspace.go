package models

type UserWorkspace struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	WorkSpaceId int64  `json:"work_space_id"`
	RoleType    string `json:"role_type"`
}
