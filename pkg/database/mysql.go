package database

import (
	"database/sql"
	"desafio-tecnico-backend/internal/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Abertura e configuração do banco de dados MySQL
func MySQL(conf *config.Config) *database_pool {

	if db_pool != nil && db_pool.DB != nil {

		return db_pool
	} else {

		db, err := sql.Open(conf.DB_DRIVE, conf.DB_DSN)
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		db.SetMaxOpenConns(8)
		db.SetConnMaxIdleTime(6)
		db.SetConnMaxLifetime(5 * time.Minute)

		db_pool = &database_pool{
			DB: db,
		}
	}
	return db_pool
}
