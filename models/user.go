package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       int `json:"id" orm:"pk;column(id)"`
	Username string `json:"username" orm:"column(username)"`
	Realname string `json:"realname" orm:"column(realname)"`
	Password string `json:"password" orm:"column(password)"`
	Number int `json:"number" orm:"column(number)"`
	Serialnum string `json:"serialnum" orm:"column(serialnum)"`
	Age int `json:"age" orm:"column(age)"`
	Ethnic string `json:"ethnic" orm:"column(ethnic)"`
	Birthday string `json:"birthday" orm:"column(birthday)"`
	LastLoginTime string `json:"last_login_time" orm:"column(last_login_time)"`
	LastLoginIp string `json:"last_login_ip" orm:"column(last_login_ip)"`
	UpdatedAt string `json:"updated_at" orm:"column(updated_at)"`
}

type AllUsers struct {
	Id int
	Username string
	Realname string
	Age int
}


func init() {
	orm.RegisterModelWithPrefix("ml_", new(Users))
}

/*
func AddUser(u User) string {
	
}
*/
func GetUserByName(username string) (Users, error) {
	o := orm.NewOrm()
	user := Users{Username: username}
    err := o.Read(&user, "username")
    //err := db.Filter("username", username).One(&user)
    if err == nil {
        return user,nil
    }
	return user, err
	
}

//更新
func UpdateById(userInfo *Users) error {
	o := orm.NewOrm()
	user := Users{Id: userInfo.Id}
	if o.Read(&user) == nil {
		if userInfo.LastLoginTime != "" {
			//最近登录时间
			user.LastLoginTime = userInfo.LastLoginTime
		}
		if userInfo.LastLoginIp != "" {
			//最近登录ip
			user.LastLoginIp = userInfo.LastLoginIp
		}
		if userInfo.Password != "" {
			//更改密码
			user.Password = userInfo.Password
		}
		user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(user)
		//return 0,nil
		_, err := o.Update(&user);
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllUsers(keywords string,condition *Users) (int64,[]*Users, error){
	o := orm.NewOrm()
	var user []*Users
	cond := orm.NewCondition()
	query := o.QueryTable("ml_users")
	if keywords != "" {
		cond = cond.And("username__contains", keywords).Or("realname__contains", keywords)
	}

	num, err := query.SetCond(cond).All(&user)
	if err == nil {
		return num, user,nil
	}
	return 0,nil,err
}
/*
func UpdateUser(uid string, uu *User) (a *User, err error) {
	
}

func Login(username, password string) bool {
	
}

func DeleteUser(uid string) {
	
}
*/