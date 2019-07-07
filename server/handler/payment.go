package handler

import (
	"net/http"
	"time"

	"github.com/ShingoYadomoto/vue-go-heroku/server/context"
	"github.com/ShingoYadomoto/vue-go-heroku/server/model"

	"github.com/ShingoYadomoto/vue-go-heroku/server/helper"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/guregu/null.v3"
)

type PaymentsForDisp struct {
	Month    *model.Month     `json:"month"`
	Payments *[]model.Payment `json:"payments"`
	TotalFee int              `json:"total_fee"`
}

func GetPaymentList(c echo.Context) (err error) {
	db := c.(*context.CustomContext).GetDB()

	t := time.Now()

	pfdl := make([]*PaymentsForDisp, 6)
	tf := 0
	for i, _ := range pfdl {
		display := t.AddDate(0, 1, 0).Format("2006-01")

		m, err := model.MonthByDisplay(db, display)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
		}

		ps, err := model.PaymentListByMonthID(db, m.ID)
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

func CreatePayment(c echo.Context) (err error) {
	db := c.(*context.CustomContext).GetDB()

	mid, err := helper.Atoi64(c.FormValue("month_id"))
	ptid, err := helper.Atoi64(c.FormValue("payment_type_id"))
	psid, err := helper.Atoi64(c.FormValue("payment_status_id"))
	am, err := helper.Atoi64(c.FormValue("amount"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": "送信内容がなんかおかしいよ"})
	}

	if mid == 0 || ptid == 0 || psid == 0 || am == 0 {
		return c.JSON(http.StatusOK, map[string]string{"error": "何か入力してないか、0入っちゃってるよ"})
	}

	p, err := model.PaymentByPaymentTypeIDMonthID(db, null.IntFrom(ptid), null.IntFrom(mid))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": "そんなtypeとstatusの支払いみつからないんだけど"})
	}
	p.PaymentStatusID = null.IntFrom(psid)
	p.Amount = null.IntFrom(am)

	err = p.Update(db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": "データベース更新できなかったすまん"})
	}

	return c.JSON(http.StatusOK, map[string]bool{"isSuccess": true})
}
