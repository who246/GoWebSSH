package main

import (
	_ "github.com/who246/GoWebSSH/routers"
	"github.com/astaxie/beego"
	"github.com/who246/GoWebSSH/models"
	_ "github.com/who246/GoWebSSH/filter"
)

func main() {
	models.CreateTable()
	beego.Run()
}

