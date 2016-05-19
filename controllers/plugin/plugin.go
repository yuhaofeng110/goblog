package plugin

import (
	"github.com/astaxie/beego"
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
}
