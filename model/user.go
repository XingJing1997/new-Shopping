package model

import (
	"fmt"
	"log"
	"shopping/db"
	"time"
)

//var DB = db.DB

//数据层
type User struct {
	ID          int
	Phone       string    //手机号，唯一
	Password    string    //密码
	Role        int       //角色
	Name        string    //昵称
	Sex         string    //性别
	Picture     string    //头像
	Create_time time.Time //创建时间
	Update_time time.Time //修改时间
	Token       string
}

//1. 查找用户是否存在
func (u *User) FoundUser() bool {
	var one User
	if db.DB.Model(&User{}).Where("phone=?", u.Phone).Find(&one).RecordNotFound() {
		log.Println("not found:")

		return false //不存在
	} else {
		log.Println("found:", u.Phone)
	}
	return true //存在
}
func (u *User) ReturnUser() *User {
	var user User
	if db.DB.Model(&User{}).Where("Phone=?", u.Phone).Find(&user).RecordNotFound() {
		log.Println("NotFound phone")
	} else {
		log.Println("Found")
	}
	return &user
}

// 2.验证用户登录
func (u *User) CheckLogin() bool {
	var user User
	if err := db.DB.Model(&User{}).Find(&user).Where("phone=?", u.Phone).Error; err != nil {
		log.Println(err)
		return false
	}
	if u.Password == user.Password {
		return true
	}
	return false
}

// 3.添加用户
func (u *User) AddUser() bool {
	tx := db.DB.Begin()
	err := tx.Create(&u).Error
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
	}
	return true
}

// 4.更新数据
func (u *User) EditUser() bool {
	tx := db.DB.Begin()
	if err := tx.Model(&User{}).Where("phone=?", u.Phone).Update(u).First(&u).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}

// 5. 保存图片
func (u *User) SaveImagePath() bool {
	tx := db.DB.Begin()
	if err := tx.Exec("update users set Picture=? where Phone=?", u.Picture, u.Phone); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 6. 获取用户信息
func (u *User) GetInfo() User {
	var userInfo User
	tx := db.DB.Begin()
	if tx.Model(&User{}).Where("phone=?", u.Phone).Find(&userInfo).RecordNotFound() {
		return userInfo
	} else {
		return userInfo
	}
}
