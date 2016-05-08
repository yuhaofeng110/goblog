package controllers

import (
	"time"

	"github.com/astaxie/beego"
	// "github.com/deepzz0/go-common/log"
	"github.com/deepzz0/goblog/cache"
	"github.com/deepzz0/goblog/models"
)

var sessionname = beego.AppConfig.String("sessionname")

type Common struct {
	beego.Controller
	domain string
	url    string
}

func (this *Common) Prepare() {
	this.url = this.Ctx.Request.URL.String()
	this.domain = beego.AppConfig.String("mydomain")
	if beego.BConfig.RunMode == beego.DEV {
		this.domain = this.domain + ":" + beego.AppConfig.String("httpport")
	}
	this.Layout = "homelayout.html"
	this.Build()
}
func (this *Common) Leftbar(cat string) {
	this.Data["Picture"] = models.Blogger.HeadIcon
	this.Data["BlogName"] = models.Blogger.BlogName
	this.Data["Introduce"] = models.Blogger.Introduce
	this.Data["Categories"] = models.Blogger.Categories
	this.Data["Socials"] = models.Blogger.Socials
	this.Data["Domain"] = this.domain
	this.Data["Choose"] = cat
	this.Data["CopyTime"] = time.Now().Year()
}

func (this *Common) Build() {
	this.Data["Build"] = cache.Cache.BuildVersion
}

func (this *Common) Verification() {
	this.Data["Verification"] = models.ManageConf.SiteVerify
}

type listOfTopic struct {
	ID        int32
	Title     string
	URL       string
	Time      string
	Preview   string
	PCategory *models.Category
	PTags     []*models.Tag
}
