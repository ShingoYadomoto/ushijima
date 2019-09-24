package repository

import "github.com/ShingoYadomoto/ushijima/server/domain/model"

type PaymentTypeRepository interface {
	GetByID(model.PAYMENT_TYPE_ID) (*model.PaymentType, error)
	GetAll() ([]model.PaymentType, error)
}
