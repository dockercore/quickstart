package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/models"
	"github.com/astaxie/beego/orm"

)

type LoginController struct {
	beego.Controller
}
/*
func (c *LoginController) Get(){
   c.TplName = "login.html"
}

func (c *LoginController) Post(){

	c.Ctx.WriteString("注册成功")
}*/

func (c *LoginController) ShowLogin()  {

	c.TplName = "login.html"
}

func (c *LoginController) HandleLogin()  {
  //1.拿到数据
   username := c.GetString("username")
   pwd := c.GetString("pwd")
  //2.数据是否合法
  if username == "" || pwd == ""{
  	beego.Info("输入数据不合法")
  	//c.TplName = "login.html"
  	c.Redirect("/login",302) //跳转
  	return
  }
  //3.查询账户和密码是否正确
   o := orm.NewOrm()
   user := models.User{}
   user.Name = username
   user.Pwd=pwd
  err :=  o.Read(&user,"Name","Pwd")
  if err != nil{
  	 beego.Info("查询失败")
  	 c.Redirect("/login",302) //跳转
  	 //c.TplName="login.html" 能给页面带参数 但是比Redirect慢
  	 return
  }
  //4.跳转到index页面
  c.Ctx.WriteString("成功登录")
}