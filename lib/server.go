package lib

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
