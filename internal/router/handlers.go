package router

import (
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("You hit the endpoint: %s\n", r.URL.Path)

	w.Write([]byte("I'm Mister Derp!"))
}
