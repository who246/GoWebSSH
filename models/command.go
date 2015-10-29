package models
import (
	"github.com/astaxie/beego/orm"
	"time"
  
)
type Command struct {
	Id int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	ModifyTime time.Time  `orm:"auto_now;type(datetime)"`
	Cmd string `orm:"type(text);null"` 
	Username string `orm:"size(60)"`
	Name string `orm:"size(60)"` 
}
func init(){
	orm.RegisterModel(new (Command))
}
func   GetCommandPageList(pageSize,pageStart int,extend map[string] string,title string) ([]orm.Params, error)  {
	o := orm.NewOrm()
    qs := o.QueryTable("command")
     for k, v:= range extend{
     qs= qs.Filter(k,v);
    }
	if title != "" {
		qs=qs.Filter("name__contains",title);
	}
	var params []orm.Params
	if _,err := qs.OrderBy( "-modify_time").Limit(pageSize,pageStart).Values(&params); err !=nil {
		return nil,err
	}
    
	return params,nil
}

func GetCommandCount(extend map[string] string,title string) (count int64,err error){
	o := orm.NewOrm()
	qs := o.QueryTable("command")
     for k, v:= range extend{
     qs= qs.Filter(k,v);
    }
	if title != "" {
		qs=qs.Filter("name__contains",title);
	}
	count, err = qs.Count()
	return count,err;
}
 