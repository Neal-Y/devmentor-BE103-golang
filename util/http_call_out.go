package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"shopping-cart/config"
	"shopping-cart/model/datatransfer/user"
	"strconv"
)

func ParseJSONResponse(body []byte, v interface{}) error {
	err := json.Unmarshal(body, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

func GetIDFromPath(c *gin.Context, paramName string) (int, error) {
	idStr := c.Param(paramName)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid " + paramName})
		return 0, err
	}
	return id, nil
}

func ParseIDToken(idToken string) (*user.LineProfileResponse, error) {
	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.LineClientSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse id token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		profileData := &user.LineProfileResponse{
			UserID:      claims["sub"].(string),
			DisplayName: claims["name"].(string),
			Email:       claims["email"].(string),
		}
		return profileData, nil
	} else {
		return nil, fmt.Errorf("failed to parse id token")
	}
}
