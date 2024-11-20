package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver import
)

// ConnectDatabase initializes the database connection
func ConnectDatabase() (*sql.DB, error) {
	// Load DB config from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Check for missing environment variables
	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}

	// Construct DSN (Data Source Name)
	var dsn string
	if dbPass == "" {
		// No password provided
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)
	} else {
		// Password provided
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	}

	fmt.Println("dsn", dsn)
	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	// Test database connection
	if err := db.Ping(); err != nil {
		// Close the database connection if Ping fails
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Return the established DB connection
	return db, nil
}
