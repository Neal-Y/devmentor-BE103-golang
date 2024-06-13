package product

import (
	"github.com/gin-gonic/gin"
)

type Product struct{}

func NewProductController(r *gin.RouterGroup) *Product {
	h := &Product{}
	newRoute(h, r)
	return h
}

func newRoute(h *Product, r *gin.RouterGroup) {
	r.POST("/products", h.CreateProduct)
	r.GET("/products/:id", h.GetProduct)
	r.GET("/products", h.GetAllProducts)
	r.PATCH("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
}