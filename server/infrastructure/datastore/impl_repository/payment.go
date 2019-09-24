package impl_repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/dto"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
)

func NewPaymentRepository(db postgres.AbsDB) repository.PaymentRepository {
	return implPaymentRepository{db}
}

type implPaymentRepository struct {
	db postgres.AbsDB
}

func (self implPaymentRepository) Update(p *model.Payment) error {
	const q = `
		UPDATE
			public.payments
		SET
			(payment_type_id, payment_status_id, amount, month_id, update_date) = ($1, $2, $3, $4, now())
		WHERE
			id = $5`

	_, err := self.db.Exec(q, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.MonthID, p.ID)
	return err
}

func (self implPaymentRepository) GetByID(id model.PAYMENT_ID) (*model.Payment, error) {
	const q = `
		SELECT
			id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id
		FROM
			public.payments
		WHERE
			id = $1`

	p := &dto.Payment{}
	err := self.db.Get(p, q, id)
	if err != nil {
		return nil, err
	}

	return p.ToModel(), nil
}

func (self implPaymentRepository) GetByPaymentTypeIDMonthID(paymentTypeID model.PAYMENT_TYPE_ID, monthID model.MONTH_ID) (*model.Payment, error) {
	const q = `
		SELECT
			id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id
		FROM
			public.payments
		WHERE
			payment_type_id = $1 AND month_id = $2`

	p := &dto.Payment{}
	err := self.db.Get(p, q, paymentTypeID, monthID)
	if err != nil {
		return nil, err
	}

	return p.ToModel(), nil
}

func (self implPaymentRepository) GetListByMonthID(id model.MONTH_ID) (model.PaymentList, error) {
	const q = `
		SELECT 
			id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id 
		FROM
			public.payments 
		WHERE
			month_id = $1 
		ORDER BY
			payment_type_id`

	pl := []dto.Payment{}
	err := self.db.Select(&pl, q, id)
	if err != nil {
		return nil, err
	}

	ret := model.PaymentList{}
	for _, p := range pl {
		ret = append(ret, *p.ToModel())
	}

	return ret, nil
}
