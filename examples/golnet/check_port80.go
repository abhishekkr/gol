package main

import (
  "fmt"
  golnet "github.com/abhishekkr/gol/golnet"
)

func main(){
  if golnet.IsPortOpen("80") {
    fmt.Println("Port 80 is open.")
  } else {
    fmt.Println("Port 80 is closed.")
  }

  if golnet.IsPortClosed("8080") {
    fmt.Println("Port 8080 is closed.")
  } else {
    fmt.Println("Port 8080 is open.")
  }
}
