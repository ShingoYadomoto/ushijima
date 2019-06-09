package handler

import (
	"net/http"

	"strconv"

	"github.com/ShingoYadomoto/vue-go-heroku/server/context"
	"github.com/ShingoYadomoto/vue-go-heroku/server/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Home(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

func GetUser(c echo.Context) (err error) {
	db := c.(*context.CustomContext).GetDB()

	userID, err := strconv.Atoi(c.Param("userID"))

	u, err := model.UserByID(db, userID)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}
