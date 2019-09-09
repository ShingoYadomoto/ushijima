package impl_repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
)

func NewPaymentStatusRepository(db postgres.AbsDB) repository.PaymentStatusRepository {
	return implPaymentStatusRepository{db}
}

type implPaymentStatusRepository struct {
	db postgres.AbsDB
}

// Insert inserts the PaymentStatus to the database.
func (self implPaymentStatusRepository) Insert(ps *model.PaymentStatus) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.payment_statuses (` +
		`name, display, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`

	// run query
	postgres.DBLog(sqlstr, ps.Name, ps.Display, ps.CreateDate, ps.UpdateDate)
	err = self.db.QueryRow(sqlstr, ps.Name, ps.Display, ps.CreateDate, ps.UpdateDate).Scan(&ps.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the PaymentStatus in the database.
func (self implPaymentStatusRepository) Update(ps *model.PaymentStatus) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.payment_statuses SET (` +
		`name, display, create_date, update_date` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE id = $5`

	// run query
	postgres.DBLog(sqlstr, ps.Name, ps.Display, ps.CreateDate, ps.UpdateDate, ps.ID)
	_, err = self.db.Exec(sqlstr, ps.Name, ps.Display, ps.CreateDate, ps.UpdateDate, ps.ID)
	return err
}

// Upsert performs an upsert for PaymentStatus.
//
// NOTE: PostgreSQL 9.5+ only
func (self implPaymentStatusRepository) Upsert(ps *model.PaymentStatus) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.payment_statuses (` +
		`id, name, display, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, name, display, create_date, update_date` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.name, EXCLUDED.display, EXCLUDED.create_date, EXCLUDED.update_date` +
		`)`

	// run query
	postgres.DBLog(sqlstr, ps.ID, ps.Name, ps.Display, ps.CreateDate, ps.UpdateDate)
	_, err = self.db.Exec(sqlstr, ps.ID, ps.Name, ps.Display, ps.CreateDate, ps.UpdateDate)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the PaymentStatus from the database.
func (self implPaymentStatusRepository) Delete(ps *model.PaymentStatus) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.payment_statuses WHERE id = $1`

	// run query
	postgres.DBLog(sqlstr, ps.ID)
	_, err = self.db.Exec(sqlstr, ps.ID)
	if err != nil {
		return err
	}

	return nil
}

// PaymentStatusByID retrieves a row from 'public.payment_statuses' as a PaymentStatus.
//
// Generated from index 'payment_statuses_pkey'.
func (self implPaymentStatusRepository) PaymentStatusByID(id int) (*model.PaymentStatus, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display, create_date, update_date ` +
		`FROM public.payment_statuses ` +
		`WHERE id = $1`

	// run query
	postgres.DBLog(sqlstr, id)
	ps := model.PaymentStatus{}

	err = self.db.QueryRow(sqlstr, id).Scan(&ps.ID, &ps.Name, &ps.Display, &ps.CreateDate, &ps.UpdateDate)
	if err != nil {
		return nil, err
	}

	return &ps, nil
}

func (self implPaymentStatusRepository) AllPaymentStatuses() (*[]model.PaymentStatus, error) {
	const sqlstr = `SELECT ` +
		`* ` +
		`FROM public.payment_statuses ` +
		`ORDER BY id `

	pl := []model.PaymentStatus{}

	err := self.db.Select(&pl, sqlstr)
	if err != nil {

		return nil, err
	}

	return &pl, nil
}
