package db

import (
	"github.com/ShingoYadomoto/ushijima/server/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(c *config.Pgsql) (*sqlx.DB, error) {
	dsn := "host=" + c.Host + " user=" + c.User + " dbname=" + c.DbName + " password=" + c.Password + " sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
