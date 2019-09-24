package model

type PAYMENT_STATUS_ID int

type PaymentStatus struct {
	ID      PAYMENT_STATUS_ID `json:"id"`
	Name    string            `json:"name"`
	Display string            `json:"display"`
}
