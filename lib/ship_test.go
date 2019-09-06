package lib

import "testing"

// Tester function for sumDataBetweenTimeFrames
func TestSumDataBetweenTimeFrames(t *testing.T) {
	type test struct {
		description string
		want        float64
		start       float64
		end         float64
		field       dataField
		multi       timeMultiplier
	}

	tests := []test{}
	tests = append(tests, test{
		description: "successful result, normal",
		want:        18,
		start:       1561546800,
		end:         1561550400,
		field:       SpeedField,
		multi:       SecondsPerHour,
	})

	tests = append(tests, test{
		description: "successful result, between points",
		want:        15.5,
		start:       1561546900,
		end:         1561550000,
		field:       SpeedField,
		multi:       SecondsPerHour,
	})

	tests = append(tests, test{
		description: "successful result, edge case",
		want:        50,
		start:       1561540000,
		end:         1561550000,
		field:       SpeedField,
		multi:       SecondsPerHour,
	})

	tests = append(tests, test{
		description: "successful result, edge case, fuel",
		want:        17,
		start:       1561542350,
		end:         1561550000,
		field:       FuelField,
		multi:       SecondsPerHour,
	})

	tests = append(tests, test{
		description: "failed result, bad input",
		want:        -1,
		start:       1561550000,
		end:         1561540000,
		field:       SpeedField,
		multi:       SecondsPerHour,
	})

	for _, aTest := range tests {
		got, _ := testShip.sumDataBetweenTimeFrames(aTest.start, aTest.end, aTest.field, aTest.multi)
		if got != aTest.want {
			t.Errorf("Invalid result, got: %v, want: %v.", got, aTest.want)
		}
	}
}

// Tester function for getEfficiencyBetweenTimeFrames
func TestGetEfficiencyBetweenTimeFrames(t *testing.T) {
	type test struct {
		description string
		want        float64
		start       float64
		end         float64
	}

	tests := []test{}
	tests = append(tests, test{
		description: "successful result, normal",
		want:        0.0375,
		start:       1561546800,
		end:         1561550400,
	})

	tests = append(tests, test{
		description: "successful result, between points",
		want:        0.0375,
		start:       1561546900,
		end:         1561550000,
	})

	tests = append(tests, test{
		description: "successful result, edge case",
		want:        0.0375,
		start:       1561542350,
		end:         1561550000,
	})

	tests = append(tests, test{
		description: "failed result, bad input",
		want:        -1,
		start:       1561550000,
		end:         1561540000,
	})

	for _, aTest := range tests {
		got, _ := testShip.getEfficiencyBetweenTimeFrames(aTest.start, aTest.end)
		if got != aTest.want {
			t.Errorf("Invalid result, got: %v, want: %v.", got, aTest.want)
		}
	}
}
