package lib

import (
	"testing"
)

// Tester function for parseDataRecordsFromCSV
func TestParseDataRecordsFromCSV(t *testing.T) {
	type test struct {
		description string
		want        *dataRecord
		wantErr     bool
		path        string
	}

	aMap := make(map[dataField]float64)
	aMap[TimestampField] = 1561546800
	aMap[SpeedField] = -1
	aMap[FuelField] = 7.914451626
	record := newDataRecord(aMap)

	tests := []test{}
	tests = append(tests, test{
		description: "successful result",
		want:        record,
		wantErr:     false,
		path:        "../res/test.csv",
	})

	tests = append(tests, test{
		description: "failed result with error",
		want:        nil,
		wantErr:     true,
		path:        "fakepath",
	})

	for _, aTest := range tests {
		got, err := parseDataRecordsFromCSV(aTest.path)
		if aTest.want != nil {
			for key, _ := range got[0].dataMap {
				if aTest.want.dataMap[key] != got[0].dataMap[key] {
					t.Errorf("Parsing was incorrect, got: %v, want: %v.", got[0].dataMap[key], aTest.want.dataMap[key])
				}
			}
		}

		if (aTest.wantErr && err == nil) || (!aTest.wantErr && err != nil) {
			t.Errorf("Error requirements failed to match, got: %v, want: %v.", err == nil, aTest.wantErr)
		}
	}
}

// Tester function for correctDataPoints
func TestCorrectDataPoints(t *testing.T) {
	type test struct {
		description string
		want        []dataRecord
		input       []dataRecord
	}

	parsed, _ := parseDataRecordsFromCSV("../res/test.csv")

	tests := []test{}
	tests = append(tests, test{
		description: "successful result",
		want:        testShip.DataRecords,
		input:       parsed,
	})

	tests = append(tests, test{
		description: "empty list",
		want:        []dataRecord{},
		input:       []dataRecord{},
	})

	for _, aTest := range tests {
		correctDataPoints(aTest.input, correctibleDataFields)
		for i := 0; i < len(aTest.input); i++ {
			for key, _ := range aTest.input[i].dataMap {
				if aTest.want[i].dataMap[key] != aTest.input[i].dataMap[key] {
					t.Errorf("Correction was incorrect, got: %v, want: %v.", aTest.input[i].dataMap[key], aTest.want[i].dataMap[key])
				}
			}
		}
	}
}
