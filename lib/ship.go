package lib

type ship struct {
	Name        string
	DataRecords []dataRecord
}

func NewShip(name string) *ship {
	return &ship{
		Name: name,
	}
}

func (ship *ship) loadDataFromFile(path string) error {
	records, err := parseDataRecordsFromCSV(path)
	if err != nil {
		return err
	}

	correctDataPoints(records, FieldsToCorrect)
	ship.DataRecords = append(ship.DataRecords, records...)

	return nil
}
