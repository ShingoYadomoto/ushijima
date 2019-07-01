package model

func PaymentListByMonthID(db XODB, id int) (*[]Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`* ` +
		`FROM public.payments ` +
		`WHERE month_id = $1 ` +
		`ORDER BY payment_type_id `

	pl := []Payment{}

	err = db.Select(&pl, sqlstr, id)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}
