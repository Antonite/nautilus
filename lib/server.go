package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		server.getDistanceHandler(w, r)
	})

	// http.HandleFunc("/total_fuel", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })

	// http.HandleFunc("/efficiency", func(w http.ResponseWriter, r *http.Request) {
	// 	server.OrdersGetHandler(w, r)
	// })
}

func (server *Server) getDistanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type view struct {
		TotalDistance float64 `json:"total_distance"`
	}

	// will return -1 in case any errors occur.
	// Ideally we would return an error here, but requirements state otherwise.
	aView := view{TotalDistance: -1.00}

	startRaw := r.URL.Query().Get("start")
	endRaw := r.URL.Query().Get("end")

	start, startErr := strconv.ParseFloat(startRaw, 64)
	end, endErr := strconv.ParseFloat(endRaw, 64)
	if startErr != nil || endErr != nil {
		log.Println("Failed to parse url parameters.")
		js, err := json.Marshal(aView)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(js)
		return
	}

	// hack, the api should provide ship id, but requirements state otherwise.
	aView.TotalDistance = server.Ships[0].getDistanceTraveled(start, end)

	js, err := json.Marshal(aView)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(js)
}
