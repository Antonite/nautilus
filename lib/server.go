package lib

type server struct {
	Ships []ship
}

func NewServer() *server {
	return &server{}
}
