// Package model contains the types for schema 'public'.
package model

func AllPaymentType(db XODB) (*[]PaymentType, error) {
	const sqlstr = `SELECT 
		id, name, display, create_date, update_date 
		FROM public.payment_types`

	ptl := []PaymentType{}

	err := db.Select(&ptl, sqlstr)
	if err != nil {
		return nil, err
	}

	return &ptl, nil
}
