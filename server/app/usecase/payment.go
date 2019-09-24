package usecase

import (
	"time"

	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/pkg/errors"
)

type PaymentUsecase interface {
	GetPaymentList() ([]model.PaymentsForDisp, error)
	CreatePayment(paymentTypeID int, paymentStatusID int, monthID int, amount int) (err error)
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository, MonthRepo repository.MonthRepository) PaymentUsecase {
	return implPaymentUsecase{
		PaymentRepo: paymentRepo,
		MonthRepo:   MonthRepo,
	}
}

type implPaymentUsecase struct {
	PaymentRepo repository.PaymentRepository
	MonthRepo   repository.MonthRepository
}

func (self implPaymentUsecase) GetPaymentList() ([]model.PaymentsForDisp, error) {
	t := time.Now()

	pfdl := make([]model.PaymentsForDisp, 6)
	for i, _ := range pfdl {
		display := t.Format("2006-01")

		m, err := self.MonthRepo.GetByDisplay(display)
		if err != nil {
			return nil, err
		}

		ps, err := self.PaymentRepo.GetListByMonthID(m.ID)
		if err != nil {
			return nil, err
		}

		pfdl[i] = model.PaymentsForDisp{
			Month:    m,
			Payments: ps,
			TotalFee: ps.GetTotalAmount(),
		}

		t = t.AddDate(0, 1, 0)
	}

	return pfdl, nil
}

func (self implPaymentUsecase) CreatePayment(paymentTypeID int, paymentStatusID int, monthID int, amount int) (err error) {
	if monthID == 0 || paymentTypeID == 0 || paymentStatusID == 0 || amount == 0 {
		return errors.New("「どこかに0が入っているな。まさか返済できねえのか？ 死にてぇなら、生命保険、加入してからにしろ。返済がまだだぜ。」")
	}

	p, err := self.PaymentRepo.GetByPaymentTypeIDMonthID(model.PAYMENT_TYPE_ID(paymentTypeID), model.MONTH_ID(monthID))
	if err != nil {
		return errors.New("「そんな金俺は貸してねえ。 月 と 何に使ったのか もう一度確認しろ。」")
	}

	p.PaymentStatusID = model.PAYMENT_STATUS_ID(paymentStatusID)
	p.Amount = amount

	err = self.PaymentRepo.Update(p)
	if err != nil {
		return errors.New("...エラーだ。おい、宿本に連絡しろ。 ギャンブルにハマった奴の明日は信用しねェ。 そこで待っとけ。")
	}

	return nil
}
