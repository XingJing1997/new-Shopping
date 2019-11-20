package model

import (
	"fmt"
	"shopping/db"
	"time"
)

type Shippings struct {
	ID               int    //主键，自增
	UserId           int    //用户ID
	ReceiverName     string //收货姓名
	ReceiverMobile   string //收货电话
	ReceiverProvince string //省份
	ReceiverCity     string //城市
	ReceiverDistrict string //区/县
	ReceiverAddress  string //详细地址
	ReceiverZip      string //邮编
	Create_time      time.Time
	Update_time      time.Time
}

func (local *Shippings) AddShipping() bool {
	tx := db.DB.Begin()
	if err := tx.Create(&local).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

//根据主键 ID 修改
func (local *Shippings) UpdateShipping() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&Shippings{}).Where("id=?", local.ID).Update(local).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

//根据ID删除地址列表
func (local *Shippings) DeleteShipping() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&Shippings{}).Where("id=?", local.ID).Delete(Shippings{}).Error; err != nil {

		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

//根据用户ID显示地址列表
func (local *Shippings) PrintLocalList() (localList []Shippings, flag bool) {
	tx := db.DB.Begin()
	if err := tx.Model(&Shippings{}).Where("user_id=?", local.UserId).Find(&localList).Error; err != nil {
		fmt.Println(err)
		flag = false
		tx.Rollback()
	} else {
		fmt.Println("yay")
		flag = true
		tx.Commit()
	}
	return
}

// 根据ID显示地址详情
