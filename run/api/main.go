package main

import (
	"flag"
	"log"
	"nautilus/lib"
	"net/http"
)

func main() {
	inputFlag := flag.String("f", "res/ship.csv", "data file")
	flag.Parse()

	log.Println("Starting server...")
	server := lib.NewServer()
	if err := server.InitFromFile(*inputFlag); err != nil {
		log.Fatal(err)
	}

	log.Println("Server started.")
	log.Fatal(http.ListenAndServe(":8080", server.Mux))
}
