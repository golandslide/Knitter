package routers

import (
	"github.com/astaxie/beego"

	"github.com/HyperNetworks/Knitter/knitter-monitor/controllers"
)

func init() {
	beego.Router("/api/v1/pods/:podns/:podname", &controllers.PodController{})
}
