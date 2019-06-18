package main

import (
	"github.com/astaxie/beego"
	_ "github.com/weibaohui/mesh-ui/routers"
)

func main() {
	beego.Run("0.0.0.0:8089")
}
