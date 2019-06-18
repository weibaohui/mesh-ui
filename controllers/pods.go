package controllers

type PodController struct {
	BaseController
}

func (c *PodController) Get() {
	c.TplName = "pods.html"
}
