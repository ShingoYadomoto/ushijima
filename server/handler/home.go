package handler

import (
	"net/http"

	"time"

	"github.com/ShingoYadomoto/vue-go-heroku/server/context"
	"github.com/ShingoYadomoto/vue-go-heroku/server/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type PaymentsForDisp struct {
	Month    *model.Month     `json:"month"`
	Payments *[]model.Payment `json:"payments"`
	TotalFee int              `json:"total_fee"`
}

func Home(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

func GetAllPaymentTypes(c echo.Context) (err error) {
	db := c.(*context.CustomContext).GetDB()

	ptl, err := model.AllPaymentTypes(db)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"paymentTypeList": ptl,
	})
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
