package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"*/


/*	//1.orm 对象
	o :=orm.NewOrm()
	//2.有一个可以插入结构体的对象
	user:= models.User{}
	//3.对结构体进行赋值
	user.Name="zhangsan"
	user.Pwd="22222"
	//4.调用orm方法插入
	o.Insert(&user)*/

	/* 查询
	//1.orm 对象
	o :=orm.NewOrm()
	//2.查询的对象
	user:= models.User{}
	//3.查询对象的值
	user.Id = 1

	err := o.Read(&user)
	if err != nil{
		beego.Info("查询失败",err)
		return
	}
	beego.Info("查询成功",user)*/
/*	//修改
	//1.要有orm对象
	 o := orm.NewOrm()
	//2.需要更新的结构体对象
	user := models.User{}
	//3.查询需要更新的数据
	user.Id = 1
	err := o.Read(&user)
	//4.给数据重新赋值
	if err == nil{
		user.Name = "lisi"
		_,err := o.Update(&user)
		if err != nil{
			beego.Info("更新失败",err)
			return
		}
	}*/

	//删除操作
	 o := orm.NewOrm()
	 user := models.User{}
	 user.Id = 1
	 _,err := o.Delete(&user)
	 if err !=nil{
	 	beego.Info("删除错误",err)
	 	return
	 }

	c.Data["Website"] = "我们的第一个go程序"
	c.TplName="test.tpl" //tplName指定页面

	}

func (c *MainController) Post() {
	c.Data["Website"] = "我们的第一个go程序"
	c.TplName="index.tpl" //tplName指定页面
}



