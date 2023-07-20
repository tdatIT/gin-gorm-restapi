package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tdatIT/gin-gorm-restapi/src/service"
)

func RegisterRouters(r *gin.Engine) {
	routers := r.Group("/api/v1/products")
	routers.POST("/", service.AddNewProduct)
	routers.GET("/", service.GetAllProduct)
	routers.GET("/:productId", service.GetProductByProductId)
	routers.DELETE("/:productId", service.DeleteProductByProductId)
	routers.PUT("/", service.UpdateProduct)
}
