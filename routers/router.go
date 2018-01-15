package routers

import (
	"github.com/zouyx/jodzadmin/controllers"
	"github.com/astaxie/beego"
	"github.com/zouyx/jodzadmin/controllers/management"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/management/synchronizer", &management.SynchronizerController{})
}
