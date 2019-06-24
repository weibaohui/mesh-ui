package controllers

type PodController struct {
	BaseController
}

func (c *PodController) Get() {
	c.TplName = "pods.html"
}
func (c *PodController) ContainerLog() {
	c.TplName = "container_log.html"
}
func (c *PodController) ContainerExec() {
	c.TplName = "container_exec.html"
}
