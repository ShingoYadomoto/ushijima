package model

func PaymentListByPaymentSummaryID(db XODB, id int) (*[]Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, payment_type_id, payment_status_id, amount, create_date, update_date, payment_summary_id ` +
		`FROM public.payments ` +
		`WHERE payment_summary_id = $1`

	pl := []Payment{}

	err = db.Select(&pl, sqlstr, id)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}
