package main

import (
	"fmt"

	"github.com/ShingoYadomoto/vue-go-heroku/server/config"
	"github.com/ShingoYadomoto/vue-go-heroku/server/context"
	"github.com/ShingoYadomoto/vue-go-heroku/server/db"
	"github.com/ShingoYadomoto/vue-go-heroku/server/handler"
	"github.com/ShingoYadomoto/vue-go-heroku/server/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	conf := config.GetConfig()

	db, err := db.NewDB(conf.Pgsql)
	if err != nil {
		log.Panic(fmt.Errorf("Faild to connect DB. %v", err))
	}
	defer func() {
		err := recover()
		if err != nil {
			db.Close()
			log.Panic(fmt.Errorf("Faild to prepare echo. %v", err))
		}
	}()

	e := initEcho(&conf, db)

	e.Debug = true

	e.GET("/", handler.Home)
	e.GET("/user/:userID", handler.GetUser)

	// Start server
	address := ":" + conf.App.Port
	e.Logger.Fatal(e.Start(address))
}

func initEcho(conf *config.Conf, db *sqlx.DB) *echo.Echo {
	// Setup
	e := echo.New()

	e.Logger.SetLevel(conf.Log.Level)
	log.SetLevel(conf.Log.Level)

	e.Use(context.CustomContextMiddleware())
	e.Use(middleware.BasicAuthMiddleware())
	e.Use(middleware.ConfigMiddleware(conf))
	e.Use(middleware.DBMiddleware(db))
	e.Use(echo_middleware.Logger())
	e.Use(echo_middleware.Recover())

	return e
}
