package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shopping/model"
	"shopping/util/tools"
	"strconv"
	"time"
)

//1。发行积分
func PushPoint(c *gin.Context) {
	// post 发行商ID PointMakerId和excel表具体账号Account和发型金额Balance
	/*excelFile, err := c.FormFile("excel")
	if err != nil {
		fmt.Println(err)
		return
	}*/
	account := c.PostForm("phone")
	balance, err := strconv.Atoi(c.PostForm("balance"))
	PointMakerId, err := strconv.Atoi(c.PostForm("PointMakerId"))
	//有效时长( '24h'/'30m'/'30s')
	timelong := c.PostForm("timelong")
	if err != nil {
		fmt.Println(err)
		return
	}
	now := time.Now()
	//24个小时后 timelong=24h
	hh, _ := time.ParseDuration(timelong)
	expiretime := now.Add(hh)
	Publish := &model.Publish{
		PointMakerId: PointMakerId,
		Account:      account,
		Balance:      balance,
		Publish_time: now,
		Expire_time:  expiretime,
	}
	if Publish.PushPoint() {
		tools.JsonMessage(c, 20, "发型积分成功", Publish)
	} else {
		tools.JsonMessage(c, 21, "发行积分失败", Publish)
	}

}

//2。显示用户未领取的积分 参数 account == phone
func ShowNoPoint(c *gin.Context) {
	var account = c.Query("phone")
	publish := &model.Publish{
		Account: account,
		Pick:    0, //1:已领取
	}
	flag, pu := publish.ShowNoPoint()
	if flag {
		tools.JsonMessage(c, 20, "显示未领取积分成功", pu)
	} else {
		tools.JsonMessage(c, 21, "显示未领取积分失败", pu)
	}
}

//3.显示用户积分 参数：userid  GetUserPoint
func GetUserPoint(c *gin.Context) {
	var userid, err = strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	userPoint := &model.UserPoint{
		UserId: userid,
	}
	if userPoint.ShowUserPoint() {
		tools.JsonMessage(c, 20, "显示用户积分成功", userPoint)
	} else {
		tools.JsonMessage(c, 21, "显示用户积分失败", userPoint)
	}
}

//4.用户领取积分，参数publish 的ID 和用户id
func GetPoint(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	publish := &model.Publish{
		ID:   id,
		Pick: 0,
	}

	if publish.GetPoint() {
		userPoint := &model.UserPoint{
			PointId:     publish.ID,
			UserId:      userid,
			Balance:     publish.Balance,
			Total:       publish.Balance, //初始未分发的积分
			Expire_time: publish.Expire_time,
			Create_time: time.Now(),
			Update_time: time.Now(),
		}
		if userPoint.CreateUserPoint() {
			tools.JsonMessage(c, 20, "领取积分成功", userPoint)
		}
	} else {
		tools.JsonMessage(c, 21, "领取积分失败", publish)
	}
}
