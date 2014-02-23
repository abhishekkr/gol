package main


import (
  "fmt"

  golhashmap "github.com/abhishekkr/gol/golhashmap"
)


func main() {
    var hmap golhashmap.HashMap
    hmap = make(golhashmap.HashMap)
    hmap["abc"] = "ABC"
    hmap["def"] = "DEF"
    hmap["d:e:f"] = "D-E-F"

    content := golhashmap.Hashmap_to_csv(hmap)
    fmt.Println(content)

    hmap2 := golhashmap.Csv_to_hashmap(content)
    fmt.Println(hmap2)

    var csv golhashmap.CSVstring
    csv = "1,2\n3,4\n5,6"
    fmt.Println(csv.To_hashmap())
    fmt.Println(hmap2.To_csv())

}
