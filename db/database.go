package db

import (
	"errors"
	"github.com/NikitaNasevich/test_task_avito/helpers"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func ConnectDatabase(err error) {
	dataSourceName := helpers.GetEnvDefault("DATABASE", "")
	if dataSourceName == "" {
		return errors.New("no DATABASE env set")
	}
	db, err = sqlx.Connect("postgres", dataSourceName)
}
