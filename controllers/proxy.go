// Package controllers provides ...
package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/deepzz0/go-com/log"
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
	response, err := http.Get("http://" + url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	this.Ctx.Output.Body(b)
}
