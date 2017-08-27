package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	appName := beego.AppConfig.String("appname")
	httpPort := beego.AppConfig.String("httpport")
	runMode := beego.AppConfig.String("runmode")

	// 刷新 http;//localhost:8080 后台打印出：appname:  myApp  httpport:  8080  runmode:  dev
	fmt.Println("appname: ", appName, " httpport: ", httpPort, " runmode: ", runMode)

	fmt.Println("==========")

	beego.SetLevel(beego.LevelInformational)
	beego.Trace("trace 1")  // 未被输出，因为trace级别小于info级别
	beego.Info("info 1")
	beego.Trace("trace 2")  // 未被输出，因为trace级别小于info级别
	beego.Info("info 2")


}
