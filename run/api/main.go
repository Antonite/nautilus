package main

import (
	"flag"
	"log"
	"nautilus/lib"
	"net/http"
)

func main() {
	inputFlag := flag.String("i", "res/ship.csv", "data file")
	flag.Parse()

	server := lib.NewServer()
	if err := server.InitFromFile(*inputFlag); err != nil {
		log.Fatal(err)
	}

	registerRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerRoutes() {
	http.HandleFunc("/total_distance", func(w http.ResponseWriter, r *http.Request) {
		lib.NewServer()
	})

	// http.HandleFunc("/total_fuel", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })

	// http.HandleFunc("/efficiency", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })
}
