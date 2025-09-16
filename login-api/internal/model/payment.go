package model

import "time"

// Payment merepresentasikan satu data pembayaran
type Payment struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"customer_name"`
	Amount       float64   `json:"amount"`
	Status       string    `json:"status"`
	PaymentDate  time.Time `json:"payment_date"`
}