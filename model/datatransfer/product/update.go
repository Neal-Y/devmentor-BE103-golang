package product

import "time"

type Update struct {
	Name           string    `json:"name"`
	Picture        string    `json:"picture"`
	Price          float64   `json:"price" `
	Stock          int       `json:"stock" `
	Description    string    `json:"description"`
	ExpirationTime time.Time `json:"expiration_time" `
}
