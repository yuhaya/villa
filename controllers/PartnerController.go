package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"villa/models"
)

type PartnerController struct {
	BaseController
}

func (this *PartnerController) List() {
	this.Data["TitleName"] = "合作方管理"
	status, err := this.GetInt8("state")
	if err != nil {
		status = 1
	}
	partener_model := models.PartnerModel{}
	partners, num, errl := partener_model.ListByStatus(byte(status))
	this.Data["partners"] = partners
	this.Data["num"] = num
	this.Data["state"] = status
	if errl == nil {

		this.Data["err"] = ""
	} else {

		this.Data["err"] = errl.Error()
	}
	for index, value := range partners {
		fmt.Printf("arr[%d]=%d xx\n", index, value.Name)
	}
	this.TplNames = "partner/list.tpl"
}

func (this *PartnerController) Add() {
	this.Data["TitleName"] = "添加合作方"
	this.TplNames = "partner/add.tpl"
}

func (this *PartnerController) AddSubmit() {

	p_model := models.Partner{}
	o := orm.NewOrm()

	if err := this.ParseForm(&p_model); err != nil {
		//handle error
		urlmsg := make(map[string]string)
		urlmsg["返回上一页"] = "javascript:history.go(-1)"
		urlmsg["回到列表页"] = "/partner/list"
		urlmsg["重新添加一个"] = "/partner"

		this.OutputMsg(err.Error(), urlmsg)
	} else {
		p_model.CreatedDate = time.Now().Format("2013-08-11")

		urlmsg := make(map[string]string)
		urlmsg["返回上一页"] = "javascript:history.go(-1)"
		urlmsg["回到列表页"] = "/partner/list"
		urlmsg["再添加一个"] = "/partner"
		_, err := o.Insert(&p_model)
		if err == nil {
			this.OutputMsg("添加成功", urlmsg)
		} else {
			this.OutputMsg(err.Error(), urlmsg)
		}
	}
	this.Data["TitleName"] = "添加合作方"
	this.TplNames = "partner/add.tpl"
}
