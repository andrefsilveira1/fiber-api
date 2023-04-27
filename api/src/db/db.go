package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 "172.17.0.2:3306",
		DBName:               "portfolio",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if pingErr := db.Ping(); pingErr != nil {
		db.Close()
		log.Fatal(pingErr)
		return nil, pingErr
	}
	fmt.Println("Connected!")

	return db, nil
}
