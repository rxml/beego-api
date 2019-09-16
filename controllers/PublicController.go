package controllers

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"apigo/models"

)

type PublicController struct {
	beego.Controller
}

//密码验证
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
//密码验证
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func (p *PublicController) Index(){
	fmt.Println("hello world!");
	mystruct := map[string]interface{}{"code": 1, "message": "hello world!!!"} 
	p.Data["json"] = &mystruct
	p.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (this *PublicController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	userInfo, err := models.GetUserByName(username)


	if err != nil{
		errorMessage1 := map[string]interface{}{"code": 0, "message": "用户不存在"}
		this.Data["json"] = &errorMessage1
		this.ServeJSON()
		return
	}
	//fmt.Println(userInfo.Password)
	//hash, _ := HashPassword(password)
	//fmt.Println("Password:", password)
	//fmt.Println("Hash:    ", userInfo.Password)
	match := CheckPasswordHash(password, userInfo.Password)

	fmt.Println(match)
	if !match {
		errorMessage2 := map[string]interface{}{"code": 0, "message": "密码错误,请重新输入"}
		this.Data["json"] = &errorMessage2
		this.ServeJSON()
		return
	}
	//登录成功把jwt字符串发给客户端，客户端需要保存起来比如localStorage中，访问别的API时加到header里面, 注意这个里面exp是必须要的，生命周期。其他都是根据自己需求添加的键值对
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	if username == "admin" {
		claims["admin"] = "true"
	} else {
		claims["admin"] = "false"
	}
	claims["exp"] = time.Now().Add(time.Hour * 480).Unix() //20天有效期，过期需要重新登录获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("mykey"))
	if err != nil {
		beego.Error("jwt.SignedString:", err)
		return
	}
	
	user := map[string]interface{}{"username": userInfo.Username,"realname": userInfo.Realname,"age": userInfo.Age,"last_login_ip": userInfo.LastLoginIp,"last_login_time": userInfo.LastLoginTime}
	data := map[string]interface{}{"token_type": "Bearer", "access_token": tokenString, "userInfo": user}
	mystruct :=  map[string]interface{}{"code": 1, "message": "登录成功", "data": data}
	this.Data["json"] = &mystruct
	// userInfo := map[string]interface{}{"username": username, "password": password}
	// mystruct := map[string]interface{}{"code": 1, "message": "获取成功", "data": map[string]interface{}{"userInfo": &userInfo}}
	// p.Data["json"] = &mystruct
	//更新时间
	updateData := models.Users{Id: userInfo.Id, LastLoginTime: time.Now().Format("2006-01-02 15:04:05"), LastLoginIp: this.Ctx.Input.IP()}
	uerr := models.UpdateById(&updateData)
	if uerr != nil {
		beego.Error("更新出错")
		return
	}
	
	this.ServeJSON()
}
