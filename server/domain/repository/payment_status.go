package repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
)

type PaymentStatusRepository interface {
	Insert(*model.PaymentStatus) error
	Update(*model.PaymentStatus) error
	Upsert(*model.PaymentStatus) error
	Delete(*model.PaymentStatus) error
	PaymentStatusByID(id int) (*model.PaymentStatus, error)
	AllPaymentStatuses() (*[]model.PaymentStatus, error)
}
