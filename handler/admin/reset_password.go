package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/datatransfer/util/forgot_pwd"
)

func (h *Admin) GetAdminEmail(c *gin.Context) {
	var req forgot_pwd.GetEmailByUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.adminService.GetAdminByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": admin.Email})
}

func (h *Admin) RequestPasswordReset(c *gin.Context) {
	var req forgot_pwd.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.adminService.RequestPasswordReset(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password reset email sent"})
}

func (h *Admin) ResetPassword(c *gin.Context) {
	var req forgot_pwd.NewPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.adminService.ResetPassword(req.Email, req.Code, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})
}
