package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mon-api/db"

	routes "mon-api/Routes"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	db.InitDB()

	router := routes.InitRoutes()
	port := os.Getenv("PORT")
	fmt.Println("Serveur démarré sur le port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
