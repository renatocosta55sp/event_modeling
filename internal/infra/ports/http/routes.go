package http

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	controller HttpServer) {
	r.POST("/products", createProductValidator, controller.CreateProduct)
	r.DELETE("/products/:id", deleteProductValidator, controller.DeleteProduct)
	r.PUT("/products/:id", updateProductValidator, controller.UpdateProduct)
	r.GET("/products", controller.GetProducts)
}
