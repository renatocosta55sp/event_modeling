package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const requestDataKey = "requestData"

type CreateProductRequest struct {
	Code       int     `json:"code"  binding:"required"`
	Name       string  `json:"name"  binding:"required"`
	Stock      int     `json:"stock"  binding:"required"`
	TotalStock int     `json:"total_stock"  binding:"required"`
	CutStock   int     `json:"cut_stock"  binding:"required"`
	PriceFrom  float64 `json:"price_from"  binding:"required"`
	PriceTo    float64 `json:"price_to"  binding:"required"`
}

func createProductValidator(ctx *gin.Context) {

	var requestData CreateProductRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Set("requestData", requestData)

	ctx.Next()

}

type UpdateProductRequest struct {
	Id         string
	Name       string  `json:"name"  binding:"required"`
	Stock      int     `json:"stock"  binding:"required"`
	TotalStock int     `json:"total_stock"  binding:"required"`
	CutStock   int     `json:"cut_stock"  binding:"required"`
	PriceFrom  float64 `json:"price_from"  binding:"required"`
	PriceTo    float64 `json:"price_to"  binding:"required"`
}

type DeleteProductRequest struct {
	Id string
}

func updateProductValidator(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		ctx.Abort()
		return
	}

	var requestData UpdateProductRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorUpd": err.Error()})
		ctx.Abort()
		return
	}

	requestData.Id = id
	ctx.Set("requestData", requestData)

	ctx.Next()

}

func deleteProductValidator(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorD": "ID parameter is required"})
		ctx.Abort()
		return
	}

	var requestData DeleteProductRequest

	requestData.Id = id
	ctx.Set("requestData", requestData)

	ctx.Next()

}
