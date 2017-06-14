package plugin

import (
	"github.com/astaxie/beego"
	"github.com/yuhaofeng110/goblog/models"
)

type Plugin struct {
	beego.Controller
}

func (this *Plugin) Prepare() {
	this.DoRequest()
}

func (this *Plugin) DoRequest() {
	requst := models.NewRequest(this.Ctx.Request)
	models.RequestM.Ch <- requst
}
