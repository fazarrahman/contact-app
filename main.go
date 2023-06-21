package main

import (
	sqllitedb "github.com/fazarrahman/contact-app/config/sqllitedb"
)

func main() {
	_, err := sqllitedb.New()
	if err != nil {
		panic(err)
	}
}
