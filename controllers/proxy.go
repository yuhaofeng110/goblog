// Package controllers provides ...
package controllers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/cache"
)

type ProxyController struct {
	Common
}

func (this *ProxyController) Get() {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			this.Ctx.WriteString(err.Error())
		}
	}()
	url := this.Ctx.Input.Param(":url")
	if icon := cache.Cache.Icons[url]; icon != nil {
		icon.Time = time.Now()
		this.Ctx.Output.Body(icon.Data)
		return
	}
	response, err := http.Get("http://" + url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	cache.Cache.Icons[url] = &cache.Icon{Data: b, Time: time.Now()}
	this.Ctx.Output.Body(b)
}
