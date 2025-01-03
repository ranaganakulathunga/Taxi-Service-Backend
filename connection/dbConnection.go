package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB = Connection()

func Connection() *sql.DB {
	// DSN with no password and root as the default username
	dsn := "root:@tcp(127.0.0.1:3306)/db_taxi"

	// Open the database connection
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	//defer database.Close()

	// Test the connection
	if err := database.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	_, connectionBeginError := database.Begin()

	if connectionBeginError != nil {
		fmt.Println("connecting begin error : %v", connectionBeginError)
	}
	result, _ := database.Exec("SELECT * FROM TEST")

	fmt.Println("Connected to MySQL database!", result)

	return database
}
