package repository

import "github.com/ShingoYadomoto/ushijima/server/domain/model"

type MonthRepository interface {
	Insert(*model.Month) error
	Update(*model.Month) error
	Upsert(*model.Month) error
	Delete(*model.Month) error
	MonthByID(int) (*model.Month, error)
	AllMonths() (*[]model.Month, error)
	MonthByDisplay(string) (*model.Month, error)
}
