package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shopping/model"
	"shopping/util/tools"
	"strconv"
	"time"
)

// 添加到购物车
func AddUserCarts(c *gin.Context) {
	userid, err := strconv.Atoi(c.PostForm("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	productid, err := strconv.Atoi(c.PostForm("productid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		fmt.Println(err)
		return
	}

	CartProduct := &model.Carts{
		UserId:      userid,    //用户ID
		ProductId:   productid, //商品ID
		Quantity:    quantity,  //数量
		Checked:     2,         //是否勾选 1：勾选，2：未勾选
		Create_time: time.Now(),
		Update_time: time.Now(),
	}

	cp, flag := CartProduct.SelectCart()
	if flag == false { //购物车里不存在该用户ID和商品ID
		if CartProduct.AddCart() {
			tools.JsonMessage(c, 20, "success in add", "添加购物车成功")
			return
		} else {
			tools.JsonMessage(c, 21, "fail", "添加购物车失败")
		}
	} else {
		CartProduct.Quantity = cp.Quantity + quantity
		if CartProduct.ModifyCarts() {
			tools.JsonMessage(c, 20, "success in Quantity", "添加购物车成功")
		} else {
			tools.JsonMessage(c, 21, "fail", "添加购物车失败")
		}
	}
}

// 删除购物车里的商品(单个)
func DeleteUserCarts(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	productid, err := strconv.Atoi(c.Query("productid"))
	if err != nil {
		fmt.Println(err)
		return
	}

	cart := &model.Carts{
		UserId:    userid,
		ProductId: productid,
	}
	if cart.DeleteCart() {
		tools.JsonMessage(c, 20, "success", "购物车删除成功")
	} else {
		tools.JsonMessage(c, 21, "fail", "购物车删除失败")
	}
}

// 修改购物车
func UpdateCarts(c *gin.Context) {
	userid, err := strconv.Atoi(c.PostForm("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	productid, err := strconv.Atoi(c.PostForm("productid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		fmt.Println(err)
		return
	}
	checked, err := strconv.Atoi(c.PostForm("checked"))
	//t := tools.NowTime()
	//time, err := time.ParseInLocation(t, t, time.Local)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	CartProduct := &model.Carts{
		UserId:      userid,    //用户ID
		ProductId:   productid, //商品ID
		Quantity:    quantity,  //数量
		Checked:     checked,   //是否勾选 1：勾选，2：未勾选
		Update_time: time.Now(),
	}
	if CartProduct.ModifyCarts() {
		tools.JsonMessage(c, 20, "success", CartProduct.Update_time)
	} else {
		tools.JsonMessage(c, 20, "fail", "修改购物车失败")

	}
}

//显示购物车列表
func PrintCarts(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	cart := &model.Carts{
		UserId: userid,
	}
	cartlist, flag := cart.PrintCarts()
	if flag {
		if len(cartlist) == 0 {
			tools.JsonMessage(c, 22, "success", "购物车为空，快点填充吧")

		} else {
			tools.JsonMessage(c, 20, "success", cartlist)
		}
	} else {
		tools.JsonMessage(c, 21, "success", "显示购物车操作错误")
	}

}

//删除购物车里的商品（多个）
func DeleteMore(c *gin.Context) {

}
