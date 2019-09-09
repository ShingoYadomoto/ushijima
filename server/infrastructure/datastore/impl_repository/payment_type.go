package impl_repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
)

func NewPaymentTypeRepository(db postgres.AbsDB) repository.PaymentTypeRepository {
	return implPaymentTypeRepository{db}
}

type implPaymentTypeRepository struct {
	db postgres.AbsDB
}

// Insert inserts the PaymentType to the database.
func (self implPaymentTypeRepository) Insert(pt *model.PaymentType) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.payment_types (` +
		`name, display, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`

	// run query
	postgres.DBLog(sqlstr, pt.Name, pt.Display, pt.CreateDate, pt.UpdateDate)
	err = self.db.QueryRow(sqlstr, pt.Name, pt.Display, pt.CreateDate, pt.UpdateDate).Scan(&pt.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the PaymentType in the database.
func (self implPaymentTypeRepository) Update(pt *model.PaymentType) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.payment_types SET (` +
		`name, display, create_date, update_date` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE id = $5`

	// run query
	postgres.DBLog(sqlstr, pt.Name, pt.Display, pt.CreateDate, pt.UpdateDate, pt.ID)
	_, err = self.db.Exec(sqlstr, pt.Name, pt.Display, pt.CreateDate, pt.UpdateDate, pt.ID)
	return err
}

// Upsert performs an upsert for PaymentType.
//
// NOTE: PostgreSQL 9.5+ only
func (self implPaymentTypeRepository) Upsert(pt *model.PaymentType) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.payment_types (` +
		`id, name, display, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, name, display, create_date, update_date` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.name, EXCLUDED.display, EXCLUDED.create_date, EXCLUDED.update_date` +
		`)`

	// run query
	postgres.DBLog(sqlstr, pt.ID, pt.Name, pt.Display, pt.CreateDate, pt.UpdateDate)
	_, err = self.db.Exec(sqlstr, pt.ID, pt.Name, pt.Display, pt.CreateDate, pt.UpdateDate)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the PaymentType from the database.
func (self implPaymentTypeRepository) Delete(pt *model.PaymentType) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.payment_types WHERE id = $1`

	// run query
	postgres.DBLog(sqlstr, pt.ID)
	_, err = self.db.Exec(sqlstr, pt.ID)
	if err != nil {
		return err
	}

	return nil
}

// PaymentTypeByID retrieves a row from 'public.payment_types' as a PaymentType.
//
// Generated from index 'payment_types_pkey'.
func (self implPaymentTypeRepository) PaymentTypeByID(id int) (*model.PaymentType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display, create_date, update_date ` +
		`FROM public.payment_types ` +
		`WHERE id = $1`

	// run query
	postgres.DBLog(sqlstr, id)
	pt := &model.PaymentType{}

	err = self.db.QueryRow(sqlstr, id).Scan(&pt.ID, &pt.Name, &pt.Display, &pt.CreateDate, &pt.UpdateDate)
	if err != nil {
		return nil, err
	}

	return pt, nil
}

func (self implPaymentTypeRepository) AllPaymentTypes() (*[]model.PaymentType, error) {
	const sqlstr = `SELECT 
		* 
		FROM public.payment_types
		ORDER BY id`

	ptl := []model.PaymentType{}

	err := self.db.Select(&ptl, sqlstr)
	if err != nil {
		return nil, err
	}

	return &ptl, nil
}
