package handler

import (
	"net/http"
	"time"

	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/helper"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/guregu/null.v3"
)

type PaymentHandler interface {
	GetPaymentList(c echo.Context) (err error)
	CreatePayment(c echo.Context) (err error)
}

func NewPaymentHandler(paymentRepo repository.PaymentRepository, MonthRepo repository.MonthRepository) PaymentHandler {
	return implPaymentHandler{
		PaymentRepository: paymentRepo,
		MonthRepository:   MonthRepo,
	}
}

type implPaymentHandler struct {
	PaymentRepository repository.PaymentRepository
	MonthRepository   repository.MonthRepository
}

type PaymentsForDisp struct {
	Month    *model.Month     `json:"month"`
	Payments *[]model.Payment `json:"payments"`
	TotalFee int              `json:"total_fee"`
}

func (self implPaymentHandler) GetPaymentList(c echo.Context) (err error) {
	t := time.Now()

	pfdl := make([]*PaymentsForDisp, 6)
	tf := 0
	for i, _ := range pfdl {
		display := t.AddDate(0, 1, 0).Format("2006-01")

		m, err := self.MonthRepository.MonthByDisplay(display)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
		}

		ps, err := self.PaymentRepository.PaymentListByMonthID(m.ID)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
		}
		for _, p := range *ps {
			tf += int(p.Amount.Int64)
		}

		pfdl[i] = &PaymentsForDisp{Month: m, Payments: ps, TotalFee: tf}
		// 一ヶ月前
		t = t.AddDate(0, -1, 0)
		tf = 0
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"paymentsList": pfdl,
	})
}

func (self implPaymentHandler) CreatePayment(c echo.Context) (err error) {
	mid, err := helper.Atoi64(c.FormValue("month_id"))
	ptid, err := helper.Atoi64(c.FormValue("payment_type_id"))
	psid, err := helper.Atoi64(c.FormValue("payment_status_id"))
	am, err := helper.Atoi64(c.FormValue("amount"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": "「しっかり じっくり 確実に入力しろ。 責任てのは、責任とれるやつが言う言葉だ。」"})
	}

	if mid == 0 || ptid == 0 || psid == 0 || am == 0 {
		return c.JSON(http.StatusOK, map[string]string{"error": "「どこかに0が入っているな。まさか返済できねえのか？ 死にてぇなら、生命保険、加入してからにしろ。返済がまだだぜ。」"})
	}

	p, err := self.PaymentRepository.PaymentByPaymentTypeIDMonthID(null.IntFrom(ptid), null.IntFrom(mid))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": "「そんな金俺は貸してねえ。 月 と 何に使ったのか もう一度確認しろ。」"})
	}
	p.PaymentStatusID = null.IntFrom(psid)
	p.Amount = null.IntFrom(am)

	err = self.PaymentRepository.Update(p)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": "...エラーだ。おい、宿本に連絡しろ。 ギャンブルにハマった奴の明日は信用しねェ。 そこで待っとけ。"})
	}

	return c.JSON(http.StatusOK, map[string]bool{"isSuccess": true})
}
