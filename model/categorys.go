package model

import "time"

//产品分类表
type Categorys struct {
	ID          int    //主键，自增
	PatentId    int    //父类ID
	Name        string //分类名称
	Status      int    //状态，1 正常，2 废弃
	Create_time time.Time
	Update_time time.Time
}
