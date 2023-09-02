package models

import "time"

type User struct {
	UserName string `json:"userName"`
	//Password string `json:"password"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	//GoogleSecret  string    `json:"googleSecret"`
	LastLoginTime time.Time `json:"lastLoginTime"`
	RegisterTime  time.Time `json:"registerTime"`
	Token         string    `json:"token"`
}
