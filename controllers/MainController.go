package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Main() {
	if c.IsLogin {
		c.Data["Website"] = "beego.me222"
		c.Data["Email"] = "astaxie@gmail.com222222"
		c.TplNames = "index.tpl"
	} else {
		c.Redirect("/login", 302)
	}
}
