package management

import "github.com/astaxie/beego"

type SynchronizerController struct {
	beego.Controller
}

func (c *SynchronizerController) Get() {
	c.TplName = "synchronizer.tpl"
}
