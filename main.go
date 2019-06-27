package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/apple-app-site-association", appleHandler())
	http.HandleFunc("/.well-known/", wellknownHandler())

	port := 4000

	log.Printf("Serving http at port %d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

func appleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileServer := http.FileServer(http.Dir("./static/well-known/"))
		w.Header().Set("content-type", "application/json")
		fileServer.ServeHTTP(w, r)
	}
}

func wellknownHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileServer := http.FileServer(http.Dir("./static/well-known/"))
		w.Header().Set("content-type", "application/json")
		s := http.StripPrefix("/.well-known/", fileServer)
		s.ServeHTTP(w, r)
	}
}
