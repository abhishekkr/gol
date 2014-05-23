package gollist

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// init: register CSVmap to DataMap
func init() {
	RegisterListEngine("csv", new(CSVmap))
}

type CSVmap struct{}

/* convert list elements to single line CSV */
func ListToCSV(list []string) string {
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

/* CSVmap proxy for ListToCSV */
func (csvmap CSVmap) FromList(list []string) string {
	return ListToCSV(list)
}

/* entertains each and every field in CSV as element of list */
func CSVToList(csvalues string) []string {
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

/* CSVmap proxy for CSVToList */
func (csvmap CSVmap) ToList(csvalues string) []string {
	return CSVToList(string(csvalues))
}

/* CSVmap proxy for CSVToList */
func CSVToNumbers(csvalues string) ([]int, error) {
	var (
		numElement int
		err        error
		numList    []int
	)

	list := CSVToList(string(csvalues))
	numList = make([]int, len(list))

	for idx, element := range list {
		numElement, err = strconv.Atoi(element)
		numList[idx] = numElement
	}
	if err != nil {
		return numList, err
	}
	return numList, nil
}

/* CSVmap proxy for CSVToNumbers */
func (csvmap CSVmap) ToNumbers(csvalues string) ([]int, error) {
	return CSVToNumbers(string(csvalues))
}
