package main

import (
	_ "quickstart/routers"  //给roouters包init方法做初始化
	"github.com/astaxie/beego"
	_ "quickstart/models"
)

func main() {
	beego.Run()
}

