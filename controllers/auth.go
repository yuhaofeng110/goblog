package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Get() {
	if beego.BConfig.Listen.EnableHTTPS && this.Ctx.Input.Scheme() == "http" {
		this.Redirect(fmt.Sprintf("%s%s", domain, this.Ctx.Input.URL()), 301)
	}
	if logout := this.GetString("logout"); logout == "now" {
		this.DelSession(sessionname)
	} else if val, ok := this.GetSession(sessionname).(string); ok && val != "" {
		this.Redirect("/admin/data", 302)
	}
	this.TplName = "login.html"
	this.Data["Name"] = models.Blogger.BlogName
	this.Data["URL"] = "/"
}

func (this *AuthController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	username := this.GetString("username")
	password := this.GetString("password")

	if username == "" || password == "" {
		resp.Status = RS.RS_params_error
		resp.Tips(helper.WARNING, RS.RS_params_error)
		resp.WriteJson(this.Ctx.ResponseWriter)
		return
	}
	if code := models.UMgr.Login(username, password); code == RS.RS_user_inexistence {
		resp.Status = code
		resp.Tips(helper.WARNING, code)
	} else if code == RS.RS_password_error {
		resp.Status = code
		resp.Tips(helper.WARNING, code)
	} else {
		models.Blogger.LoginIp = this.Ctx.Request.RemoteAddr
		this.SetSession(sessionname, username)
		resp.Data = "/admin/data"
	}
}
