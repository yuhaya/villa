package controllers

type MainController struct {
	BaseController
}

func (c *MainController) NestPrepare() {
	if !c.IsLogin {
		c.Redirect("/login", 302)
	}
}

func (c *MainController) Main() {
	c.TplNames = "frame.tpl"
}

func (c *MainController) TopBody() {
	c.TplNames = "index.tpl"
}

func (c *MainController) LeftBody() {
	c.TplNames = "index.tpl"
}

func (c *MainController) RightBody() {
	c.TplNames = "index.tpl"
}
