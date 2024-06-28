package admin

import (
	"github.com/gin-gonic/gin"
	middleware "shopping-cart/middlerware"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type Admin struct {
	adminService service.AdminService
	userService  service.UserService
}

func NewAdminController(r *gin.RouterGroup) *Admin {
	adminRepo := repository.NewAdminRepository()
	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)
	adminService := service.NewAdminService(adminRepo, userRepo)

	h := &Admin{
		adminService: adminService,
		userService:  userService,
	}

	Register(h, r)

	r.Use(middleware.JWTAuthMiddleware())
	{
		adminRoute(h, r)
		manageUser(h, r)
	}

	return h
}

func Register(h *Admin, r *gin.RouterGroup) {
	r.POST("/admin/register", h.Register)
	r.POST("/admin/login", h.Login)
}

func adminRoute(h *Admin, r *gin.RouterGroup) {
	r.GET("/admin/:id", h.GetAdmin)
	r.PATCH("/admin/:id", h.UpdateAdmin)
	r.DELETE("/admin/:id", h.DeleteAdmin)
}

func manageUser(h *Admin, r *gin.RouterGroup) {
	r.POST("/admin/users", h.CreateUser)
	r.GET("/admin/users/:id", h.GetUser)
	r.PATCH("/admin/users/:id", h.UpdateUser)
	r.DELETE("/admin/users/:id", h.DeleteUser)
}
