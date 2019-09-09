package handler

import (
	"net/http"

	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type HomeHandler interface {
	Home(c echo.Context) (err error)
	GetAllPaymentTypes(c echo.Context) (err error)
	GetAllPaymentStatuses(c echo.Context) (err error)
	GetAllMonths(c echo.Context) (err error)
}

func NewHomeHandler(paymentTypeRepo repository.PaymentTypeRepository, paymentStatusRepo repository.PaymentStatusRepository, MonthRepo repository.MonthRepository) HomeHandler {
	return implHomeHandler{
		PaymentTypeRepository:   paymentTypeRepo,
		PaymentStatusRepository: paymentStatusRepo,
		MonthRepository:         MonthRepo,
	}
}

type implHomeHandler struct {
	PaymentTypeRepository   repository.PaymentTypeRepository
	PaymentStatusRepository repository.PaymentStatusRepository
	MonthRepository         repository.MonthRepository
}

func (self implHomeHandler) Home(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

func (self implHomeHandler) GetAllPaymentTypes(c echo.Context) (err error) {
	ptl, err := self.PaymentTypeRepository.AllPaymentTypes()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"paymentTypeList": ptl,
	})
}

func (self implHomeHandler) GetAllPaymentStatuses(c echo.Context) (err error) {
	ptl, err := self.PaymentStatusRepository.AllPaymentStatuses()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"paymentStatusList": ptl,
	})
}

func (self implHomeHandler) GetAllMonths(c echo.Context) (err error) {
	ml, err := self.MonthRepository.AllMonths()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"monthList": ml,
	})
}
