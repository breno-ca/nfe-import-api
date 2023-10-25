package database

import (
	"database/sql"
	"desafio-tecnico-backend/internal/config"
	"fmt"
)

type DatabaseInterface interface {
	GetDB() (DB *sql.DB)
	Close() error
}

type database_pool struct {
	DB *sql.DB
}

var db_pool = &database_pool{}

func NewBD(conf *config.Config) *database_pool {
	// Atribuição do Data Source Name
	conf.DBConfig.DB_DSN = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.DBConfig.DB_USER, conf.DBConfig.DB_PASS, conf.DBConfig.DB_HOST, conf.DBConfig.DB_PORT, conf.DBConfig.DB_NAME)

	db_pool = MySQL(conf)

	return db_pool
}

func (d *database_pool) GetDB() (DB *sql.DB) {
	return d.DB
}

func (d *database_pool) Close() error {
	err := d.DB.Close()
	if err != nil {
		return err
	}

	db_pool = &database_pool{}

	return err
}
