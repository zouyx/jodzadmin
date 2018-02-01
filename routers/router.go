package routers

import (
	"github.com/astaxie/beego"
	"github.com/zouyx/jodzadmin/controllers"
	"github.com/zouyx/jodzadmin/controllers/management"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/management/synchronizer/?:path", &management.SynchronizerController{})
	beego.Router("/management/synchronizer/config/:path", &management.SynchronizerController{},"get:GetConfig")
	beego.Router("/management/synchronizer/nodes/:path", &management.SynchronizerController{},"get:GetNodes")
}
