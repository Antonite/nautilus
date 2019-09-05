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

func (ship *ship) getDistanceTraveled(startTime float64, endTime float64) float64 {
	// ideally, we would return an error here, but requirements state otherwise.
	if startTime < 0 || endTime < 0 || endTime < startTime {
		return -1
	}

	distance := 0.00
	firstVal := true
	for i := 0; i < len(ship.DataRecords); i++ {
		currTime := ship.DataRecords[i].dataMap[TimestampField]

		// while before
		if currTime <= startTime || i == 0 {
			continue
		}

		// when after
		if currTime >= endTime {
			// add the last fragment between the ending time frames
			prevTime := ship.DataRecords[i-1].dataMap[TimestampField]
			if firstVal {
				prevTime = startTime
			}

			distance += ship.DataRecords[i-1].dataMap[SpeedField] * (endTime - prevTime) / 3600
			break
		}

		// special case for first distance between two time frames
		if firstVal {
			distance += ship.DataRecords[i-1].dataMap[SpeedField] * (currTime - startTime) / 3600
			firstVal = false
			continue
		}

		// log.Printf("%v, %v, %v\n", distance, ship.DataRecords[i-1].dataMap[SpeedField], currTime-ship.DataRecords[i-1].dataMap[TimestampField])
		distance += ship.DataRecords[i-1].dataMap[SpeedField] * (currTime - ship.DataRecords[i-1].dataMap[TimestampField]) / 3600
	}

	return distance
}
