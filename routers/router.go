package routers

import (
	"Gemini/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/about",&controllers.AboutController{})
}
