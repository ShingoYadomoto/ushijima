package repository

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
)

type PaymentRepository interface {
	Update(*model.Payment) error
	GetByID(model.PAYMENT_ID) (*model.Payment, error)
	GetByPaymentTypeIDMonthID(model.PAYMENT_TYPE_ID, model.MONTH_ID) (*model.Payment, error)
	GetListByMonthID(model.MONTH_ID) (model.PaymentList, error)
}
