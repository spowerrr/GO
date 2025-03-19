package config

//Database Connection
import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// global variable tohold  the database connection
var DB *sql.DB

// ConnectDB initializes a connection to postgreSql
func ConnectDB() {
	//load environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env files")
	}
	// Build PostgreSQL connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	//open the database connection
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("Failed to open database: ", err)
		return
	}

	//test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
		return
	}

	fmt.Println("✅ Connect to the database!")

}
