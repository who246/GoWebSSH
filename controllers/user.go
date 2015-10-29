package controllers
import (
	"github.com/astaxie/beego"
	"github.com/who246/GoWebSSH/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	
)
type UserController struct {
	baseController
}
func (this *UserController) AddShow(){
	m := models.User{}
	this.Data["m"] = m;
	this.Data["op"] = "a";
	this.TplNames="user/form.html"
 }
func (this *UserController) List(){
	page,err := this.GetInt("pageNum",1);
	numPerPage,err := this.GetInt("numPerPage",20);
	keyword := this.GetString("keyword","")
	
	if err != nil {
	beego.Error(err);
	}
	extend := map[string]string{}; 
	maps,err := models.GetUserPageList(numPerPage,(page-1)*numPerPage,extend,keyword)
	count, err := models.GetUserCount(extend,""); 
	if(err != nil){
		beego.Error(err);
	}
	this.Ctx.Request.Form.Add("p",this.GetString("pageNum"))
	p := pagination.NewPaginator(this.Ctx.Request, numPerPage, count)
	this.Data["paginator"] = p
	this.Data["users"] = maps
	this.Data["keyword"] = keyword
	this.TplNames="user/list.html"
}
func (this *UserController) Add() {
	username := this.GetString("username") 
	password := this.GetString("password") 
	user := models.User{Username:username}
	exit := orm.NewOrm().QueryTable("user").Filter("username",username).Exist();
	if exit { 
		this.ToJsonFail("已经存在！");
		return;
	}
	user.Password = password;
	_,err := orm.NewOrm().Insert(&user)
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.JsonSuccess();
}
func (this *UserController) Del() {
	id,err := this.GetInt("id",-1);
	 
	 _,err = orm.NewOrm().Delete(&models.User{Id: id});
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.JsonSuccess();
}
func (this *UserController) ModifyShow() {
	id,err := this.GetInt("id",-1)
	m := models.User{Id: id}
	err = orm.NewOrm().Read(&m,"id")
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	this.Data["m"] = m;
	this.Data["op"] = "m";
	this.TplNames="user/form.html"
}
func (this *UserController) Modify() {
	id,err := this.GetInt("id",-1)
	username := this.GetString("username","") 
	password := this.GetString("password","") 
	
	m := models.User{Id:id,Password:password,Username:username}
	_,err = orm.NewOrm().Update(&m,"Password");
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
     this.JsonSuccess();
}

func (this *UserController) Login() {
   username := this.GetString("username","");
   password := this.GetString("password","");
   isRemeber:=this.GetString("isRemeber",""); 
   m := models.User{}
   err := orm.NewOrm().QueryTable("user").Filter("username",username).One(&m)

   if err!=nil || m.Password != password {
	this.Data["error"] = "用户名或密码错误";
	this.TplNames="login.html"
	return;
   }
    
   this.SetSession("username",username)
   this.SetSession("uid",m.Id)
   if isRemeber == "on" {
   this.SetSecureCookie("gowebssh","username",username)
   this.SetSecureCookie("gowebssh","password",password)
   } 
	this.Redirect("/admin/index",302)
}
func (this *UserController) Logout() {
    this.DelSession("username");
	 this.DelSession("uid");
	this.Redirect("/",302)
}
func (this *UserController) ShowChangePwd() {
	this.TplNames="user/cgpwd.html"
}
func (this *UserController) ChangePwd() {
	newpassword1 := this.GetString("newpassword1","") 
	newpassword2 := this.GetString("newpassword2","") 
	password := this.GetString("password","") 
	if newpassword1 == "" {
		this.ToJsonFail("密码不能为空！！");
		return;
	}
	if(newpassword1 != newpassword2){ 
		this.ToJsonFail("两次密码不一样！！");
		return;
	}
	id := this.GetSession("uid").(int);
	m := models.User{Id: id};
	 err := orm.NewOrm().Read(&m); 
	if err !=nil{
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	if password!=m.Password{
		this.ToJsonFail("密码错误！！");
		return;
	}
	m.Password = newpassword1
	_,err = orm.NewOrm().Update(&m,"Password");
	println(m.Password)
	if err!=nil {
		beego.Error(err)
		this.ToJsonFail(err.Error());
		return;
	}
	 this.DelSession("username");
	 this.DelSession("uid");
     this.JsonSuccess();
}