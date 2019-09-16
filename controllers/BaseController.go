package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	ControllerName string
    ActionName     string
    TplNames       string
}