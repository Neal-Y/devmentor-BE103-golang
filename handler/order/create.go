package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/datatransfer"
)

func (h *Order) CreateOrder(c *gin.Context) {
	var orderRequest datatransfer.OrderRequest
	err := c.ShouldBindJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.CreateOrder(&orderRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": order})
}
