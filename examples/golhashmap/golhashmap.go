package main

import (
  "fmt"

  golhashmap "github.com/abhishekkr/gol/golhashmap"
)

func main(){
  j := golhashmap.GetDataMap("json")
  c := golhashmap.GetDataMap("csv")

  j.SetDataMap([]byte("{\"1\": \"23\", \"4\": \"56\"}"))
  c.SetDataMap([]byte("1,23\n4,56\n"))

  j.EncodeToHashMap()
  c.EncodeToHashMap()

  fmt.Println(j.GetHashMap())
  fmt.Println(c.GetHashMap())

  var hmap golhashmap.HashMap
  hmap = make(golhashmap.HashMap)
  hmap["abc"] = "ABC"
  hmap["xyz"] = "XYZ"
  j.SetHashMap(hmap)
  c.SetHashMap(hmap)

  fmt.Println(string(j.GetDataMap()))
  fmt.Println(string(c.GetDataMap()))
}
