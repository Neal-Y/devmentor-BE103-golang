package order

import "shopping-cart/model/database"

type Request struct {
	UserID       int                    `json:"user_id"`
	Note         string                 `json:"note"`
	Status       string                 `json:"status"`
	OrderDetails []database.OrderDetail `json:"order_details"`
}
