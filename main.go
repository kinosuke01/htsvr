package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./"))
	mux.Handle("/", http.StripPrefix("/", fileServer))
	server := http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	server.ListenAndServe()
}
