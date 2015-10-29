package controllers

import (
	"github.com/astaxie/beego"
	)
	
type baseController struct {
	beego.Controller
}
func (this *baseController) JsonSuccess() {
	this.ToJsonSuccess("操作成功")
}
func (this *baseController) ToJsonSuccess( msg string) {
	this.Data["json"] = map[string]string{
      "statusCode":"200",
      "message":msg,
      "navTabId":"",
      "rel":"",
      "callbackType":"",
      "forwardUrl":"",
}
	this.ServeJson();
}
func (this *baseController) JsonFail() {
	this.ToJsonFail("操作失败")
}
func (this *baseController) ToJsonFail(msg string) {
	this.Data["json"] = map[string]string{
      "statusCode":"300",
      "message":msg,
}
	this.ServeJson();
}

func (this *baseController) Getusername() (username string)  {
	return this.GetSession("username").(string);
}