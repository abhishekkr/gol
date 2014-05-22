package gollist

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// init: register CSVmap to DataMap
func init() {
	RegisterListEngine("csv", new(CSVmap))
}

type CSVmap struct{}

/* convert list elements to single line CSV */
func List_to_csv(list []string) string {
	csvio := bytes.NewBufferString("")
	csvWriter := csv.NewWriter(csvio)

	err := csvWriter.Write(list)
	if err != nil {
		fmt.Printf("Error converting List values to CSV:\n%q\n", list)
		fmt.Println(err)
	}
	csvWriter.Flush()
	return strings.TrimSpace(csvio.String())
}

/* CSVmap proxy for List_to_csv */
func (csvmap CSVmap) FromList(list []string) string {
	return List_to_csv(list)
}

/* entertains each and every field in CSV as element of list */
func Csv_to_list(csvalues string) []string {
	var list []string

	csvalues = strings.TrimSpace(csvalues)
	csvio := bytes.NewBufferString(csvalues)
	csvReader := csv.NewReader(csvio)

	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error converting CSV values to List:\n%s\n", csvalues)
			fmt.Println(err)
		}

		for _, val := range fields {
			list = append(list, val)
		}
	}
	return list
}

/* CSVmap proxy for Csv_to_list */
func (csvmap CSVmap) ToList(csvalues string) []string {
	return Csv_to_list(string(csvalues))
}
