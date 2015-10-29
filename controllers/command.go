package controllers
import (
	"github.com/astaxie/beego"
	"github.com/who246/GoWebSSH/models"
	"github.com/astaxie/beego/orm" 
	"github.com/astaxie/beego/utils/pagination"
)
type CommandController struct {
	baseController
}
 func (this *CommandController) AddShow(){
	m := models.Command{}
	this.Data["m"] = m;
	this.Data["op"] = "a";
	this.TplNames="command/form.html"
 }
func (this *CommandController) Add() {
	cmd := this.GetString("cmd")  
	name := this.GetString("name")  
	command := models.Command{Cmd:cmd,Username:this.Getusername(),Name:name}
	_,err := orm.NewOrm().Insert(&command)
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.JsonSuccess();
}
func (this *CommandController) List(){
	page,err := this.GetInt("pageNum",1);
	numPerPage,err := this.GetInt("numPerPage",20);
	keyword := this.GetString("keyword","") 
	if err != nil {
	beego.Error(err);
	}
	extend := map[string]string{}; 
    extend["username"] = this.Getusername()
	maps,err := models.GetCommandPageList(numPerPage,(page-1)*numPerPage,extend,keyword)
	count, err := models.GetCommandCount(extend,keyword);
	beego.Info(count);
	if(err != nil){
		beego.Error(err);
	}
	this.Ctx.Request.Form.Add("p",this.GetString("pageNum"))
	p := pagination.NewPaginator(this.Ctx.Request, numPerPage, count)
	this.Data["paginator"] = p
	this.Data["commands"] = maps
	this.Data["keyword"] = keyword
	this.TplNames="command/list.html"
}
func (this *CommandController) Del() {
	id,err := this.GetInt("id",-1);
	 
	 _,err = orm.NewOrm().Delete(&models.Command{Id: id});
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.JsonSuccess();
}
func (this *CommandController) ModifyShow() {
	id,err := this.GetInt("id",-1)
	m := models.Command{Id: id}
	err = orm.NewOrm().Read(&m,"id")
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.Data["m"] = m;
	this.Data["op"] = "m";
	this.TplNames="command/form.html"
}
func (this *CommandController) Modify() {
	id,err := this.GetInt("id",-1)
	name := this.GetString("name","") 
	cmd := this.GetString("cmd","") 
	m := models.Command{Id: id,Name:name,Cmd:cmd}
	_,err = orm.NewOrm().Update(&m,"Name","Cmd","ModifyTime");
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
     this.JsonSuccess();
}
func (this *CommandController) CmdFilterList() {
	id,_ := this.GetInt("id",-1);
	m := models.Server{Id: id}
	err := orm.NewOrm().Read(&m,"id")
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.List(); 
	this.Data["server"]=m;
	this.TplNames="command/table_list.html"
}