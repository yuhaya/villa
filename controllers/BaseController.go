package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type AjaxReturn struct {
	Code bool
	Msg  string
	Data interface{}
}

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
	IsLogin bool
}

// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {

	v := this.GetSession("LOGIN_SESSION_KEY")
	if v == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
	}

	// page start time
	this.Data["PageStartTime"] = time.Now()

	// Setting properties.
	this.Data["AppDescription"] = beego.AppConfig.String("AppDescription")
	this.Data["AppKeywords"] = beego.AppConfig.String("AppKeywords")
	this.Data["AppName"] = beego.AppConfig.String("AppName")
	this.Data["AppVer"] = beego.AppConfig.String("AppVer")
	this.Data["AppUrl"] = beego.AppConfig.String("AppUrl")
	this.Data["AppLogo"] = beego.AppConfig.String("AppLogo")
	this.Data["IsProMode"] = beego.AppConfig.String("IsProMode")

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (this *BaseController) AjaxReturnFun(code bool, msg string, data interface{}) {
	m := AjaxReturn{code, msg, data}
	b, err := json.Marshal(m)
	if err == nil {
		this.Ctx.WriteString(string(b))
	} else {
		this.Ctx.WriteString("{\"code\":0,\"msg\":\"系统异常\",\"data\":\"\"}")
	}
}
