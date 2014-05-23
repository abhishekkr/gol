package golhashmap

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// init: register CSVmap to DataMap
func init() {
	RegisterHashMapEngine("csv", new(CSVmap))
}

type CSVmap struct{}

func HashMapToCSV(hmap HashMap) string {
	csvio := bytes.NewBufferString("")
	csvWriter := csv.NewWriter(csvio)

	for key, value := range hmap {
		key_val := []string{key, value}
		err := csvWriter.Write(key_val)
		if err != nil {
			fmt.Printf("Error converting Hashmap values to CSV:\n%q\n", hmap)
			fmt.Println(err)
		}
	}
	csvWriter.Flush()
	return strings.TrimSpace(csvio.String())
}

func (csvmap CSVmap) FromHashMap(hmap HashMap) string {
	return HashMapToCSV(hmap)
}

func CSVToHashMap(csvalues string) HashMap {
	var hmap HashMap
	hmap = make(HashMap)
	csvalues = strings.TrimSpace(csvalues)
	csvio := bytes.NewBufferString(csvalues)
	csvReader := csv.NewReader(csvio)

	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error converting CSV values to Hashmap:\n%s\n", csvalues)
			fmt.Println(err)
		}
		hmap[fields[0]] = strings.Join(fields[1:], ",")
	}
	return hmap
}

func (csvmap CSVmap) ToHashMap(csvalues string) HashMap {
	return CSVToHashMap(string(csvalues))
}
