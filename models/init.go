package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//register all tables
	orm.RegisterModel(new(Userinfo), new(Admininfo), new(Deviceinfo), new(Historyinfo), new(Command), new(OperationRecord), new(Alarm), new(Script))
	//register mysql driver
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//register default database
	orm.RegisterDataBase("default", "mysql", "root:autelan@/lte_test?charset=utf8&&loc=Asia%2FShanghai")

	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
}

func CheckDatabase() bool {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		return false
	} else {
		return true
	}
}
