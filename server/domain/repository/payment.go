package repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"gopkg.in/guregu/null.v3"
)

type PaymentRepository interface {
	Insert(*model.Payment) error
	Update(*model.Payment) error
	Upsert(*model.Payment) error
	Delete(*model.Payment) error
	PaymentByPaymentTypeIDMonthID(null.Int, null.Int) (*model.Payment, error)
	PaymentByID(int) (*model.Payment, error)
	PaymentListByMonthID(id int) (*[]model.Payment, error)
}
