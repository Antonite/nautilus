package lib

import (
	"log"
	"net/http"
)

type Server struct {
	Mux   *http.ServeMux
	Ships []ship
}

func NewServer() *Server {
	server := &Server{
		Mux: http.NewServeMux(),
	}

	server.registerRoutes()
	return server
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

func (server *Server) registerRoutes() {
	server.Mux.HandleFunc("/total_distance", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit endpoint: /total_distance")
		server.GetDistanceHandler(w, r)
	})

	// http.HandleFunc("/total_fuel", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })

	// http.HandleFunc("/efficiency", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })
}

func (server *Server) GetDistanceHandler(w http.ResponseWriter, r *http.Request) {

	// js, err := json.Marshal(views)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)
}
