package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"apigo/models"
	"path/filepath"
	"os"
	"time"
)
// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
/*func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	fmt.Println("11111")
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	fmt.Println(username)
	fmt.Println(password)
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
*/


func (this *UserController) Users() {
	keywords := this.GetString("keywords")
	condition := models.Users{}
	num, rows, err := models.GetAllUsers(keywords,&condition)
	if err != nil {
		errorMessage := map[string]interface{}{"code": 0, "message": "内部错误"}
		this.Data["json"] = &errorMessage
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{"total_count": num, "userList": rows}
	response :=  map[string]interface{}{"code": 1, "message": "获取成功", "data": data}
	this.Data["json"] = &response
	this.ServeJSON()
}


/*
* @author molin
* @date 2019-09-21
* @上传头像
*/
func (this *UserController) Avatar() {
	f, h, err := this.GetFile("file")
	defer f.Close()
	var code int64
	var message string
	fmt.Println(h)
    if err == nil {
		filesuffix := filepath.Ext(h.Filename)
		fmt.Println(filesuffix)
		if filesuffix == ".jpg" || filesuffix == ".png" || filesuffix == "jpeg" || filesuffix == ".gif" {
			if fileSizer, ok := f.(Size); ok {
				fileSize := fileSizer.Size()
				// fmt.Printf("上传%v文件的大小为: %v", fileSize, h.Filename)
				if fileSize > int64(Filebytes) {
					code = 0
					message = "获取上传文件错误:文件大小超出33M"
					goto END
				}
			} else {
				code = 0
				message = "获取上传文件错误:无法读取文件大小"
				goto END
			}
			
			dirpath := "static/upload/avatar"
			_, err := os.Stat(dirpath)
			if err != nil {
				//文件夹不存在  则创建文件夹
				err = os.MkdirAll(dirpath, os.ModePerm)
				if err != nil {
					code  = 0
					message = "创建文件夹失败"
					goto END
                }
			}

			//新建文件名
			timeStamp := time.Now().Unix()
			fileName := fmt.Sprintf("%d-%s", timeStamp, h.Filename)
			
			err = this.SaveToFile("file", filepath.Join(dirpath, fileName)) // 保存位置在 static/upload, 没有文件夹要先创建
			if err == nil {
				code = 1
				message = "上传成功"
			} else {
				code = 0
				message = "上传失败"
			}
		} else {
			code = 0
			message = "不支持上传该文件类型"
		}
		
	} else {
		code = 0
		message = "没有发现上传文件"
	}
	END:
	response :=  map[string]interface{}{"code": code, "message": message}
	this.Data["json"] = &response
	this.ServeJSON()
}
