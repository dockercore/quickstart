package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
   c.Ctx.WriteString("这个第一个自己的控制器")
}

func (c *TestController) ShowIndex()  {
   c.Ctx.WriteString("showIndex方法")
}


func (c *TestController) Post() {
	c.Ctx.WriteString("这个第一个自己的控制器 post方法")
}
