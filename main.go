package main

import (
	"github.com/astaxie/beego"
	"github.com/deepzz0/go-com/log"
	_ "github.com/deepzz0/goblog/routers"
)

func main() {
	beego.Run()
	log.WaitFlush()
}
