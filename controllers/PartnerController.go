package controllers

import (
	"fmt"
	"villa/models"
)

type PartnerController struct {
	BaseController
}

func (this *PartnerController) List() {
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
	this.TplNames = "partner/add.tpl"
}
