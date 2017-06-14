package main

import (
	"github.com/astaxie/beego"
	_ "github.com/yuhaofeng110/goblog/routers"
	_ "github.com/yuhaofeng110/goblog/cron"
)

func main() {
	beego.Run()
}
