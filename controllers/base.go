package controllers

import (
	"github.com/astaxie/beego"
	"github.com/weibaohui/mesh-ui/constant"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Data["MeshServer"] = constant.MeshServer()
}
