package controllers

import (
	"villa/models"
)

type UserController struct {
	BaseController
}

func (u *UserController) LoginPage() {
	if u.isLogin {
		u.Redirect("/", 302)
	} else {
		admin_mode := models.AdminModel{Name: "yuhaya", Email: "3wmaocomputer@gmail.com"}
		admin_mode.GetSession()
		u.TplNames = "user/login.tpl"
	}
}

func (u *UserController) LoginSubmit() {
	username := u.GetString("username")
	password := u.GetString("password")
	if username == "yuhaya" && password == "123" {
		u.isLogin = true

		var userid int = 1
		u.SetSession("LOGIN_SESSION_KEY", userid)

		u.Redirect("/", 302)
	} else {
		u.Ctx.WriteString("验证失败!")
	}
	return
}
