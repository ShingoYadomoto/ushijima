package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type HomeHandler interface {
	Home(c echo.Context) (err error)
}

func NewHomeHandler() HomeHandler {
	return implHomeHandler{}
}

type implHomeHandler struct{}

func (self implHomeHandler) Home(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
