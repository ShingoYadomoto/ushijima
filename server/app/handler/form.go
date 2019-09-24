package handler

import (
	"net/http"

	"github.com/ShingoYadomoto/ushijima/server/app/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type FormHandler interface {
	GetAllPaymentTypes(c echo.Context) (err error)
	GetAllPaymentStatuses(c echo.Context) (err error)
	GetAllMonths(c echo.Context) (err error)
}

func NewFormHandler(formUse usecase.FormUsecase) FormHandler {
	return implFormHandler{
		FormUsecase: formUse,
	}
}

type implFormHandler struct {
	FormUsecase usecase.FormUsecase
}

func (self implFormHandler) Home(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

func (self implFormHandler) GetAllPaymentTypes(c echo.Context) (err error) {
	ptl, err := self.FormUsecase.GetAllPaymentTypes()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"paymentTypeList": ptl,
	})
}

func (self implFormHandler) GetAllPaymentStatuses(c echo.Context) (err error) {
	psl, err := self.FormUsecase.GetAllPaymentStatuses()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"paymentStatusList": psl,
	})
}

func (self implFormHandler) GetAllMonths(c echo.Context) (err error) {
	ml, err := self.FormUsecase.GetAllMonths()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"monthList": ml,
	})
}
