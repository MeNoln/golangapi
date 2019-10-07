package db

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib" //sql provider
	"github.com/jmoiron/sqlx"
)

//DB Variables conts
const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "1234"
	DB_NAME = "gorestdb"
)

var db *sqlx.DB

//InitializeDB ...
func InitializeDB() error {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	db, err := sqlx.Connect("pgx", dbinfo)
	if err != nil {
		log.Fatalln("Failed to connect to database")
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Failed to connect to database")
		return err
	}

	return nil
}

//GetDb ...
func GetDb() *sqlx.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	db, _ := sqlx.Connect("pgx", dbinfo)
	return db
}

//Close ...
func Close() {
	db.Close()
}
