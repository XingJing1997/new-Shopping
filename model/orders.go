package model

import (
	"fmt"
	"shopping/db"
	"time"
)

type Orderdetail struct {
	ID               int       //主键，自增
	OrderNo          string    //订单号码（索引）
	UserId           int       //用户ID
	ProductId        int       //产品ID
	ProductName      string    //产品名称
	ProductImage     string    //产品图片
	CurrentUnitPrice float64   //订单生成时候的商品单价，保留两位小数
	Quantity         int       //购买商品数量
	TotalPrice       float64   //商品总价
	Create_time      time.Time //创建时间
	Update_time      time.Time //更新时间
}

type Orders struct {
	ID          int       //主键，自增
	OrderNo     string    //订单号码（索引）
	UserId      int       //用户ID
	ShippingId  int       //收货地址ID
	Payment     float64   //总价格
	Postage     float64   //运费
	Status      int       //订单状态 0：待支付，1：代发货，2：待收货，3：待评价
	PaymentTime time.Time //支付时间
	SendTime    time.Time //发货时间
	EndTime     time.Time //交易结束时间
	CloseTime   time.Time //交易关闭时间
	Create_time time.Time //创建时间
	Update_time time.Time //更新时间
	Detail      []Orderdetail
}

//1. 创建订单
func (order *Orders) CreateOrder(orderdetail []Orderdetail) bool {
	tx := db.DB.Begin()
	//fmt.Printf("order:%T\n", &order)
	//fmt.Println(order)
	//fmt.Println(&order)
	if err := tx.Create(&order).Error; err != nil {
		fmt.Println("order err:  ", err)
		tx.Rollback()
		return false
	} else {
		// 2.创建订单详情表
		for _, v := range orderdetail {
			fmt.Printf("%T\n", v)
			fmt.Println(&v)
			if err := tx.Model(&Orderdetail{}).Create(&v).Error; err != nil {
				fmt.Println("orderDetail err:  ", err)
				tx.Rollback()
				return false
			}
		}
		tx.Commit()
		return true

	}
}

// 2.创建订单详情表
func CreateOrderDetail(orderdetail []Orderdetail) bool {
	tx := db.DB.Begin()
	for _, v := range orderdetail {
		if err := tx.Create(&v).Error; err != nil {
			fmt.Println("err:  ", err)
			tx.Rollback()
			return false
		} else {
			tx.Commit()
			return true

		}
	}
	return false
}

// 3. 取消订单 通过订单号
func (order *Orders) DeleteOrder() bool {
	tx := db.DB.Begin()
	fmt.Println(order.OrderNo)
	if err := tx.Model(&Orders{}).Where("order_no=?", order.OrderNo).Delete(Orders{}).Error; err != nil {

		fmt.Println("order:", err)
		tx.Rollback()
		return false
	} else {
		if err := tx.Model(&Orderdetail{}).Where("order_no=?", order.OrderNo).Delete(Orderdetail{}).Error; err != nil {
			fmt.Println("Orderdetail:", err)
			tx.Rollback()
			return false
		}
		tx.Commit()
		return true
	}
}

// 4. 根据订单号,修改订单状态和地址ID
func (order *Orders) UpdateOrder() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&Orders{}).Where("order_no=?", order.OrderNo).Update(order).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

// 5. 显示全部订单,根据用户ID
func (order *Orders) PrintAllOrder() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&Orders{}).Where("user_id=?", order.UserId).Find(&order).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

// 6. 根据用户ID和状态显示订单 :0：待支付，1：代发货，2：待收货，3：待评价
func (order *Orders) PrintOrderStatus() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&Orders{}).Where("user_id=? and status=?", order.UserId, order.Status).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

// 7. 根据用户ID和订单号和商品ID显示订单详情表
func (orderdetail *Orderdetail) PrintOrderDetail() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&Orderdetail{}).Where("user_id=? and order_no=? and product_id=?", orderdetail.UserId, orderdetail.OrderNo, orderdetail.ProductId).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}
