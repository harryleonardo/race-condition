package database

import (
	"database/sql"
	"fmt"
)

type (
	PostgresInterface interface {
		OpenPostgresConn() *sql.DB
	}
)

func (d *database) OpenPostgresConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.SharedConfig.POSTGRES.Host,
		d.SharedConfig.POSTGRES.Port,
		d.SharedConfig.POSTGRES.Username,
		d.SharedConfig.POSTGRES.Password,
		d.SharedConfig.POSTGRES.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
