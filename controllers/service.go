package controllers

import (
	"github.com/astaxie/beego"
	"github.com/who246/GoWebSSH/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	
)

type ServiceController struct {
	baseController
}
 func (this *ServiceController) AddShow(){
	m := models.Server{}
	this.Data["m"] = m;
	this.Data["op"] = "a";
	this.TplNames="service/form.html"
 }
func (this *ServiceController) List(){
	page,err := this.GetInt("pageNum",1);
	numPerPage,err := this.GetInt("numPerPage",20);
	keyword := this.GetString("keyword","")
	
	if err != nil {
	beego.Error(err);
	}
	extend := map[string]string{}; 
	extend["username"] = this.Getusername()
	maps,err := models.GetServerPageList(numPerPage,(page-1)*numPerPage,extend,keyword)
	count, err := models.GetServerCount(extend,"");
	beego.Info(count);
	if(err != nil){
		beego.Error(err);
	}
	this.Ctx.Request.Form.Add("p",this.GetString("pageNum"))
	p := pagination.NewPaginator(this.Ctx.Request, numPerPage, count)
	this.Data["paginator"] = p
	this.Data["servers"] = maps
	this.Data["keyword"] = keyword
	this.TplNames="service/list.html"
}
func (this *ServiceController) Trim() {
	this.Data["host"] = this.Ctx.Request.Host
	this.Data["id"]= this.GetString("id")
	this.Data["cmdId"]= this.GetString("cmdId") 
	this.TplNames = "service/term2.html"
}

func (this *ServiceController) Add() {
	serverName := this.GetString("serverName")
	ip := this.GetString("ip")
	loginName := this.GetString("loginName")
	loginPassword := this.GetString("loginPassword")
	server := models.Server{Ip:ip,ServerName:serverName,LoginName:loginName,LoginPassword:loginPassword,Username:this.Getusername()}
	_,err := orm.NewOrm().Insert(&server)
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.JsonSuccess();
}
func (this *ServiceController) Del() {
	id,err := this.GetInt("id",-1);
	 
	 _,err = orm.NewOrm().Delete(&models.Server{Id: id});
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.JsonSuccess();
}
func (this *ServiceController) ModifyShow() {
	id,err := this.GetInt("id",-1)
	m := models.Server{Id: id}
	err = orm.NewOrm().Read(&m,"id")
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.Data["m"] = m;
	this.Data["op"] = "m";
	this.TplNames="service/form.html"
}
func (this *ServiceController) Modify() {
	id,err := this.GetInt("id",-1)
	serverName := this.GetString("serverName","")
	ip := this.GetString("ip","")
	loginName := this.GetString("loginName","")
	loginPassword := this.GetString("loginPassword","")
	m := models.Server{Id: id,ServerName:serverName,Ip:ip,LoginName:loginName,LoginPassword:loginPassword}
	_,err = orm.NewOrm().Update(&m,"Ip","ServerName","LoginName","LoginPassword","ModifyTime");
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
     this.JsonSuccess();
}

