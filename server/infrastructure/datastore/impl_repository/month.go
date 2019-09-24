package impl_repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/dto"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
)

func NewMonthRepository(db postgres.AbsDB) repository.MonthRepository {
	return implMonthRepository{db}
}

type implMonthRepository struct {
	db postgres.AbsDB
}

func (self implMonthRepository) GetByID(id model.MONTH_ID) (*model.Month, error) {
	const q = `SELECT id, display FROM public.months WHERE id = $1`

	m := &dto.Month{}
	err := self.db.Get(m, q, id)
	if err != nil {
		return nil, err
	}

	return &model.Month{
		ID:      model.MONTH_ID(m.ID),
		Display: m.Display.String,
	}, nil
}

func (self implMonthRepository) GetByDisplay(d string) (*model.Month, error) {
	const q = `SELECT id, display FROM public.months WHERE display = $1`

	m := &dto.Month{}
	err := self.db.Get(m, q, d)
	if err != nil {
		return nil, err
	}

	return &model.Month{
		ID:      model.MONTH_ID(m.ID),
		Display: m.Display.String,
	}, nil
}

func (self implMonthRepository) GetAll() ([]model.Month, error) {
	const q = `SELECT id, display FROM public.months ORDER BY id`

	ml := []dto.Month{}
	err := self.db.Select(&ml, q)
	if err != nil {
		return nil, err
	}

	ret := make([]model.Month, len(ml))
	for i, m := range ml {
		ret[i] = model.Month{
			ID:      model.MONTH_ID(m.ID),
			Display: m.Display.String,
		}
	}

	return ret, nil
}
