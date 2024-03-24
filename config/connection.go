package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// This variable use to hold database connection and will be use accross all handlers
var DB *sql.DB

// Function to connect with database
func ConnectDB() {
	// Load .env variable
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error is occured on load .env file")
	}

	// assign the .env variable value 
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")

	// For local test with database run on local set the sslmode=disable
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=require",
       host, port, user, dbname, pass)
	 // Connect to Postgresql database
   db, errSql := sql.Open("postgres", psqlSetup)
   if errSql != nil {
      fmt.Println("There is an error while connecting to the database ", err)
      panic(err)
   } else {
		  // Assign the success database connection into DB variable
      DB = db
      fmt.Println("Successfully connected to database!")
   }
	
}