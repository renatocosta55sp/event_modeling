package http

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	controller HttpServer) {
	r.POST("/products", controller.CreateProduct)
	r.DELETE("/products/:code", controller.DeleteProduct)
	r.PUT("/products/:code", controller.UpdateProduct)
	r.GET("/products", controller.GetProducts)
}
