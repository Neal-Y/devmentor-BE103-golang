package order

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/constant"
	"shopping-cart/middleware"
	"shopping-cart/service"
)

type Order struct {
	orderService service.OrderService
}

func NewOrderHandler(r *gin.RouterGroup, orderService service.OrderService) *Order {
	h := &Order{
		orderService: orderService,
	}

	adminRoute(h, r)
	newRoute(h, r)

	return h
}

func newRoute(h *Order, r *gin.RouterGroup) {
	orderRoute := r.Group("/orders")
	orderRoute.POST("/history", h.ListHistoryOrdersByUser)
	orderRoute.Use(middleware.JWTAuthMiddleware(constant.UserType))
	orderRoute.POST("", h.CreateOrder)
	orderRoute.DELETE("/:id", h.DeleteOrder)
}

func adminRoute(h *Order, r *gin.RouterGroup) {
	adminRoute := r.Group("/orders")
	adminRoute.Use(middleware.JWTAuthMiddleware(constant.AdminType))
	adminRoute.PATCH("/:id", h.UpdateOrder)
	adminRoute.GET("/search", h.SearchOrders)
	adminRoute.GET("/revenue", h.GetRevenue)
}
