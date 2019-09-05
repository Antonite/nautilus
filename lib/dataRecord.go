package lib

// A DataField, used to manipulate data based on column headers
type dataField string

const (
	SpeedField     dataField = "speed"
	FuelField      dataField = "fuel"
	UndefinedField dataField = "undefined"
	TimestampField dataField = "timestamp"
)

// dataRecord struct
// dataMap - a map of a dataField to value
type dataRecord struct {
	dataMap map[dataField]float64
}

// a list of fields that can be cleaned up after loading data
var correctibleDataFields = []dataField{SpeedField, FuelField}

// Given a dataMap, create a new dataRecord.
// Returns a pointer to the created dataRecord, which holds the given dataMap
func newDataRecord(dataMap map[dataField]float64) *dataRecord {
	return &dataRecord{
		dataMap: dataMap,
	}
}

// Given a string field name, convert to respective dataField.
func newDataField(field string) dataField {
	switch field {
	case "speed":
		return SpeedField
	case "fuel":
		return FuelField
	case "timestamp":
		return TimestampField
	default:
		return UndefinedField
	}
}
