package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	 username,ok := this.GetSecureCookie("gowebssh","username");
   if ok {
	this.Data["username"] = username;
	   beego.Info(username);
   }
   password,ok := this.GetSecureCookie("gowebssh","password");
   if ok {
	this.Data["password"] = password;
	 beego.Info(password);
   }
	this.Data["host"] = this.Ctx.Request.Host
	this.TplNames = "login.html"
}
func (this *MainController) Admin() {
	this.Data["host"] = this.Ctx.Request.Host
	this.TplNames = "index.html"
}
func (this *MainController) Error() { 
    this.Data["error"] = this.GetString("error"); 
	this.TplNames = "error.html"
}