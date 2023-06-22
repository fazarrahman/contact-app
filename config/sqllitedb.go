package config

import (
	"log"

	"github.com/fazarrahman/contact-app/lib"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func New() (*sqlx.DB, error) {
	dbname := lib.GetEnv("DB_NAME")
	db, err := sqlx.Connect("sqlite3", dbname+".db")
	if err != nil {
		panic(err)
	}

	// Create a table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS contacts (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		gender varchar(1) NOT NULL,
		phone VARCHAR(255) NULL,
		email VARCHAR(255) NULL,
		created_at timestamp NOT NULL,
		updated_at timestamp NULL
	);`)
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}
