package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

var LoginUrlManager = func(ctx *context.Context) {
     username, ok := ctx.Input.Session("username").(string) 
	if !ok {
		ctx.Redirect(302, "/")
		return;
	}
	ctx.Input.SetData("username",username);
	if strings.Index(ctx.Request.RequestURI,"/admin/user")!=-1 && username != "admin"{
		ctx.WriteString("错误：权限不足❌");
		return;
	}
}

func init() {
  beego.InsertFilter("/admin/*",beego.BeforeRouter,LoginUrlManager)
}

