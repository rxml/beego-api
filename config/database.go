package config

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitMysql(){
	driverName := beego.AppConfig.String("driverName")
	dbuser := beego.AppConfig.String("mysqlUser")
	dbpwd := beego.AppConfig.String("mysqlpwd")
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbname := beego.AppConfig.String("dbname")
	/*
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		db = db1
		fmt.Println("链接成功")
	}
	*/
	orm.Debug= true
	orm.RegisterDriver(driverName, orm.DRMySQL)
	orm.RegisterDataBase("default", driverName, dbuser + ":"+dbpwd+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8")
	o := orm.NewOrm()
    o.Using("default") // 默认使用 default，你可以指定为其他数据库
	
}