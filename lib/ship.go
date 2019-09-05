package lib

type ship struct {
	Name        string
	DataRecords []dataRecord
}

func NewShip(name string) *ship {
	return &ship{Name: name}
}
