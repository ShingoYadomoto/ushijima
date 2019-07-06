package model

func AllPaymentStatuses(db XODB) (*[]PaymentStatus, error) {
	const sqlstr = `SELECT ` +
		`* ` +
		`FROM public.payment_statuses ` +
		`ORDER BY id `

	pl := []PaymentStatus{}

	err := db.Select(&pl, sqlstr)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}
