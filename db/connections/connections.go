package connections

import (
	"database/sql"
	"fmt"
	"main/infra/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func OpenMySQLConnection() *sql.DB {
	USER := viper.GetString("MASTER_DB_USER")
	PASSWORD := viper.GetString("MASTER_DB_PASSWORD")
	HOST := viper.GetString("MASTER_DB_HOST")
	PORT := viper.GetString("MASTER_DB_PORT")
	DB_NAME := viper.GetString("MASTER_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASSWORD, HOST, PORT, DB_NAME)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Fatalf("failed to open database connection: %v", err)
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		logger.Fatalf("Failed to ping database connection: %v", err)
	}

	logger.Infof("Successfully connected to the database!")

	return db
}
