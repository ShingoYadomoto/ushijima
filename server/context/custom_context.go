package context

import (
	"github.com/ShingoYadomoto/vue-go-heroku/server/config"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func CustomContextMiddleware() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	}
}

type CustomContext struct {
	echo.Context
}

const (
	ConfigKey = "__CONFIG__"
	DBKey     = "__DB__"
)

func (c *CustomContext) GetConfig() *config.Conf {
	conf, ok := c.Get(ConfigKey).(*config.Conf)
	if !ok {
		log.Panic("*config.Conf assertion error")
	}
	return conf
}

func (c *CustomContext) GetDB() *sqlx.DB {
	db, ok := c.Get(DBKey).(*sqlx.DB)
	if !ok {
		log.Panic("*sqlx.DB assertion error")
	}
	return db
}
