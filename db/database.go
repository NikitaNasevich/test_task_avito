package db

import (
	"errors"
	"github.com/NikitaNasevich/test_task_avito/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var db *sqlx.DB

func ConnectDatabase() (err error) {
	dataSourceName := helpers.GetEnvDefault("DATABASE", "")
	if dataSourceName == "" {
		return errors.New("no DATABASE env set")
	}
	if db, err = sqlx.Connect("mysql", dataSourceName); err != nil {
		return err
	}

	db.SetMaxOpenConns(helpers.GetEnvAsInt("DB_MAX_OPEN_CONNECTIONS", 50))

	db.SetMaxIdleConns(helpers.GetEnvAsInt("DB_MAX_IDLE_CONNECTIONS", 50))

	db.SetConnMaxLifetime(time.Minute * time.Duration(helpers.GetEnvAsInt("DB_CONNECTION_MAX_LIFETIME", 10)))

	db.MapperFunc(func(column string) string {
		return column
	})

	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func Database() *sqlx.DB {
	return db
}
