package controllers

import (
	"github.com/weibaohui/mesh-ui/constant"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["MeshServer"] = constant.MeshServer()
	c.TplName = "index.html"
}
