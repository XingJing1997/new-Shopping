package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"shopping/model"
	"shopping/util/tools"
	"strconv"
	"time"
)

//1.创建订单(直接创建)
func CreateOrder(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	var allOrder model.Orders
	json.Unmarshal([]byte(b), &allOrder)
	createtime := time.Now()
	updatetime := time.Now()
	tt := time.Date(0001, 01, 01, 01, 01, 01, 12, time.Local)
	//生成订单号
	orderno := tools.GetOrderNo(allOrder.UserId)
	allOrder.OrderNo = orderno
	var totalPrice float64
	for k, v := range allOrder.Detail {
		// 1. 查找商品表以获取商品名称和商品图片，商品当前价格
		products := &model.Products{
			ID: v.ProductId,
		}
		product := products.SelectPro() //根据商品ID查找商品获取名称和价格等
		allOrder.Detail[k].ProductName = product.Name
		allOrder.Detail[k].ProductImage = product.MainImage
		allOrder.Detail[k].CurrentUnitPrice = product.Price
		allOrder.Detail[k].OrderNo = allOrder.OrderNo
		allOrder.Detail[k].UserId = allOrder.UserId
		allOrder.Detail[k].Create_time = createtime
		allOrder.Detail[k].Update_time = updatetime
		quantity := float64(v.Quantity)
		allOrder.Detail[k].TotalPrice = allOrder.Detail[k].CurrentUnitPrice * quantity
		totalPrice += allOrder.Detail[k].TotalPrice
	}

	allOrder.Payment = totalPrice
	allOrder.Create_time = createtime
	allOrder.Update_time = updatetime
	allOrder.PaymentTime = tt
	allOrder.CloseTime = tt
	allOrder.EndTime = tt
	allOrder.SendTime = tt
	if allOrder.CreateOrder(allOrder.Detail) {
		tools.JsonMessage(c, 20, "success", allOrder)
	} else {
		tools.JsonMessage(c, 21, "fail", "订单提交失败")
	}
}

//2.通过购物车创建订单（参数为用户ID，地址ID，运费，状态）
func CreateOrderByCart(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	var allOrder model.Orders
	json.Unmarshal([]byte(b), &allOrder)
	//1.遍历购物车，查找是否勾选
	cart := &model.Carts{
		UserId:  allOrder.UserId,
		Checked: 1,
	}
	cartlist, flag := cart.PrintCarts() //已勾选的物品
	if flag {
		if len(cartlist) == 0 {
			tools.JsonMessage(c, 22, "success", "购物车为空，快点填充吧")

		} else {
			createtime := time.Now()
			updatetime := time.Now()
			tt := time.Date(0001, 01, 01, 01, 01, 01, 12, time.Local)
			//生成订单号
			orderno := tools.GetOrderNo(allOrder.UserId)
			allOrder.OrderNo = orderno
			var totalPrice float64
			for k, v := range cartlist { //遍历选中商品列表获信息
				// 1. 查找商品表以获取商品名称和商品图片，商品当前价格
				products := &model.Products{
					ID: v.ProductId,
				}
				product := products.SelectPro() //根据商品ID查找商品获取名称和价格等
				allOrder.Detail[k].ProductName = product.Name
				allOrder.Detail[k].ProductImage = product.MainImage
				allOrder.Detail[k].CurrentUnitPrice = product.Price
				allOrder.Detail[k].OrderNo = allOrder.OrderNo
				allOrder.Detail[k].UserId = allOrder.UserId
				allOrder.Detail[k].Create_time = createtime
				allOrder.Detail[k].Update_time = updatetime
				quantity := float64(v.Quantity)
				allOrder.Detail[k].TotalPrice = allOrder.Detail[k].CurrentUnitPrice * quantity
				totalPrice += allOrder.Detail[k].TotalPrice
			}

			allOrder.Payment = totalPrice
			allOrder.Create_time = createtime
			allOrder.Update_time = updatetime
			allOrder.PaymentTime = tt
			allOrder.CloseTime = tt
			allOrder.EndTime = tt
			allOrder.SendTime = tt
			if allOrder.CreateOrder(allOrder.Detail) {
				if cart.DeleteMore() {
					tools.JsonMessage(c, 20, "success", allOrder)
				} else {
					tools.JsonMessage(c, 21, "购物车里的删除失败", allOrder)
				}
			} else {
				tools.JsonMessage(c, 21, "fail", "订单提交失败")
			}
		}
	} else {
		tools.JsonMessage(c, 21, "success", "购物车操作错误")
	}

}

// 3. 删除订单 通过订单号
func DeleteOrder(c *gin.Context) {
	orderno := c.Query("orderno")
	if len(orderno) == 0 {
		fmt.Println("orderno.len==0")
		return
	}
	order := &model.Orders{
		OrderNo: orderno,
	}
	if order.DeleteOrder() {
		tools.JsonMessage(c, 20, "success", "订单删除成功")
	} else {
		tools.JsonMessage(c, 21, "fail", "订单删除失败")

	}
}

// 4. 根据订单号,修改订单状态和地址ID
func UpdateOrder(c *gin.Context) {
	orderno := c.Query("orderno")
	status, err := strconv.Atoi(c.Query("status"))
	localid, err := strconv.Atoi(c.Query("localid"))
	if err != nil {
		fmt.Println(err)
	}
	order := &model.Orders{
		OrderNo:     orderno,
		Status:      status,
		ShippingId:  localid,
		Update_time: time.Now(),
	}
	if order.UpdateOrder() {
		tools.JsonMessage(c, 20, "success", "订单修改成功")
	} else {
		tools.JsonMessage(c, 21, "fail", "订单修改失败")

	}
}

// 5. 显示全部订单,根据用户ID
func PrintAllOrder(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Println(err)
	}
	order := &model.Orders{
		UserId: userid,
	}
	if order.PrintAllOrder() {
		tools.JsonMessage(c, 20, "success", order)

	} else {
		tools.JsonMessage(c, 21, "fail", "显示订单失败")

	}

}

// 6. 根据用户ID和状态显示订单 :0：待支付，1：代发货，2：待收货，3：待评价
func PrintOrderStatus(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		fmt.Println(err)
		return
	}
	order := &model.Orders{
		UserId: userid,
		Status: status,
	}
	if order.PrintOrderStatus() {
		tools.JsonMessage(c, 20, "success", order)
	} else {
		tools.JsonMessage(c, 21, "fail", "显示状态订单失败")

	}
}

// 7. 根据用户ID和订单号和商品ID显示订单详情表
func PrintOrderDetail(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	orderno := c.Query("orderno")
	productid, err := strconv.Atoi(c.Query("productid"))
	if err != nil {
		fmt.Println(err)
		return
	}
	detail := &model.Orderdetail{
		UserId:    userid,
		OrderNo:   orderno,
		ProductId: productid,
	}
	if detail.PrintOrderDetail() {
		tools.JsonMessage(c, 20, "success", detail)

	} else {
		tools.JsonMessage(c, 21, "fail", "显示订单详情表失败")

	}

}
