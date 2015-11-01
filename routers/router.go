package routers

import (
	"github.com/who246/GoWebSSH/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Index")
	 beego.Router("/error", &controllers.MainController{},"*:Error")
	beego.Router("/admin/index", &controllers.MainController{},"*:Admin")
	beego.Router("/admin/ws", &controllers.WSController{})
	beego.Router("/admin/service/list", &controllers.ServiceController{},"*:List")
	beego.Router("/admin/service/trim", &controllers.ServiceController{},"*:Trim") 
	beego.Router("/admin/service/add", &controllers.ServiceController{},"get:AddShow")
	beego.Router("/admin/service/add", &controllers.ServiceController{},"post:Add")
	beego.Router("/admin/service/del", &controllers.ServiceController{},"post:Del")
	beego.Router("/admin/service/mdy", &controllers.ServiceController{},"get:ModifyShow")
	beego.Router("/admin/service/mdy", &controllers.ServiceController{},"post:Modify")
	beego.Router("/admin/cmd/add", &controllers.CommandController{},"get:AddShow")
	beego.Router("/admin/cmd/add", &controllers.CommandController{},"post:Add")
	beego.Router("/admin/cmd/list", &controllers.CommandController{},"*:List")
	beego.Router("/admin/cmd/del", &controllers.CommandController{},"post:Del")
	beego.Router("/admin/cmd/mdy", &controllers.CommandController{},"get:ModifyShow")
	beego.Router("/admin/cmd/mdy", &controllers.CommandController{},"post:Modify")
	beego.Router("/admin/cmd/filterlist", &controllers.CommandController{},"get:CmdFilterList")
	
	beego.Router("/admin/user/list", &controllers.UserController{},"*:List") 
	beego.Router("/admin/user/add", &controllers.UserController{},"get:AddShow")
	beego.Router("/admin/user/add", &controllers.UserController{},"post:Add")
	beego.Router("/admin/user/del", &controllers.UserController{},"post:Del")
	beego.Router("/admin/user/mdy", &controllers.UserController{},"get:ModifyShow")
	beego.Router("/admin/user/mdy", &controllers.UserController{},"post:Modify")
	beego.Router("/login", &controllers.UserController{},"post:Login")
	beego.Router("/admin/logout", &controllers.UserController{},"*:Logout")
	beego.Router("/admin/showchangepwd", &controllers.UserController{},"get:ShowChangePwd")
	beego.Router("/admin/changepwd", &controllers.UserController{},"post:ChangePwd")
}
