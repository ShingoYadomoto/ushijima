package model

func MonthByDisplay(db XODB, d string) (*Month, error) {
	const sqlstr = `SELECT ` +
		`id, display, create_date, update_date ` +
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
