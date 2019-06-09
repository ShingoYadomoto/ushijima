package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func Home(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{"status":"OK"})
}
