package model

import (
	"fmt"
	"shopping/db"
	"time"
)

//商品表
type Products struct {
	ID          int     //主键，自增
	CategoryId  int     //分类ID，对应分类列表主键
	Name        string  //商品名称
	Subtitle    string  //商品副标题
	MainImage   string  //商品主图
	SubImages   string  //其他图片
	Detail      string  //商品详情
	Price       float64 //价格
	Stock       int     //商品库存
	Status      int     //商品状态 1：在售，2：下架，3：删除
	Create_time time.Time
	Update_time time.Time
}

// 1. 添加商品
func (product *Products) AddProduct() {

}

// 2. 展示商品列表
func (proInfo Products) DisplayProduct(msgNum, PageIndex int) (info []Products, dis bool) {
	tx := db.DB.Begin()
	allMsg := msgNum * PageIndex
	fmt.Println("allmsg:", allMsg)
	if err := tx.Model(&Products{}).Where("id>=?", allMsg).Limit(msgNum).Find(&info).Error; err != nil {
		fmt.Println(err)
		dis = false
	} else {
		dis = true
		fmt.Println(info)
	}
	return
}

// 3. 根据商品ID查找商品
func (proInfo Products) SelectPro() (pro Products) {
	tx := db.DB.Begin()
	if err := tx.Model(&Products{}).Where("id=?", proInfo.ID).Find(&pro).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}
