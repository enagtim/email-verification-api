package db

import (
	"api/email-verification/configs"
	"database/sql"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := sql.Open("postgres", conf.Db.Dsn)
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
