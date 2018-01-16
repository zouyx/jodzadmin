package management

import (
	"github.com/astaxie/beego"
	"github.com/zouyx/jodzadmin/service/zk"
	"fmt"
)

type SynchronizerController struct {
	beego.Controller
}

func (c *SynchronizerController) Get() {
	conn:= zk.GetConn()
	strings, stat, e := conn.Children("/jodz/jobScheduler")
	fmt.Println("stat:",stat)
	fmt.Println("e:",e)
	fmt.Println("strings:",strings)

	c.TplName = "synchronizer.tpl"
}
