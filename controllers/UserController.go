package controllers

import (
	"fmt"
	"villa/models"
)

var _ = fmt.Sprint("yuahya")

type UserController struct {
	BaseController
}

func (this *UserController) LoginPage() {
	if this.IsLogin {
		this.Redirect("/", 302)
	} else {
		this.TplNames = "user/login.tpl"
	}
}

func (this *UserController) LoginSubmit() {
	username := this.GetString("username")
	password := this.GetString("password")
	admin_model := models.AdminModel{}
	code, msg, uid := admin_model.Login(username, password)

	if code == false {
		this.AjaxReturnFun(false, msg.Error(), nil)
	} else {
		this.IsLogin = true
		this.SetSession("LOGIN_SESSION_KEY", uid)
		this.AjaxReturnFun(true, "", nil)
	}
}
