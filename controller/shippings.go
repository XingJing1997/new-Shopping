package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shopping/model"
	"shopping/util/tools"
	"strconv"
	"time"
)

// 添加地址
func AddShipping(c *gin.Context) {
	userid, err := strconv.Atoi(c.PostForm("userid"))
	if err != nil {
		fmt.Println("add:", err)
		return
	}
	mobile := c.PostForm("mobile")
	if len(mobile) != 0 {
		if !tools.CheckPhone(mobile) {
			tools.JsonMessage(c, 10, "fail", "手机号格式错误")
			return
		}
	}
	local := &model.Shippings{
		UserId:           userid,
		ReceiverName:     c.PostForm("name"),
		ReceiverMobile:   mobile,
		ReceiverProvince: c.PostForm("province"),
		ReceiverCity:     c.PostForm("city"),
		ReceiverDistrict: c.PostForm("district"),
		ReceiverAddress:  c.PostForm("address"),
		ReceiverZip:      c.PostForm("zip"),
		Create_time:      time.Now(),
		Update_time:      time.Now(),
	}
	if local.AddShipping() {
		tools.JsonMessage(c, 20, "success", local)
	} else {
		tools.JsonMessage(c, 20, "fail", "添加收货地址失败")

	}
}

//根据用户ID修改地址
func UpdateShippings(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		fmt.Println(err)
		return
	}
	mobile := c.PostForm("mobile")
	if len(mobile) != 0 {
		if !tools.CheckPhone(mobile) {
			tools.JsonMessage(c, 10, "fail", "手机号格式错误")
			return
		}
	}
	local := &model.Shippings{
		ID:               id,
		ReceiverName:     c.PostForm("name"),
		ReceiverMobile:   mobile,
		ReceiverProvince: c.PostForm("province"),
		ReceiverCity:     c.PostForm("city"),
		ReceiverDistrict: c.PostForm("district"),
		ReceiverAddress:  c.PostForm("address"),
		ReceiverZip:      c.PostForm("zip"),
		Update_time:      time.Now(),
	}
	if local.UpdateShipping() {
		tools.JsonMessage(c, 20, "success", local)

	} else {
		tools.JsonMessage(c, 20, "fail", "修改失败")
	}
}

// 根据ID删除地址列表
func DeleteShipping(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println(err)
		return
	}
	local := &model.Shippings{
		ID: id,
	}
	if local.DeleteShipping() {
		tools.JsonMessage(c, 20, "删除地址成功", "success")
	} else {
		tools.JsonMessage(c, 20, "删除地址失败", "fail")

	}
}

func PrintShipping(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	local := &model.Shippings{
		UserId: userid,
	}
	locallist, flag := local.PrintLocalList()
	if flag { //有地址
		if len(locallist) == 0 {
			tools.JsonMessage(c, 22, "success", "地址列表为空")

		} else {
			tools.JsonMessage(c, 20, "success1", locallist)

		}
	} else {
		tools.JsonMessage(c, 21, "fail", "显示地址列表错误")
	}

}
