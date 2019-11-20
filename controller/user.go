package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shopping/config"
	"shopping/model"
	"shopping/redis"
	"shopping/util/tools"
	"time"
)

func LoginHandler(c *gin.Context) {
	phone := c.Query("phone")
	pwd := c.Query("password")

	// 1. 判断用户是否存在
	u := &model.User{Phone: phone, Password: pwd}
	if u.FoundUser() { // 用户名存在
		if u.CheckLogin() { //验证登录密码
			tools.JsonMessage(c, 20, "登录成功", u)
			tokenStr, err := model.CreateToken([]byte(model.SecretKey), "", phone, "MAC")
			if err != nil {
				tools.JsonMessage(c, 21, "生成token成功", tokenStr)

			} else {
				token := &model.User{
					Phone: phone,
					Token: tokenStr,
				}
				if token.EditUser() {
					tools.JsonMessage(c, 20, "修改数据库的token成功", tokenStr)
					cmd := "HSET user " + phone + " " + tokenStr
					reply, e := redis.DoRedisCommand(cmd)
					fmt.Println(reply, e)
					if e != nil {
						tools.JsonMessage(c, 21, "将token添加入缓存失败", reply)
						return
					} else {
						tools.JsonMessage(c, 20, "8888将token添加入缓存成功", token)

					}
				}
			}
			return
		} else {
			tools.JsonMessage(c, 21, "fail", "登录失败，密码错误")

		}
	} else {
		tools.JsonMessage(c, 22, "fail", "登陆失败，用户名不存在")
	}
}

// 注册
func RegisterHandler(c *gin.Context) {

	user := &model.User{
		Phone:       c.PostForm("phone"),
		Password:    c.PostForm("password"),
		Name:        c.PostForm("name"),
		Sex:         c.PostForm("sex"),
		Role:        0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	}
	if !tools.CheckPhone(user.Phone) {
		tools.JsonMessage(c, 10, "fail", "用户名格式错误")
		return
	}
	if !tools.CheckPassword(user.Password) {
		tools.JsonMessage(c, 10, "fail", "密码格式错误")
		return
	}
	if user.FoundUser() { // 用户名存在
		tools.JsonMessage(c, 22, "fail", "注册失败,用户名已存在")
		return
	} else {
		if user.AddUser() {
			tools.JsonMessage(c, 20, "success", "注册成功")
		} else {
			tools.JsonMessage(c, 21, "fail", "注册失败，添加用户不成功")
		}
	}
}

// 修改用户密码
func ModifyUser(c *gin.Context) {
	user := &model.User{
		Phone:       c.PostForm("phone"),
		Password:    c.PostForm("password"),
		Update_time: time.Now(),
	}
	if !tools.CheckPassword(user.Password) {
		tools.JsonMessage(c, 10, "fail", "密码格式错误")
		return
	}
	if user.EditUser() {
		tools.JsonMessage(c, 20, "success", "修改成功")
	} else {
		tools.JsonMessage(c, 21, "fail", "修改失败")
	}
}

// 修改用户信息
func ModifyUserInfo(c *gin.Context) {

	user := &model.User{
		Phone:       c.PostForm("phone"),
		Name:        c.PostForm("name"),
		Sex:         c.PostForm("sex"),
		Update_time: time.Now(),
	}
	if user.EditUser() {
		tools.JsonMessage(c, 20, "success", "修改成功")
	} else {
		tools.JsonMessage(c, 21, "fail", "修改失败")
	}
}

//头像上传
func UploadImage(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	phone := c.PostForm("phone")
	filename := phone + ".png"
	path := config.GetLocalPath() + filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		tools.JsonMessage(c, 21, "fail", "上传图片失败")
		return
	} else {
		user := &model.User{
			Phone:   phone,
			Picture: path,
		}
		if user.EditUser() {
			tools.JsonMessage(c, 22, "success", path)
			return

		} else {
			tools.JsonMessage(c, 20, "fail", "图片路径保存失败")

		}

	}
}

//显示用户信息
func ShowUserInfo(c *gin.Context) {
	phone := c.Query("phone")
	user := &model.User{
		Phone: phone,
	}
	userInfo := user.GetInfo()
	tools.JsonMessage(c, 20, "success", userInfo)

}
