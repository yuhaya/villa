package controllers

import (
	"fmt"
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
    if err := this.ParseForm(&p_model); err != nil {
        //handle error
        fmt.Printf("%s", err.Error())
    }else{
    	fmt.Printf("%v", p_model)
    }
	this.Data["TitleName"] = "添加合作方"
	this.TplNames = "partner/add.tpl"
}