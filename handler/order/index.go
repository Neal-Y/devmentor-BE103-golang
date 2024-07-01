package order

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type Order struct {
	orderService service.OrderService
}

func NewOrderHandler(r *gin.RouterGroup) *Order {
	orderRepo := repository.NewOrderRepository()
	productRepo := repository.NewProductRepository()

	orderService := service.NewOrderService(orderRepo, productRepo)

	h := &Order{
		orderService: orderService,
	}

	newRoute(h, r)
	return h
}

func newRoute(h *Order, r *gin.RouterGroup) {
	r.POST("/orders", h.CreateOrder)
	r.GET("/orders/:id", h.GetOrder)
	r.PATCH("/orders/:id", h.UpdateOrder)
	r.DELETE("/orders/:id", h.DeleteOrder)
	r.GET("/orders", h.ListOrders)
}
