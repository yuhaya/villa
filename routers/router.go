package routers

import (
	"github.com/astaxie/beego"
	"villa/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Main")
	beego.Router("/login", &controllers.UserController{}, "get:LoginPage;post:LoginSubmit")
}
