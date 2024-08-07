package route

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/handler/admin"
	adminRender "shopping-cart/handler/admin/render"
	"shopping-cart/handler/order"
	"shopping-cart/handler/product"
	productRender "shopping-cart/handler/product/render"
	"shopping-cart/handler/user"
	userRender "shopping-cart/handler/user/render"
	"shopping-cart/repository"
	"shopping-cart/service"
	"shopping-cart/util"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()
	server.Use(gin.Logger())
	server.LoadHTMLGlob("frontend/*")

	adminRender.RegisterHomeRoutes(server)
	userRender.RegisterUserHomeRoutes(server)
	productRender.ProductPage(server)

	api := server.Group("/api")

	orderRepo := repository.NewOrderRepository()
	productRepo := repository.NewProductRepository()
	userRepo := repository.NewUserRepository()
	notificationService := service.NewNotificationService()
	notificationCache := util.NewNotificationCache()
	adminRepo := repository.NewAdminRepository()
	verifyRepo := repository.NewVerifyRepository()

	adminService := service.NewAdminService(adminRepo, verifyRepo)
	orderService := service.NewOrderService(orderRepo, productRepo, userRepo, adminRepo, notificationService, notificationCache)
	productService := service.NewProductService(productRepo, notificationCache)
	userService := service.NewUserService(userRepo, orderRepo, verifyRepo)

	product.NewProductController(api, productService)
	order.NewOrderHandler(api, orderService)

	user.NewAuthorization(api, userService)
	admin.NewAdminController(api, adminService)

	return server
}
