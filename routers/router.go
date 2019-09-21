// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"apigo/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"

)

var authFilter = func (ctx *context.Context) {

	//登录才可以访问
	response := make(map[string]interface{})
	
	authString := ctx.Input.Header("Authorization")
	if authString == "" || len(authString) == 0{
		response["code"] = 0
		response["message"] = "请传入授权令牌"
		ctx.Output.JSON(response, false, true)
		return
	}
	bearer := authString[:6] 
	tokenString := authString[7:]
	fmt.Println(bearer,tokenString)
    if bearer != "Bearer" {
        response["code"] = 0
		response["message"] = "令牌无效,请重新登录"
		ctx.Output.JSON(response, false, true)
		return
    }
    // Parse token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("mykey"), nil
    })
    if err != nil {
        beego.Error("Parse token:", err)
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                response["code"] = 0
				response["message"] = "令牌无效,请重新登录"
				ctx.Output.JSON(response, false, true)
				return
            } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
                response["code"] = 0
				response["message"] = "令牌无效,请重新登录"
				ctx.Output.JSON(response, false, true)
				return
            } else {
                response["code"] = 0
				response["message"] = "令牌无效,请重新登录"
				ctx.Output.JSON(response, false, true)
				return
            }
        } else {
            response["code"] = 0
			response["message"] = "令牌无效,请重新登录"
			ctx.Output.JSON(response, false, true)
			return
        }
    }
    if !token.Valid {
        response["code"] = 0
		response["message"] = "令牌无效,请重新登录"
		ctx.Output.JSON(response, false, true)
		return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        response["code"] = 0
		response["message"] = "令牌无效,请重新登录"
        ctx.Output.JSON(response, false, true)
        return
    }
	fmt.Println("令牌通过", claims["username"].(string))
}

func init() {

	public := 
		beego.NewNamespace("/api",
			beego.NSRouter("/", &controllers.PublicController{}, "*:Index"),
			beego.NSRouter("/login", &controllers.PublicController{}, "post:Login"),
			
		)

	auth := 
		beego.NewNamespace("/api/auth",
			beego.NSBefore(authFilter),
			beego.NSRouter("/users", &controllers.UserController{}, "*:Users"),
			beego.NSRouter("/avatar", &controllers.UserController{}, "post:Avatar"),
		)
	
	beego.AddNamespace(public,auth)
	
}
