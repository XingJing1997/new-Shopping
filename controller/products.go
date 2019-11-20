package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shopping/model"
	"shopping/util/tools"
	"strconv"
	"time"
)

func CreateProduct(c *gin.Context) {

	categoryId, _ := strconv.Atoi(c.PostForm("category_id"))
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	product := &model.Products{
		CategoryId:  categoryId,
		Name:        c.PostForm("name"),
		Subtitle:    c.PostForm("subtitle"),
		MainImage:   c.PostForm("mainimage"),
		SubImages:   c.PostForm("subimages"),
		Detail:      c.PostForm("detail"),
		Price:       price,
		Stock:       stock,
		Status:      status,
		Create_time: time.Now(),
		Update_time: time.Now(),
	}
	product.AddProduct()

}
func ShowProducts(c *gin.Context) {
	//var products []model.Products
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err)
		return
	}
	var pro model.Products
	products, dis := pro.DisplayProduct(10, pageindex)
	if dis {
		tools.JsonMessage(c, 20, "success", products)
	} else {
		tools.JsonMessage(c, 21, "fail", "获取商品列表失败")
	}
}
