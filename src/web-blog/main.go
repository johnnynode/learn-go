package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller // 在beego 里面默认实现了一些方法,如get, post
}

func (h *MainController) Get() {
	h.Ctx.WriteString("Hello World")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}

