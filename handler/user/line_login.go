package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/config"
	"shopping-cart/constant"
	"shopping-cart/util"
)

func (h *User) LineLogin(c *gin.Context) {
	state := "randomStateString"
	lineURL := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=profile%%20openid%%20email", constant.LineAuthURL, config.AppConfig.LineClientID, config.AppConfig.LineRedirectURI, state)
	c.Redirect(http.StatusFound, lineURL)
}

func (h *User) LineCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if state != "randomStateString" {
		handleLineServerError(c, "invalid state")
		return
	}

	user, err := h.userService.ExchangeTokenAndGetProfile(code)
	if err != nil {
		handleLineServerError(c, err.Error())
		return
	}

	err = h.userService.SaveOrUpdateUser(user)
	if err != nil {
		handleLineServerError(c, "failed to save or update user")
		return
	}

	token, err := util.GenerateJWT(constant.UserType)
	if err != nil {
		handleLineServerError(c, "failed to generate JWT")
		return
	}

	redirectURL := fmt.Sprintf("%s/buffer?token=%s&user_id=%d", config.AppConfig.NgrokURL, token, user.ID)
	c.Redirect(http.StatusFound, redirectURL)
}

func handleLineServerError(c *gin.Context, errorMessage string) {
	fmt.Printf("handleLineServerError called with message: %s\n", errorMessage)
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/error?message=%s", config.AppConfig.NgrokURL, errorMessage))
}
