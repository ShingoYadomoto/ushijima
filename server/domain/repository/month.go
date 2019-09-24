package repository

import "github.com/ShingoYadomoto/ushijima/server/domain/model"

type MonthRepository interface {
	GetByID(model.MONTH_ID) (*model.Month, error)
	GetByDisplay(string) (*model.Month, error)
	GetAll() ([]model.Month, error)
}
