package main

import (
	_ "apigo/routers"
	"github.com/astaxie/beego"
	"fmt"
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
	/*if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}*/
	//递增
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a,b,c,d)
	var f = 1
	var h = &f
	fmt.Println(h,*h)
	var aa [10]int
	var i int
	for i = 0; i < 10; i++ {
		aa[i] = i + 100
	}
	fmt.Println(aa)
	cc := aa[:3] 
	fmt.Println(cc)

	//切片
	var j,k int
	qiePian := make([][]int, 3)
	for j = 0; j < 3; j++ {
		len := j+1
		qiePian[j] = make([]int, len)
		for k = 0; k < len; k++ {
			qiePian[j][k] = j+k
		}
	}
	fmt.Println("2d: ", qiePian)
	
	var p Phone
	p = new (NokiaPhone)
	p.call()

	p = new (IPhone)
	p.call()
	config.InitMysql()
	beego.Run()
}
