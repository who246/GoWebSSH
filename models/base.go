package models



import ( 
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" 
)


func init() {
	
 
	orm.Debug = true
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "./gowebssh.db")
    
}
func CreateTable() {
    name := "default"                          //数据库别名
    force := false                             //不强制建数据库
    verbose := true                            //打印建表过程
    err := orm.RunSyncdb(name, force, verbose) //建表
    if err != nil {
        beego.Error(err)
    }
}
