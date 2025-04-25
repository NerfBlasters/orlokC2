package router

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) { 1 usage new *
	log.Printf(format:"You hit the endpoint: %s\n", r.URL.Path)
	
	w.Write([]byte("I'm Mister Derp!"))
}
