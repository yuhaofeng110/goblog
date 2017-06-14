package background

import (
	"github.com/yuhaofeng110/goblog/models"
)

type ADController struct {
	Common
}

func (this *ADController) Get() {
	this.TplName = "manage/ad/adTemplate.html"
	this.Data["Title"] = "广告管理 | " + models.Blogger.BlogName
	this.LeftBar("ad")
	this.Content()
}

func (this *ADController) Content() {
	this.Data["Content"] = ""
}
