package golhashmap

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// init: register CSVMap to DataMap
func init() {
	RegisterDataMap("csv", new(CSVMap))
}

// CSV's DataMap struct
type CSVMap struct {
	HMap      HashMap
	CSVString string
}

// getter CSVMap's HashMap
func (csvmap *CSVMap) GetHashMap() HashMap {
	return csvmap.HMap
}

// getter CSVMap's CSVString
func (csvmap *CSVMap) GetDataMap() []byte {
	return []byte(csvmap.CSVString)
}

// setter CSVMap's HashMap
func (csvmap *CSVMap) SetHashMap(hashmap HashMap) {
	csvmap.HMap = make(HashMap)
	csvmap.HMap = hashmap
}

// setter CSVMap's CSVString
func (csvmap *CSVMap) SetDataMap(csv_bytes []byte) {
	csvmap.CSVString = string(csv_bytes)
}

// CSV DataMap's Encode csv to hashmap
func (csvmap *CSVMap) EncodeToHashMap() {
	csvmap.HMap = csv_to_hashmap(csvmap.CSVString)
}

// CSV DataMap's Decode hashmap to csv
func (csvmap *CSVMap) DecodeToDataMap() {
	csvmap.CSVString = hashmap_to_csv(csvmap.HMap)
}

// private: convert hashmap to csv string
func hashmap_to_csv(hmap HashMap) string {
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

// private: convert csv string to hashmap
func csv_to_hashmap(csvalues string) HashMap {
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
