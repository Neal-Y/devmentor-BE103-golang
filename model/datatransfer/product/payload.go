package product

import "time"

type Payload struct {
	Name           string    `json:"name" binding:"required"`
	Picture        string    `json:"picture"`
	Price          float64   `json:"price" binding:"required"`
	Stock          int       `json:"stock" binding:"required"`
	Description    string    `json:"description"`
	ExpirationTime time.Time `json:"expiration_time" binding:"required"`
}
