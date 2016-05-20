package plugin

import (
	"github.com/astaxie/beego"
	"github.com/deepzz0/goblog/models"
)

type Plugin struct {
	beego.Controller
	domain string
}

func (this *Plugin) Prepare() {
	this.domain = beego.AppConfig.String("mydomain")
	if beego.BConfig.RunMode == beego.DEV {
		this.domain = this.domain + ":" + beego.AppConfig.String("httpport")
	}
	this.DoRequest()
}

func (this *Plugin) DoRequest() {
	requst := models.NewRequest(this.Ctx.Request)
	models.RequestM.Ch <- requst
}
