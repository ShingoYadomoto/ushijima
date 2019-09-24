package model

type PAYMENT_TYPE_ID int

type PaymentType struct {
	ID      PAYMENT_TYPE_ID `json:"id"`
	Name    string          `json:"name"`
	Display string          `json:"display"`
}
