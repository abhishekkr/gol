package golhashmap

import (
    "fmt"
    "encoding/csv"
    "bytes"
    "io"
    "strings"
)


type CSVstring string


func Hashmap_to_csv(hmap HashMap) string{
  csvio := bytes.NewBufferString("")
  csvWriter := csv.NewWriter(csvio)

  for  key, value := range hmap {
    key_val := []string {key, value}
    err := csvWriter.Write(key_val)
    if err != nil { fmt.Println(err) }
  }
  csvWriter.Flush()
  return csvio.String()
}

func (hmap HashMap) To_csv() CSVstring{
  return CSVstring(Hashmap_to_csv(hmap))
}


func Csv_to_hashmap(csvalues string) HashMap{
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
    hmap[fields[0]] = strings.Join(fields[1:],",")
  }
  return hmap
}

func (csvalues CSVstring) To_hashmap() HashMap{
  return Csv_to_hashmap(string(csvalues))
}
