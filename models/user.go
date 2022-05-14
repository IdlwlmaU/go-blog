package models

type User struct {
	Uid      int    `json:"uid" db:"uid"`
	UserName string `json:"user_name" db:"user_name"`
	Passwd   string `json:"passwd" db:"passwd"`
	Avatar   string `json:"avatar" db:"avatar"`
	CreateAt string `json:"create_at" db:"create_at"`
	UpdateAt string `json:"update_at" db:"update_at"`
}

type UserInfo struct {
	Uid      int    `json:"uid" db:"uid"`
	UserName string `json:"userName" db:"user_name"`
	Avatar   string `json:"avatar" db:"avatar"`
}
