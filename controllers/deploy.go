package controllers

type DeployController struct {
	BaseController
}

func (c *DeployController) Get() {
	c.TplName = "deploys.html"
}
