package controllers

import (
	"villa/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) NestPrepare() {
	if !this.IsLogin {
		this.Redirect("/login", 302)
	}
}

func (this *MainController) Main() {
	this.TplNames = "frame.tpl"
}

func (this *MainController) TopBody() {

	u := models.AdminModel{}
	admin_id := this.GetSession("LOGIN_SESSION_KEY")

	if admin_id_val, ok := admin_id.(uint); ok {
		/* act on str */
		u_admin, err := u.UserInfo(admin_id_val)
		if err == nil {
			this.Data["Code"] = 1
			this.Data["Name"] = u_admin.Name
		} else {
			this.Data["Code"] = 0
			this.Data["Name"] = "未知"
		}

	} else {
		/* not string */
		this.Data["Code"] = 0
		this.Data["Name"] = "未知"
	}

	this.TplNames = "topbody.tpl"
}

func (this *MainController) LeftBody() {
	this.TplNames = "leftbody.tpl"
}

func (this *MainController) RightBody() {
	this.TplNames = "rightbody.tpl"
}
