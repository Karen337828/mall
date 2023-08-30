package models

import "time"

type User struct {
	UserName      string    `json:"userName"`
	Age           int       `json:"age"`
	NickName      string    `json:"nickName"`
	LastLoginTime time.Time `json:"lastLoginTime"`
}
