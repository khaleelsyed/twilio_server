package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /voice", voice)
	mux.HandleFunc("GET /token", token)

	const addr string = ":4000"

	log.Print("Staring server on " + addr)

	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
