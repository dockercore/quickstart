package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
)

type RegController struct {
	beego.Controller
}

func (c *RegController) Get(){
   c.TplName = "register.html"
}

func (c *RegController) Post(){
	//1.拿到数据
	username := c.GetString("userName") //用户名
	pwd := c.GetString("pwd")
	//2.对数据进行校验
	if username == "" || pwd == ""{
		beego.Info("数据不能为空")
		c.Redirect("/reg",302) //跳转
		return
	}
	//3.插入数据库
	/*beego.Info("username=",username)
	beego.Info("pwd=",pwd)*/
	o := orm.NewOrm()
	user := models.User{}
	user.Name = username
	user.Pwd =pwd
	_,err := o.Insert(&user)
	if err != nil{
		beego.Info("插入数据失败")
		c.Redirect("/reg",302)
		return
	}

	c.Ctx.WriteString("注册成功")
}