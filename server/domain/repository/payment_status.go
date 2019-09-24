package repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
)

type PaymentStatusRepository interface {
	GetByID(model.PAYMENT_STATUS_ID) (*model.PaymentStatus, error)
	GetAll() ([]model.PaymentStatus, error)
}
