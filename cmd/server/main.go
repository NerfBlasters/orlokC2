package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//import "github.com/NerfBlasters/OrlokC2/internal/router"

const serverAddr = "127.0.0.1"
const serverPort = 8080

func main() {
	r := chi.NewRouter()

	router.SetupRoutes(r)

	serverAddrPort := fmt.Sprintf("%s:%s", serverAddr, serverPort)

	log.Printf("Server listening on %s\n", serverAddrPort)

	err := http.ListenAndServe(serverAddrPort, r)
	if err != nil {
		log.Fatal("Failed to start server: %v\n", err)
	}
}
