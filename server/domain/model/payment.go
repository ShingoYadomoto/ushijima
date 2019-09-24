package model

type PAYMENT_ID int

type Payment struct {
	ID              PAYMENT_ID        `json:"id"`
	PaymentTypeID   PAYMENT_TYPE_ID   `json:"payment_type_id"`
	PaymentStatusID PAYMENT_STATUS_ID `json:"payment_status_id"`
	MonthID         MONTH_ID          `json:"month_id"`
	Amount          int               `json:"amount"`
}

type PaymentList []Payment

func (pl PaymentList) GetTotalAmount() int {
	var total int
	for _, p := range pl {
		total = total + p.Amount
	}
	return total
}

type PaymentsForDisp struct {
	Month    *Month    `json:"month"`
	Payments []Payment `json:"payments"`
	TotalFee int       `json:"total_fee"`
}
