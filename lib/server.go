package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Server struct
// Mux - custom http muxer
// Ships - list of ships in the system
type Server struct {
	Mux   *http.ServeMux
	Ships []ship
}

// Create new server.
// Registers http routes to the internal muxer as part of initialization.
// Returns pointer to the newly created server.
func NewServer() *Server {
	server := &Server{
		Mux: http.NewServeMux(),
	}

	server.registerRoutes()
	return server
}

// Given a path to a local file, initialize the server data.
// Returns any encountered errors.
func (server *Server) InitFromFile(path string) error {
	// load ship data
	aShip := NewShip(path)
	if err := aShip.loadDataFromFile(path); err != nil {
		return err
	}

	server.Ships = append(server.Ships, *aShip)
	return nil
}

// Register routes to the given server's muxer.
func (server *Server) registerRoutes() {
	server.Mux.HandleFunc("/total_distance", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit endpoint: /total_distance")
		server.getDistanceHandler(w, r)
	})

	server.Mux.HandleFunc("/total_fuel", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit endpoint: /total_fuel")
		server.getFuelHandler(w, r)
	})

	server.Mux.HandleFunc("/efficiency", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit endpoint: /efficiency")
		server.getEfficiencyHandler(w, r)
	})
}

// GET distance http handler
// Accepts start and end time as parameters in epoch format.
// Computes the total distance in miles within the given timeframe.
// Writes the distance as JSON to the ResponseWriter.
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

	// parse times
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
	var err error
	aView.TotalDistance, err = server.Ships[0].sumDataBetweenTimeFrames(start, end, SpeedField, SecondsPerHour)
	if err != nil {
		// Ideally we would return an error here, but requirements state otherwise.
		log.Println(err)
	}

	js, err := json.Marshal(aView)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(js)
}

// GET fuel http handler
// Accepts start and end time as parameters in epoch format.
// Computes the total fuel spent in gallons within the given timeframe.
// Writes the fuel amount as JSON to the ResponseWriter.
func (server *Server) getFuelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type view struct {
		TotalFuel float64 `json:"total_fuel"`
	}

	// will return -1 in case any errors occur.
	// Ideally we would return an error here, but requirements state otherwise.
	aView := view{TotalFuel: -1.00}

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
	var err error
	aView.TotalFuel, err = server.Ships[0].sumDataBetweenTimeFrames(start, end, FuelField, SecondsPerMinute)
	if err != nil {
		// Ideally we would return an error here, but requirements state otherwise.
		log.Println(err)
	}

	js, err := json.Marshal(aView)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(js)
}

// GET efficiency http handler
// Accepts start and end time as parameters in epoch format.
// Computes the miles per gallon of fuel efficiency within the given timeframe.
// Writes the efficiency as JSON to the ResponseWriter.
func (server *Server) getEfficiencyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type view struct {
		Efficiency float64 `json:"efficiency"`
	}

	// will return -1 in case any errors occur.
	// Ideally we would return an error here, but requirements state otherwise.
	aView := view{Efficiency: -1.00}

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
	var err error
	aView.Efficiency, err = server.Ships[0].getEfficiencyBetweenTimeFrames(start, end)
	if err != nil {
		// Ideally we would return an error here, but requirements state otherwise.
		log.Println(err)
	}

	js, err := json.Marshal(aView)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(js)
}
