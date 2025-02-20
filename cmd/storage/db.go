package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var  db *sql.DB

func InitDB(){
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env files")
    }

    // dbHost := os.Getenv("DB_HOST")
    // dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")

    db, err := sql.Open("postgres", fmt.Sprintf(
        "user=%s password=%s dbname=%s sslmode=disable",
        dbUser,
        dbPass,
        dbName,
    ))

    if err != nil {
        panic(err.Error())
    }

    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
    return db
}


