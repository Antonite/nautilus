package lib

import (
	"log"
	"net/http"
)

type Server struct {
	Ships []ship
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) InitFromFile(path string) error {
	// load ship data
	aShip := NewShip(path)
	if err := aShip.loadDataFromFile(path); err != nil {
		return err
	}

	server.Ships = append(server.Ships, *aShip)
	return nil
}

func (server *Server) RegisterRoutes() {
	http.HandleFunc("/total_distance", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hi")
		// lib.NewServer()
	})

	// http.HandleFunc("/total_fuel", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })

	// http.HandleFunc("/efficiency", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })
}
