package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *pgxpool.Pool

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	initPGConnection()
}

func initPGConnection() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, database, password)
	conn, err := pgxpool.Connect(context.Background(), uri)
	if err != nil {
		log.Fatal(err)
	}

	db = conn
}

func GetDB() *pgxpool.Pool {
	return db
}
