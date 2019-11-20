package model

import (
	"fmt"
	"shopping/db"
	"time"
)

//积分发行商
type PointMaker struct {
	ID     int    //ID
	Name   string //名称
	Desc   string //介绍
	Logo   string //Logo商标图片
	Status int    //发行状态 10-待审核 20-发行成功 30-合作终止
	//发行总量或者历史发行量之后需要添加,历史发行，正在流通，已兑换等
	Total       int     //单次发行总量
	Fate        float64 //通用积分比值
	Create_time time.Time
	Update_time time.Time
}

//用户积分表
type UserPoint struct {
	ID          int       //ID
	PointId     int       //PointID
	UserId      int       //用户ID
	Balance     int       //一个品牌积分余额
	Total       int       //一个品牌积分总量
	Expire_time time.Time //过期时间
	Create_time time.Time //创建时间
	Update_time time.Time //更新时间
}

//发行积分记录在数据库中
type Publish struct {
	ID           int       //主键ID
	PointMakerId int       //积分发行商ID，标识是哪一个积分发行商
	No           int       //发行序号，只是标识Execl的序号而已
	Account      string    //发行给具体账号
	Balance      int       //发行金额
	Pick         int       //0未领取  1领取
	Publish_time time.Time //发行时间
	Expire_time  time.Time //过期时间
}

//1.发型积分
func (pubish *Publish) PushPoint() bool {
	tx := db.DB.Begin()
	if err := tx.Create(&pubish).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

//2.显示用户未领取积分
func (publish *Publish) ShowNoPoint() (flag bool, pu []*Publish) {
	tx := db.DB.Begin()
	if err := tx.Model(&Publish{}).Where("account=? and pick=?", publish.Account, publish.Pick).Find(&pu).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		flag = false
	} else {
		tx.Commit()
		flag = true
	}
	return
}

//4.用户领取积分，参数publish 的ID
func (publish *Publish) GetPoint() bool {
	tx := db.DB.Begin()
	//var p *Publish
	if err := tx.Model(&Publish{}).Where("id=? and pick=?", publish.ID, publish.Pick).First(&publish).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		if err := tx.Model(&Publish{}).Where("id=? and pick=?", publish.ID, publish.Pick).Update(&Publish{Pick: 1}).Error; err != nil {
			fmt.Println("33", err)
			tx.Rollback()
			return false
		} else {
			tx.Commit()
			return true

		}
		//tx.Commit()
		//return true
	}
}

//5.积分领取成功后，将其加入用户积分表中
func (userPoint *UserPoint) AddUserPoint() bool {
	tx := db.DB.Begin()
	if err := tx.Create(&userPoint).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

//6.计算积分总量
func (userPoint *UserPoint) CalculateTotal() (total int, flag bool) {
	// 1.先获取当前积分总量,根据用户ID
	var temp *UserPoint
	tx := db.DB.Begin()
	if err := tx.Model(&UserPoint{}).Where("user_id=?", userPoint.UserId).Find(&temp).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		flag = false
	} else {
		tx.Commit()
		flag = true
	}
	total = temp.Total + userPoint.Balance
	return
}

//7.创建用户积分表
func (userPoint *UserPoint) CreateUserPoint() bool {
	tx := db.DB.Begin()
	if err := tx.Create(&userPoint).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

//8.显示用户积分
func (userPoint *UserPoint) ShowUserPoint() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&UserPoint{}).Where("user_id=?", userPoint.UserId).Find(&userPoint).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}
