package middleware

import (
	"github.com/ShingoYadomoto/vue-go-heroku/server/context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func DBMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(context.DBKey, db)
			log.Debug("set db to echo.Context.")
			err := next(c)
			if err != nil {
				c.Error(err)
			}
			return err
		}
	}
}
