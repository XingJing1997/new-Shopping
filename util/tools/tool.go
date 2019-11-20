package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"regexp"
	"shopping/model"
	"shopping/redis"
	"strconv"
	"time"
)

// 1. 跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorize")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
	}
}

func JsonMessage(c *gin.Context, code int, message string, date interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
		"data": date,
	})
}

//检查手机号
func CheckPhone(phone string) bool {
	phoneReg := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(phoneReg)
	return reg.MatchString(phone)
}
func CheckPassword(pwd string) bool {
	pwdReg := `^[a-z0-9_-]{6,18}$`
	reg := regexp.MustCompile(pwdReg)
	return reg.MatchString(pwd)
}

func NowTime() string {
	time := time.Now().Format("2006-01-02 15:04:03")
	return time
}

// 1. 生成唯一的订单号
func GetOrderNo(userid int) (orderno string) {
	rand.Seed(time.Now().UnixNano())
	rand1 := rand.Intn(8999) + 1000
	time := time.Now().Format("2006-01-02@15:04:03")
	//随机产生字符
	bytes := make([]byte, 3)
	for i := 0; i < 3; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	orderno = strconv.Itoa(userid) + string(bytes[0]) + strconv.Itoa(rand1) + string(bytes[1]) + time + string(bytes)
	return orderno
}
func ValidateToken() gin.HandlerFunc {
	fmt.Println("aaa")
	return func(c *gin.Context) {
		//1.检查token参数
		token := c.Request.FormValue("token")
		if token == "" {
			JsonMessage(c, 0, "缺少参数token", "")
			c.Abort()
			return
		}
		//2 :解析token参数，第一步对比是否过时，第二步对应与数据库的是 否一致，都通过了了就可以继续了
		claims, err := model.ParseToken(token, []byte(model.SecretKey)) //fmt.Println(claims)
		if nil != err {
			JsonMessage(c, 0, "token异常，重新登陆", "")
			c.Abort()
			return
		}
		userID := claims.(jwt.MapClaims)["uid"].(string)

		fmt.Println(userID)
		cmd := "HGET user " + userID //到redis缓存中查找key 对应的 token
		reply, _ := redis.DoRedisCommand(cmd)
		if reply == "" { //redis缓存中没有，则查找mysql获取token并添加到redis中
			user := &model.User{
				Phone: userID,
			}
			tokenInfo := user.GetInfo() //查找mysql
			cmd := "HSET user " + userID + " " + tokenInfo.Token
			reply, e := redis.DoRedisCommand(cmd)
			fmt.Println(reply, e)
			if e != nil {
				JsonMessage(c, 21, "将token添加入缓存失败", reply)
				return
			} else {
				JsonMessage(c, 20, "将token添加入缓存成功", reply)

			}
		}
		c.Next()
	}
}
