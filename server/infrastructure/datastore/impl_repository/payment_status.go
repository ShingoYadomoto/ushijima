package impl_repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/dto"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
)

func NewPaymentStatusRepository(db postgres.AbsDB) repository.PaymentStatusRepository {
	return implPaymentStatusRepository{db}
}

type implPaymentStatusRepository struct {
	db postgres.AbsDB
}

func (self implPaymentStatusRepository) GetByID(id model.PAYMENT_STATUS_ID) (*model.PaymentStatus, error) {
	const q = `SELECT id, name, display, create_date, update_date FROM public.payment_statuses WHERE id = $1`

	ps := &dto.PaymentStatus{}
	err := self.db.Get(ps, q, id)
	if err != nil {
		return nil, err
	}

	return &model.PaymentStatus{
		ID:      model.PAYMENT_STATUS_ID(ps.ID),
		Name:    ps.Name.String,
		Display: ps.Display.String,
	}, nil
}

func (self implPaymentStatusRepository) GetAll() ([]model.PaymentStatus, error) {
	const q = `SELECT id, name, display FROM public.payment_statuses ORDER BY id`

	psl := []dto.PaymentStatus{}
	err := self.db.Select(&psl, q)
	if err != nil {
		return nil, err
	}

	ret := make([]model.PaymentStatus, len(psl))
	for i, ps := range psl {
		ret[i] = model.PaymentStatus{
			ID:      model.PAYMENT_STATUS_ID(ps.ID),
			Name:    ps.Name.String,
			Display: ps.Display.String,
		}
	}

	return ret, nil
}
