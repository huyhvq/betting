package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Connector interface {
	Open() (*sql.DB, error)
	Close() error
}

type db struct {
	DB *sql.DB
}

func (d *db) Open() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		viper.Get("db_user"),
		viper.Get("db_pass"),
		viper.Get("db_host"),
		viper.Get("db_port"),
		viper.Get("db_name"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	d.DB = db
	return db, nil
}

func (d *db) Close() error {
	return d.DB.Close()
}

func NewDB() Connector {
	return &db{}
}
