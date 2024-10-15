package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erreur lors de la connexion à la base de données :", err)
	}

	fmt.Println("Connexion réussie à la base de données PostgreSQL")
}
