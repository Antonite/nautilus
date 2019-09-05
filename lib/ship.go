package lib

import (
	"errors"
)

// ship struct
// Name - the name of the ship
// DataRecords - list of data records of the ship
type ship struct {
	Name        string
	DataRecords []dataRecord
}

// helper type to organize unit conversions
type timeMultiplier uint

const (
	SecondsPerHour   timeMultiplier = 3600
	SecondsPerMinute timeMultiplier = 60
)

// Create a new ship object
func NewShip(name string) *ship {
	return &ship{
		Name: name,
	}
}

// Given a path to a local file, load data records for a given ship.
// Returns any encountered errors.
func (ship *ship) loadDataFromFile(path string) error {
	records, err := parseDataRecordsFromCSV(path)
	if err != nil {
		return err
	}

	// fill in any missing points
	correctDataPoints(records, FieldsToCorrect)

	ship.DataRecords = append(ship.DataRecords, records...)

	return nil
}

// Sum the values of data records that fall in a given time frame for a ship.
// Takes a dataField that determines which type of data point to sum.
// Takes a multiplier to convert between units of measurement of the underlying data point.
// Returns any encountered errors.
func (ship *ship) sumDataBetweenTimeFrames(startTime float64, endTime float64, field dataField, multiplier timeMultiplier) (float64, error) {
	if startTime < 0 || endTime < 0 || endTime < startTime {
		return -1, errors.New("Invalid start/end times")
	}

	sum := 0.00
	firstVal := true
	multi := float64(multiplier)

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

			sum += ship.DataRecords[i-1].dataMap[field] * (endTime - prevTime) / multi
			break
		}

		// special case for first difference between two time frames
		if firstVal {
			sum += ship.DataRecords[i-1].dataMap[field] * (currTime - startTime) / multi
			firstVal = false
			continue
		}

		sum += ship.DataRecords[i-1].dataMap[field] * (currTime - ship.DataRecords[i-1].dataMap[TimestampField]) / multi
	}

	return sum, nil
}

func (ship *ship) getEfficiencyBetweenTimeFrames(start float64, end float64) (float64, error) {
	distance, errD := ship.sumDataBetweenTimeFrames(start, end, SpeedField, SecondsPerHour)
	if errD != nil {
		return -1, errD
	}

	fuel, errF := ship.sumDataBetweenTimeFrames(start, end, FuelField, SecondsPerMinute)
	if errF != nil {
		return -1, errF
	}

	// can't divide by zero
	if fuel == 0 {
		return -1, errors.New("No fuel was spent.")
	}

	efficiency := distance / fuel
	return efficiency, nil
}
