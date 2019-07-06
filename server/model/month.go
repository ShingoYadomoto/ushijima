package model

func AllMonths(db XODB) (*[]Month, error) {
	const sqlstr = `SELECT 
		* 
		FROM public.months
		ORDER BY id`

	ptl := []Month{}

	err := db.Select(&ptl, sqlstr)
	if err != nil {
		return nil, err
	}

	return &ptl, nil
}

func MonthByDisplay(db XODB, d string) (*Month, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.months ` +
		`WHERE display = $1`

	// run query
	XOLog(sqlstr, d)
	m := Month{
		_exists: true,
	}

	err := db.QueryRow(sqlstr, d).Scan(&m.ID, &m.Display, &m.CreateDate, &m.UpdateDate)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
