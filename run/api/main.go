package main

import (
	"log"
	"nautilus/lib"
	"net/http"
)

func main() {

	server := lib.NewServer()
	nautilus := lib.NewShip("Nautilus")

	server.Ships = append(server.Ships, *nautilus)

	http.HandleFunc("/total_distance", func(w http.ResponseWriter, r *http.Request) {
		lib.NewServer()
	})

	// http.HandleFunc("/total_fuel", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })

	// http.HandleFunc("/efficiency", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })

	log.Fatal(http.ListenAndServe(":8080", nil))
}
