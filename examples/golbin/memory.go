package main

import (
  "fmt"
  golbin "github.com/abhishekkr/gol/golbin"
)

func main() {
  fmt.Println(golbin.MemInfo("MemFree"))
  fmt.Println(golbin.MemInfo("MemTotal"))
}
