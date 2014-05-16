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

func Hashmap_to_csv(hmap HashMap) string {
	csvio := bytes.NewBufferString("")
	csvWriter := csv.NewWriter(csvio)

	for key, value := range hmap {
		key_val := []string{key, value}
		err := csvWriter.Write(key_val)
		if err != nil {
			fmt.Println(err)
		}
	}
	csvWriter.Flush()
	return csvio.String()
}

func (csvmap CSVmap) FromHashMap(hmap HashMap) string {
	return Hashmap_to_csv(hmap)
}

func Csv_to_hashmap(csvalues string) HashMap {
	var hmap HashMap
	hmap = make(HashMap)
	csvio := bytes.NewBufferString(csvalues)
	csvReader := csv.NewReader(csvio)

	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		hmap[fields[0]] = strings.Join(fields[1:], ",")
	}
	return hmap
}

func (csvmap CSVmap) ToHashMap(csvalues string) HashMap {
	return Csv_to_hashmap(string(csvalues))
}
