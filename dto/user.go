package dto

import "time"

// UserProfile 用户信息
type UserProfile struct {
	ID          string    `json:"id"`                    // ID 用户id
	Username    string    `json:"username,omitempty"`    // Username 用户名
	Email       string    `json:"email,omitempty"`       // Email 邮箱
	IsAdmin     bool      `json:"isAdmin,omitempty"`     // IsAdmin 是否为管理员
	IsAbleLogin bool      `json:"isAbleLogin,omitempty"` // IsAbleLogin 是否可以登录
	CreateTime  time.Time `json:"createTime,omitempty"`  // CreateTime 创建时间
}
