package impl_repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
	"gopkg.in/guregu/null.v3"
)

func NewPaymentRepository(db postgres.AbsDB) repository.PaymentRepository {
	return implPaymentRepository{db}
}

type implPaymentRepository struct {
	db postgres.AbsDB
}

// Insert inserts the Payment to the database.
func (self implPaymentRepository) Insert(p *model.Payment) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.payments (` +
		`payment_type_id, payment_status_id, amount, create_date, update_date, month_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING id`

	// run query
	postgres.DBLog(sqlstr, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID)
	err = self.db.QueryRow(sqlstr, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID).Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the Payment in the database.
func (self implPaymentRepository) Update(p *model.Payment) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.payments SET (` +
		`payment_type_id, payment_status_id, amount, create_date, update_date, month_id` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE id = $7`

	// run query
	postgres.DBLog(sqlstr, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID, p.ID)
	_, err = self.db.Exec(sqlstr, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID, p.ID)
	return err
}

// Upsert performs an upsert for Payment.
//
// NOTE: PostgreSQL 9.5+ only
func (self implPaymentRepository) Upsert(p *model.Payment) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.payments (` +
		`id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.payment_type_id, EXCLUDED.payment_status_id, EXCLUDED.amount, EXCLUDED.create_date, EXCLUDED.update_date, EXCLUDED.month_id` +
		`)`

	// run query
	postgres.DBLog(sqlstr, p.ID, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID)
	_, err = self.db.Exec(sqlstr, p.ID, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the Payment from the database.
func (self implPaymentRepository) Delete(p *model.Payment) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.payments WHERE id = $1`

	// run query
	postgres.DBLog(sqlstr, p.ID)
	_, err = self.db.Exec(sqlstr, p.ID)
	if err != nil {
		return err
	}

	return nil
}

// PaymentByPaymentTypeIDPaymentID retrieves a row from 'public.payments' as a Payment.
//
// Generated from index 'payment_type_month_unique'.
func (self implPaymentRepository) PaymentByPaymentTypeIDMonthID(paymentTypeID null.Int, monthID null.Int) (*model.Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id ` +
		`FROM public.payments ` +
		`WHERE payment_type_id = $1 AND month_id = $2`

	// run query
	postgres.DBLog(sqlstr, paymentTypeID, monthID)
	p := model.Payment{}

	err = self.db.QueryRow(sqlstr, paymentTypeID, monthID).Scan(&p.ID, &p.PaymentTypeID, &p.PaymentStatusID, &p.Amount, &p.CreateDate, &p.UpdateDate, &p.MonthID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// PaymentByID retrieves a row from 'public.payments' as a Payment.
//
// Generated from index 'payments_pkey'.
func (self implPaymentRepository) PaymentByID(id int) (*model.Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id ` +
		`FROM public.payments ` +
		`WHERE id = $1`

	// run query
	postgres.DBLog(sqlstr, id)
	p := model.Payment{}

	err = self.db.QueryRow(sqlstr, id).Scan(&p.ID, &p.PaymentTypeID, &p.PaymentStatusID, &p.Amount, &p.CreateDate, &p.UpdateDate, &p.MonthID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (self implPaymentRepository) PaymentListByMonthID(id int) (*[]model.Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`* ` +
		`FROM public.payments ` +
		`WHERE month_id = $1 ` +
		`ORDER BY payment_type_id `

	pl := []model.Payment{}

	err = self.db.Select(&pl, sqlstr, id)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}
