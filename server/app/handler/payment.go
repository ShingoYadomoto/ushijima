package handler

import (
	"net/http"

	"github.com/ShingoYadomoto/ushijima/server/app/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type PaymentHandler interface {
	GetPaymentList(c echo.Context) (err error)
	CreatePayment(c echo.Context) (err error)
}

func NewPaymentHandler(paymentUse usecase.PaymentUsecase) PaymentHandler {
	return implPaymentHandler{
		PaymentUse: paymentUse,
	}
}

type implPaymentHandler struct {
	PaymentUse usecase.PaymentUsecase
}

func (self implPaymentHandler) GetPaymentList(c echo.Context) (err error) {
	pfdl, err := self.PaymentUse.GetPaymentList()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"paymentsList": pfdl})
}

// なんかエラーハンドリングが変
func (self implPaymentHandler) CreatePayment(c echo.Context) (err error) {
	type request struct {
		PaymentTypeID   int `form:"payment_type_id"`
		PaymentStatusID int `form:"payment_status_id"`
		MonthID         int `form:"month_id"`
		Amount          int `form:"amount"`
	}

	req := &request{}
	if err = c.Bind(req); err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": "「しっかり じっくり 確実に入力しろ。 責任てのは、責任とれるやつが言う言葉だ。」"})
	}

	err = self.PaymentUse.CreatePayment(req.PaymentTypeID, req.PaymentStatusID, req.MonthID, req.Amount)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]bool{"isSuccess": true})
}
