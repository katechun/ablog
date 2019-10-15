package main

import (
	"ablog/helpers"
	"ablog/router"
)




func main() {

	//初始化命令行配置
	helpers.InitConfig()

	//初始化数据
	helpers.InitDB()

	//初始化日志
	helpers.InitLog()

	//初始化路由
	r := router.InitRouter()

    //开启日志记录调用
	helpers.Logger.Println("test")

	//启动监听
	err :=r.Run(helpers.GetOneValue("app","host")+":"+helpers.GetOneValue("app","port"))
	if err != nil {
		helpers.Logger.Println("启动监听错误！")
	}



}
