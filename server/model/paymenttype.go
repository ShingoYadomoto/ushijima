package model

func AllPaymentTypes(db XODB) (*[]PaymentType, error) {
	const sqlstr = `SELECT 
		* 
		FROM public.payment_types
		ORDER BY id`

	ptl := []PaymentType{}

	err := db.Select(&ptl, sqlstr)
	if err != nil {
		return nil, err
	}

	return &ptl, nil
}
