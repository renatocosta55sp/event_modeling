package http

type CreateProductRequest struct {
	Code       int     `json:"code"  binding:"required"`
	Name       string  `json:"name"  binding:"required"`
	Stock      int     `json:"stock"  binding:"required"`
	TotalStock int     `json:"total_stock"  binding:"required"`
	CutStock   int     `json:"cut_stock"  binding:"required"`
	PriceFrom  float64 `json:"price_from"  binding:"required"`
	PriceTo    float64 `json:"price_to"  binding:"required"`
}

type UpdateProductRequest struct {
	Name       string  `json:"name"  binding:"required"`
	Stock      int     `json:"stock"  binding:"required"`
	TotalStock int     `json:"total_stock"  binding:"required"`
	CutStock   int     `json:"cut_stock"  binding:"required"`
	PriceFrom  float64 `json:"price_from"  binding:"required"`
	PriceTo    float64 `json:"price_to"  binding:"required"`
}
