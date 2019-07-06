package main

import (
	"fmt"

	"net/http"

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
	e.GET("/month", handler.GetAllMonths)
	e.GET("/payment_type", handler.GetAllPaymentTypes)
	e.GET("/payment_status", handler.GetAllPaymentStatuses)

	e.GET("/payment", handler.GetPaymentList)
	e.POST("/payment/create", handler.CreatePayment)

	// Start server
	address := ":" + conf.App.Port
	e.Logger.Fatal(e.Start(address))
}

func initEcho(conf *config.Conf, db *sqlx.DB) *echo.Echo {
	// Setup
	e := echo.New()

	e.Logger.SetLevel(conf.Log.Level)
	log.SetLevel(conf.Log.Level)

	e.Use(echo_middleware.CORSWithConfig(echo_middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"origin", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}))
	e.Use(context.CustomContextMiddleware())
	e.Use(middleware.BasicAuthMiddleware())
	e.Use(middleware.ConfigMiddleware(conf))
	e.Use(middleware.DBMiddleware(db))
	e.Use(echo_middleware.Logger())
	e.Use(echo_middleware.Recover())

	return e
}
