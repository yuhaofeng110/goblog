package background

import (
	"github.com/yuhaofeng110/goblog/controllers"
	"github.com/yuhaofeng110/goblog/models"
)

type MessageController struct {
	Common
}

func (this *MessageController) Get() {
	this.TplName = "manage/message.html"
	this.Data["Title"] = "留言管理 | " + models.Blogger.BlogName
	this.LeftBar("message")
	this.Content()
}

func (this *MessageController) Content() {
	this.Data["ID"] = 99999
	this.Data["URL"] = "/message"
	this.Data["Domain"] = controllers.Domain
	this.Data["Title"] = "给我留言"
}
