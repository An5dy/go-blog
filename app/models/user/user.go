package user

import "go-blog/app/models"

type User struct {
	models.PrimaryKey
	Nickname string `gorm:"column:nickname;type:varchar(255);not null;default:;comment:昵称;" json:"nickname"`
	Username string `gorm:"column:username;type:varchar(255);uniqueIndex:uk_idx_username;not null;default:;comment:用户名;" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);not null;default:;comment:密码;" json:"-"`
	Token    string `gorm:"column:token;type:text;comment:token;" json:"token"`
	models.Timestamps
}

type Token struct {
	Token string `json:"token"`
}
