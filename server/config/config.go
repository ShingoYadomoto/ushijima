package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

type Conf struct {
	App     *App     `envconfig:"app"`
	Log     *Log     `envconfig:"log"`
	Pgsql   *Pgsql   `envconfig:"pgsql"`
	LineApi *LineApi `envconfig:"lineapi"`
}

type App struct {
	Name   string
	Domain string
	URL    string
	Port   string
}

type Log struct {
	Level log.Lvl
}

type Pgsql struct {
	User     string
	Password string
	Host     string
	DbName   string `split_words:"true"`
}

type LineApi struct {
	ChannelSecret      string
	ChannelAccessToken string
}

func GetConfig() (conf Conf) {
	err := godotenv.Load()
	if err != nil {
		err = errors.Wrap(err, "Error loading .env file")
		log.Error(err)
	}

	err = envconfig.Process("ushijima", &conf)
	if err != nil {
		err = errors.Wrap(err, "Error mapping .env file")
		log.Error(err)
	}

	return
}
