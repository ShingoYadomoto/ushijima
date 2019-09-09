package impl_repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
)

func NewMonthRepository(db postgres.AbsDB) repository.MonthRepository {
	return implMonthRepository{db}
}

type implMonthRepository struct {
	db postgres.AbsDB
}

// Insert inserts the Month to the database.
func (self implMonthRepository) Insert(m *model.Month) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.months (` +
		`display, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	postgres.DBLog(sqlstr, m.Display, m.CreateDate, m.UpdateDate)
	err = self.db.QueryRow(sqlstr, m.Display, m.CreateDate, m.UpdateDate).Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the Month in the database.
func (self implMonthRepository) Update(m *model.Month) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.months SET (` +
		`display, create_date, update_date` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	postgres.DBLog(sqlstr, m.Display, m.CreateDate, m.UpdateDate, m.ID)
	_, err = self.db.Exec(sqlstr, m.Display, m.CreateDate, m.UpdateDate, m.ID)
	return err
}

// Upsert performs an upsert for Month.
//
// NOTE: PostgreSQL 9.5+ only
func (self implMonthRepository) Upsert(m *model.Month) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.months (` +
		`id, display, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, display, create_date, update_date` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.display, EXCLUDED.create_date, EXCLUDED.update_date` +
		`)`

	// run query
	postgres.DBLog(sqlstr, m.ID, m.Display, m.CreateDate, m.UpdateDate)
	_, err = self.db.Exec(sqlstr, m.ID, m.Display, m.CreateDate, m.UpdateDate)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the Month from the database.
func (self implMonthRepository) Delete(m *model.Month) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.months WHERE id = $1`

	// run query
	postgres.DBLog(sqlstr, m.ID)
	_, err = self.db.Exec(sqlstr, m.ID)
	if err != nil {
		return err
	}

	return nil
}

// MonthByID retrieves a row from 'public.months' as a Month.
//
// Generated from index 'months_pkey'.
func (self implMonthRepository) MonthByID(id int) (*model.Month, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, display, create_date, update_date ` +
		`FROM public.months ` +
		`WHERE id = $1`

	// run query
	postgres.DBLog(sqlstr, id)
	m := new(model.Month)

	err = self.db.QueryRow(sqlstr, id).Scan(m.ID, m.Display, m.CreateDate, m.UpdateDate)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (self implMonthRepository) AllMonths() (*[]model.Month, error) {
	const sqlstr = `SELECT 
		* 
		FROM public.months
		ORDER BY id`

	ptl := []model.Month{}

	err := self.db.Select(&ptl, sqlstr)
	if err != nil {
		return nil, err
	}

	return &ptl, nil
}

func (self implMonthRepository) MonthByDisplay(d string) (*model.Month, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.months ` +
		`WHERE display = $1`

	// run query
	postgres.DBLog(sqlstr, d)
	m := new(model.Month)

	err := self.db.QueryRow(sqlstr, d).Scan(m.ID, m.Display, m.CreateDate, m.UpdateDate)
	if err != nil {
		return nil, err
	}

	return m, nil
}
