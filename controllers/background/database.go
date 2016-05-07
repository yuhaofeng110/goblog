package background

import (
	"github.com/deepzz0/goblog/models"
)

type DataBaseController struct {
	Common
}

func (this *DataBaseController) Get() {
	this.TplName = "manage/database/databaseTemplate.html"
	this.Data["Title"] = "基础数据 - " + models.Blogger.BlogName
	this.LeftBar("data")
	this.Content()
}

func (this *DataBaseController) Content() {

}
