package handler

import (
	"net/http"

	"github.com/ShingoYadomoto/vue-go-heroku/server/context"
	"github.com/ShingoYadomoto/vue-go-heroku/server/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

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

func GetAllPaymentStatuses(c echo.Context) (err error) {
	db := c.(*context.CustomContext).GetDB()

	ptl, err := model.AllPaymentStatuses(db)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"paymentStatusList": ptl,
	})
}

func GetAllMonths(c echo.Context) (err error) {
	db := c.(*context.CustomContext).GetDB()

	ml, err := model.AllMonths(db)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"monthList": ml,
	})
}
