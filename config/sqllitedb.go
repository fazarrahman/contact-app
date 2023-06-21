package sqllitedb

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func New() (*sqlx.DB, error) {
	dbname := os.Getenv("DB_NAME")
	db, err := sqlx.Connect("sqlite3", dbname+".db")
	if err != nil {
		panic(err)
	}

	// Create a table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS  contacts (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		gender varchar(1) NOT NULL,
		phone VARCHAR(255) NULL,
		email VARCHAR(255) NULL,
		created_at timestamp NOT NULL,
		updated_at timestamp NULL,
		UNIQUE(phone)
	);`)
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}
