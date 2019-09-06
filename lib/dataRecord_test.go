package lib

import "testing"

// Tester function for newDataField
func TestNewDataField(t *testing.T) {
	type test struct {
		description string
		want        dataField
		input       string
	}

	tests := []test{}
	tests = append(tests, test{
		description: "successful result",
		want:        SpeedField,
		input:       "speed",
	})

	tests = append(tests, test{
		description: "bad string",
		want:        UndefinedField,
		input:       "wef we",
	})

	for _, aTest := range tests {
		got := newDataField(aTest.input)
		if got != aTest.want {
			t.Errorf("Conversion incorrect, got: %v, want: %v.", got, aTest.want)
		}
	}
}
