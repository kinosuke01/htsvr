package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./"))
	mux.Handle("/", http.StripPrefix("/", logInfo(fileServer)))
	server := http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	server.ListenAndServe()
}

func logInfo(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rAddr := r.RemoteAddr
		method := r.Method
		path := r.URL.Path
		fmt.Printf("Remote: %s [%s] %s\n", rAddr, method, path)
		h.ServeHTTP(w, r)
	})
}
