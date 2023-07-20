package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdatIT/gin-gorm-restapi/src/config"
	"github.com/tdatIT/gin-gorm-restapi/src/model"
	"net/http"
)

func AddNewProduct(ctx *gin.Context) {
	var body model.Product
	if err := ctx.BindJSON(&body); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		fmt.Println(err.Error())
		return
	}
	if result := config.DB.Create(&body); result.Error != nil {
		ctx.String(http.StatusBadRequest, result.Error.Error())
		fmt.Println(result.Error.Error())
	} else {
		ctx.JSON(http.StatusOK, body)
	}
}
func GetAllProduct(ctx *gin.Context) {
	var products []model.Product
	if err := config.DB.Find(&products); err.Error != nil {
		ctx.String(http.StatusNotFound, err.Error.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)
}
func GetProductByProductId(ctx *gin.Context) {
	productId := ctx.Param("productId")
	var product model.Product
	if err := config.DB.First(&product, productId); err.Error != nil {
		ctx.String(http.StatusNotFound, "Not found product")
		return
	}
	ctx.JSON(http.StatusOK, product)
}
func DeleteProductByProductId(ctx *gin.Context) {
	productId := ctx.Param("productId")
	if err := config.DB.Delete(&model.Product{}, productId); err.Error != nil {
		ctx.String(http.StatusOK, "Delete product was failed cause:"+err.Error.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Delete has been successful")
}

func UpdateProduct(ctx *gin.Context) {
	var body model.Product

	if err := ctx.BindJSON(&body); err != nil {
		ctx.String(http.StatusOK, "Update failed cause:"+err.Error())
		return
	}
	var product model.Product
	err := config.DB.First(&product, body.ProductId)
	if err != nil {
		product.Name = body.Name
		product.Price = body.Price
		if err := config.DB.Save(&product); err.Error != nil {
			ctx.String(http.StatusOK, "Update product was failed cause:"+err.Error.Error())
			return
		}
	}

	var jsonResp = map[string]interface{}{
		"message":     "Update has been successful",
		"update_data": product,
	}
	ctx.JSON(http.StatusOK, jsonResp)

}
