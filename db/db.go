package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	connStr := "user=fadhadwahyuaji dbname=db_karyawan_sangkuriang sslmode=disable password=12345678"
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
}
