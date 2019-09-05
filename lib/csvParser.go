package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func parseDataRecordsFromCSV(path string) ([]dataRecord, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open csv.")
	}
	defer f.Close()

	var dataRecords []dataRecord

	scanner := bufio.NewScanner(f)
	// scan header
	scanner.Scan()
	if scanner.Err() != nil {
		return dataRecords, errors.Wrap(err, "Failed to parse CSV header")
	}

	// build datafield array of headers
	header := scanner.Text()
	var headers []dataField
	for _, field := range strings.Split(header, ",") {
		headers = append(headers, newDataField(field))
	}

	columnCount := len(headers)

	// scan file line by line
	for scanner.Scan() {
		text := scanner.Text()
		columns := strings.Split(text, ",")

		// create data points
		dataMap := make(map[dataField]float64)
		for i := 0; i < columnCount; i++ {
			// mark each column for correction in case no data is present
			parsed := -1.00

			// if data is present for this column
			if columns[i] != "" {
				var parseErr error
				parsed, parseErr = strconv.ParseFloat(columns[i], 64)
				if parseErr != nil {
					log.Printf("Failed to parse string to float %v\n", columns[i])
				}
			}

			dataMap[headers[i]] = parsed
		}

		if err != nil {
			return dataRecords, err
		}

		dataRecords = append(dataRecords, *newDataRecord(dataMap))
	}

	if scanner.Err() != nil {
		return dataRecords, errors.Wrap(err, "Failed to scan csv")
	}

	return dataRecords, nil
}

func correctDataPoints(dataRecords []dataRecord, fieldsToCorrect []dataField) {
	for _, field := range fieldsToCorrect {
		index := -1
		recordsCount := len(dataRecords)
		for i := 0; i < recordsCount; i++ {
			record := dataRecords[i]

			// check if field needs correction
			if record.dataMap[field] == -1 && index == -1 {
				index = i
			} else {
				if record.dataMap[field] == -1 {
					continue
				}

				// correct empty fields
				if index != -1 {
					endPoint := record.dataMap[field]
					// check for empty start edgecase
					if index == 0 {
						// in this case, set all missing values to the first available value
						for si := index; si < i; si++ {
							dataRecords[si].dataMap[field] = endPoint
						}
					} else {
						startPoint := dataRecords[index-1].dataMap[field]
						// average increments for all missing points
						increment := (endPoint - startPoint) / float64(i-index+1)
						// fix each point at speed index
						for si := index; si < i; si++ {
							dataRecords[si].dataMap[field] = startPoint + increment*(float64(si-index+1))
						}
					}

					index = -1
				}
			}
		}

		// edge case if ending points were missing data
		// in this case, set all missing values to the first available value
		if index != -1 {
			var val float64

			// edge case if no points had data, set all values to zero
			if index == 0 {
				val = 0
			} else {
				val = dataRecords[index-1].dataMap[field]
			}

			for si := index; si < recordsCount; si++ {
				dataRecords[si].dataMap[field] = val
			}
		}
	}
}
