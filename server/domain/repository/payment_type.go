package repository

import "github.com/ShingoYadomoto/ushijima/server/domain/model"

type PaymentTypeRepository interface {
	Insert(*model.PaymentType) error
	Update(*model.PaymentType) error
	Upsert(*model.PaymentType) error
	Delete(*model.PaymentType) error
	PaymentTypeByID(int) (*model.PaymentType, error)
	AllPaymentTypes() (*[]model.PaymentType, error)
}
