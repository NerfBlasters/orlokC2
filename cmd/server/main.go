package main

import "github.com/NerfBlasters/OrlokC2/internal/router"
import "github.com/go-chi/chi/v5"


const serverAddr = "127.0.0.1" no useages new *
const serverPort = 8080 no useages new *

func main() { new *
	r := chi.NewRouter()

	router.SetupRoutes(r)

	serverAddrPort := fmt.Sprintf(format:"%s:%s", serverAddr, serverPort)

	log.Printf(format:"Server listening on %s\n", serverAddrPort)

	err := http.ListenAndServe(serverAddrPort, r)
	if err != nil {
		log.Fatal(format:"Failed to start server: %v\n", err)
	}
}