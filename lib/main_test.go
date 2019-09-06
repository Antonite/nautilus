package lib

import (
	"os"
	"testing"
)

var testShip *ship

// Main test func setup
// Add dummy data to test ship
func TestMain(m *testing.M) {
	testShip = NewShip("test")

	aMap1 := make(map[dataField]float64)
	aMap1[TimestampField] = 1561546800
	aMap1[SpeedField] = 18
	aMap1[FuelField] = 8
	record1 := newDataRecord(aMap1)

	aMap2 := make(map[dataField]float64)
	aMap2[TimestampField] = 1561550400
	aMap2[SpeedField] = 18
	aMap2[FuelField] = 4
	record2 := newDataRecord(aMap2)

	aMap3 := make(map[dataField]float64)
	aMap3[TimestampField] = 1561579200
	aMap3[SpeedField] = 18
	aMap3[FuelField] = 0
	record3 := newDataRecord(aMap3)

	aMap4 := make(map[dataField]float64)
	aMap4[TimestampField] = 1561582800
	aMap4[SpeedField] = 18
	aMap4[FuelField] = 6
	record4 := newDataRecord(aMap4)

	records := []dataRecord{*record1, *record2, *record3, *record4}

	testShip.DataRecords = records

	os.Exit(m.Run())
}
