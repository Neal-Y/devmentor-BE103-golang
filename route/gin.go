package route

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/handler/order"
	"shopping-cart/handler/product"
	"shopping-cart/handler/user"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()

	user.RegisterHomeRoutes(server)

	api := server.Group("/api")

	user.NewAuthorization(api)
	product.NewProductController(api)
	order.NewOrderHandler(api)

	return server
}
