package controllers

import (
	"fmt"

	"github.com/deepzz0/goblog/models"
)

type AboutController struct {
	Common
}

func (this *AboutController) Get() {
	this.Layout = "homelayout.html"
	this.TplName = "aboutTemplate.html"
	this.Data["Title"] = "About Me | " + models.Blogger.BlogName
	this.Leftbar("about")
	this.Content()
}

func (this *AboutController) Content() {
	this.Data["URL"] = "/about"
	if about := models.TMgr.GetTopic(1); about != nil {
		this.Data["Content"] = string(about.Content)
	} else {
		this.Data["Content"] = "nothing。"
	}
	this.Data["Description"] = fmt.Sprintf("about me,%s,%s,blog", models.Blogger.Introduce, models.Blogger.UserName)
	this.Data["Keywords"] = fmt.Sprintf("about me,%s,%s", models.Blogger.Introduce, models.Blogger.UserName)
}
