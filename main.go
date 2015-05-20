package main

import (
	_ "LMS/controllers"
	_ "LMS/models"
	_ "LMS/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/deferpanic/deferclient/deferstats"
)

func init() {
	//设置log格式
	beego.SetLogger("file", `{"filename":"logs/server.log"}`)
	beego.SetLogFuncCall(true)
	beego.SetLevel(beego.LevelDebug)

	//设置下载脚本目录
	beego.SetStaticPath("/script_download", "script_download")
}

func main() {
	//开启调试模式
	orm.Debug = true

	//自动同步数据库表格
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		beego.Critical("sycndb error! Error:", err)
	}

	//开启defer panic支持
	//deferstats.NewClient("kxHlEw0EeO5OQj4GNqIG58jsE81p2356")

	//启动服务
	beego.Trace("LMS start running...")
	beego.Run()
}
