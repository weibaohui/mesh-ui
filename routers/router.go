package routers

import (
	"github.com/astaxie/beego"
	"github.com/weibaohui/mesh-ui/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/pods", &controllers.PodController{})
	beego.Router("/containerLog", &controllers.PodController{}, "*:ContainerLog")
	beego.Router("/containerExec", &controllers.PodController{}, "*:ContainerExec")
}
