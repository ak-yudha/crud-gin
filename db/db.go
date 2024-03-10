package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// SetupDB initializes the database connection
func SetupDB() (*sql.DB, error) {
	// Replace with your actual MySQL database connection details
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/your_database_name")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to the database")
	return db, nil
}
