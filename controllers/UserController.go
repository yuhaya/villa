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
		u.TplNames = "user/login.tpl"
	}
}

func (u *UserController) LoginSubmit() {
	username := u.GetString("username")
	password := u.GetString("password")
	admin_model := models.AdminModel{}
	code, msg, uid := admin_model.Login(username, password)

	if code == false {
		u.Ctx.WriteString(msg.Error())
	} else {
		u.isLogin = true
		u.SetSession("LOGIN_SESSION_KEY", uid)
		u.Redirect("/", 302)
	}
}
