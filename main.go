package main

import (
	_ "apigo/routers"
	"github.com/astaxie/beego"
	"apigo/config"
)

type Phone interface {
    call()
}

type NokiaPhone struct {
	id int
	name string
}

func (nokiaPhone NokiaPhone) call() {
	nokiaPhone.id = 1
	nokiaPhone.name = "zhangsan"
    fmt.Println(nokiaPhone)
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
	//链接数据库
	config.InitMysql()
	//启动
	beego.Run()
}
