package impl_repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/dto"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
)

func NewPaymentTypeRepository(db postgres.AbsDB) repository.PaymentTypeRepository {
	return implPaymentTypeRepository{db}
}

type implPaymentTypeRepository struct {
	db postgres.AbsDB
}

func (self implPaymentTypeRepository) GetByID(id model.PAYMENT_TYPE_ID) (*model.PaymentType, error) {
	const q = `SELECT id, name, display FROM public.payment_types WHERE id = $1`

	pt := &dto.PaymentType{}
	err := self.db.Get(pt, q, id)
	if err != nil {
		return nil, err
	}

	return &model.PaymentType{
		ID:      model.PAYMENT_TYPE_ID(pt.ID),
		Name:    pt.Name.String,
		Display: pt.Display.String,
	}, nil
}

func (self implPaymentTypeRepository) GetAll() ([]model.PaymentType, error) {
	const q = `SELECT id, name, display FROM public.payment_types ORDER BY id`

	ptl := []dto.PaymentType{}
	err := self.db.Select(&ptl, q)
	if err != nil {
		return nil, err
	}

	ret := make([]model.PaymentType, len(ptl))
	for i, pt := range ptl {
		ret[i] = model.PaymentType{
			ID:      model.PAYMENT_TYPE_ID(pt.ID),
			Name:    pt.Name.String,
			Display: pt.Display.String,
		}
	}

	return ret, nil
}
