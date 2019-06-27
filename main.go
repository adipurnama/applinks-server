package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/.well-known/", appleHandler())
	http.ListenAndServe(":4000", nil)
}

func appleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileServer := http.FileServer(http.Dir("./static/well-known/"))
		w.Header().Set("content-type", "application/json")
		s := http.StripPrefix("/.well-known/", fileServer)
		s.ServeHTTP(w, r)
	}
}
