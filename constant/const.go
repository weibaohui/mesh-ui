package constant

import "github.com/astaxie/beego"

func Inject(c *beego.Controller) {

}

func MeshServer() string {
	return beego.AppConfig.String("meshServerAddress")
}
