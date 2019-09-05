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

func (ship *ship) sumDataBetweenTimeFrames(startTime float64, endTime float64, field dataField, multiplier float64) float64 {
	// ideally, we would return an error here, but requirements state otherwise.
	if startTime < 0 || endTime < 0 || endTime < startTime {
		return -1
	}

	sum := 0.00
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

			sum += ship.DataRecords[i-1].dataMap[field] * (endTime - prevTime) / multiplier
			break
		}

		// special case for first difference between two time frames
		if firstVal {
			sum += ship.DataRecords[i-1].dataMap[field] * (currTime - startTime) / multiplier
			firstVal = false
			continue
		}

		sum += ship.DataRecords[i-1].dataMap[field] * (currTime - ship.DataRecords[i-1].dataMap[TimestampField]) / multiplier
	}

	return sum
}
