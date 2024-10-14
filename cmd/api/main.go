// cmd/api/main.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "example.com/mon-api/routes"
)

func main() {
    router := routes.InitRoutes()

    fmt.Println("Serveur démarré sur le port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
