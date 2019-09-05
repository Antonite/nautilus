package lib

type dataField string

const (
	SpeedField     dataField = "speed"
	FuelField      dataField = "fuel"
	UndefinedField dataField = "undefined"
	TimestampField dataField = "timestamp"
)

type dataRecord struct {
	dataMap map[dataField]float64
}

func newDataRecord(dataMap map[dataField]float64) *dataRecord {
	return &dataRecord{
		dataMap: dataMap,
	}
}

// func newDataPoint(key dataField, value float64) dataPoint {
// 	return &dataPoint{
// 		Key:   key,
// 		Value: value,
// 	}
// }

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