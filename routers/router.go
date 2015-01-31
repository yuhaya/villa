package routers

import (
	"github.com/astaxie/beego"
	"villa/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Main")
	beego.Router("/frame/top", &controllers.MainController{}, "get:TopBody")
	beego.Router("/frame/left", &controllers.MainController{}, "get:LeftBody")
	beego.Router("/frame/right", &controllers.MainController{}, "get:RightBody")

	beego.Router("/login", &controllers.UserController{}, "get:LoginPage;post:LoginSubmit")
	beego.Router("/loginout", &controllers.UserController{}, "*:LoginOut")

	beego.Router("/partner/list", &controllers.PartnerController{}, "get:List")
	beego.Router("/partner", &controllers.PartnerController{}, "get:Add")

}
