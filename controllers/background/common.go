package background

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/cache"
	"github.com/deepzz0/goblog/helper"
	// "github.com/deepzz0/goblog/models"
	// "github.com/deepzz0/go-com/log"
)

var sessionname = beego.AppConfig.String("sessionname")
var domain = beego.AppConfig.String("mydomain")

type Common struct {
	beego.Controller
	index string
}

func (this *Common) Prepare() {
	this.Layout = "manage/adminlayout.html"
}
func (this *Common) LeftBar(index string) {
	this.Data["Choose"] = index
	this.Data["LeftBar"] = cache.Cache.BackgroundLeftBars
}

// ----------------------------- 过滤登录 -----------------------------
var FilterUser = func(ctx *context.Context) {
	val, ok := ctx.Input.Session(sessionname).(string)
	if !ok || val == "" {
		if ctx.Request.Method == "GET" {
			ctx.Redirect(302, "/login")
		} else if ctx.Request.Method == "POST" {
			resp := helper.NewResponse()
			resp.Status = RS.RS_user_not_login
			resp.Data = "/login"
			resp.WriteJson(ctx.ResponseWriter)
		}
	}
}
