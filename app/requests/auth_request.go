package requests

import "github.com/thedevsaddam/govalidator"

type RegisterRequest struct {
	Nickname string `valid:"nickname" json:"nickname"`
	Username string `valid:"username" json:"username"`
	Password string `valid:"password" json:"password"`
}

func (r *RegisterRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"nickname": []string{"required"},
		"username": []string{"required"},
		"password": []string{"required"},
	}
}

func (r *RegisterRequest) Messages() govalidator.MapData {
	return govalidator.MapData{
		"nickname": []string{
			"required:昵称不能为空",
		},
		"username": []string{
			"required:用户名不能为空",
		},
		"password": []string{
			"required:密码不能为空",
		},
	}
}

type LoginRequest struct {
	Username string `valid:"username" json:"username"`
	Password string `valid:"password" json:"password"`
}

func (r *LoginRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}
}

func (r *LoginRequest) Messages() govalidator.MapData {
	return govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}
}
