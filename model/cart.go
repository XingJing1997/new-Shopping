package model

import (
	"fmt"
	"log"
	"shopping/db"
	"time"
)

type Carts struct { //购物车
	ID          int //主键，自增
	UserId      int //用户ID
	ProductId   int //商品ID
	Quantity    int //数量
	Checked     int //是否勾选 1：勾选，2：未勾选
	Create_time time.Time
	Update_time time.Time
}

//根据用户ID和商品ID，查找某件物品是否在购物车
func (cart *Carts) SelectCart() (cartInfo Carts, flag bool) {
	if db.DB.Model(&User{}).Where("user_id=? and product_id", cart.UserId, cart.ProductId).Find(&cartInfo).RecordNotFound() {
		log.Println("NotFound phone")
		flag = false
	} else {
		log.Println("Found")
		flag = true
	}
	return
}

// 添加购物车
func (cart Carts) AddCart() bool {
	tx := db.DB.Begin()
	if err := tx.Create(&cart).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

// 删除购物车

func (cart *Carts) DeleteCart() bool {
	//var c int
	tx := db.DB.Begin()
	if err := tx.Model(&Carts{}).Where("user_id=? and product_id=?", cart.UserId, cart.ProductId).Delete(Carts{}).Error; err != nil {
		fmt.Println("err:", err)
		//fmt.Println(c)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true

	}
}

//修改购物车
func (cart *Carts) ModifyCarts() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&Carts{}).Where("user_id=? and product_id=?", cart.UserId, cart.ProductId).Update(cart).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

//根据用户ID显示购物车
func (cart *Carts) PrintCarts() (cartList []Carts, flag bool) {
	tx := db.DB.Begin()
	if err := tx.Model(&Carts{}).Where("user_id=?", cart.UserId).Find(&cartList).Error; err != nil {
		flag = false
		tx.Rollback()
	} else {
		flag = true
		tx.Commit()
	}
	return
}

//根据用户ID查找已勾选的购物车
func (cart *Carts) SearchCarts() (cartList []Carts, flag bool) {
	tx := db.DB.Begin()
	if err := tx.Model(&Carts{}).Where("user_id=? and checked=?", cart.UserId, cart.Checked).Find(&cartList).Error; err != nil {
		flag = false
		tx.Rollback()
	} else {
		flag = true
		tx.Commit()
	}
	return
}

//删除多个物品
func (cart *Carts) DeleteMore() (flag bool) {
	tx := db.DB.Begin()
	if err := tx.Model(&Carts{}).Where("user_id=? and checked=?", cart.UserId, cart.Checked).Delete(Carts{}).Error; err != nil {
		flag = false
		tx.Rollback()
	} else {
		flag = true
		tx.Commit()
	}
	return
}
