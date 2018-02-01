package management

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/zouyx/jodzadmin/service/zk"
)

const SOURCE_PATH  ="/jodz/jobScheduler"

type SynchronizerController struct {
	beego.Controller
}

func (this *SynchronizerController) Get() {
	path := this.GetString("path")
	if path=="" {
		path=SOURCE_PATH
	}
	names := getInfo(path)
	if names!=nil{
		this.Data["names"] = names
	}
	this.TplName = "synchronizer.tpl"
}

func (this *SynchronizerController) GetConfig() {
	path := this.Ctx.Input.Param(":path")
	if path=="" {
		path=SOURCE_PATH
	}
	logs.Info("path:",path)

	this.TplName = "synchronizerConfig.tpl"
}

func (this *SynchronizerController) GetNodes() {
	path := this.Ctx.Input.Param(":path")
	if path=="" {
		path=SOURCE_PATH
	}
	logs.Info("path:",path)

	this.TplName = "synchronizerNodes.tpl"
}

/**
根据传入的path获取信息
 */
func getInfo(path string) []string {
	conn := zk.GetConn()
	strings, _, e := conn.Children(path)
	if e != nil {
		logs.Error("get config fail!", e)
		return nil
	}

	return strings
}
