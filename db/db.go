package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() error {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Println("error while starting the psql server: ", err)
		return err
	}
	if err = DB.Ping(); err != nil {
		log.Println("error making a test ping to the server: ", err)
		return err
	}
	log.Println("Database connected successfully")
	return nil
}
