package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    /*beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})
     //t1访问的get方式访问走的是Get方法 Post请求走post方法
	//beego.Router("/t1",&controllers.TestController{})
	//t1访问的get方式访问走的是ShowIndex
	beego.Router("/t1",&controllers.TestController{},"get:ShowIndex")*/
	beego.Router("/reg", &controllers.RegController{})
	beego.Router("/",  &controllers.RegController{})
    beego.Router("/login",&controllers.LoginController{},"get:ShowLogin;post:HandleLogin")
	}
