package main


import (
  "fmt"
  "time"

  "github.com/abhishekkr/gol"
)


func wait_n_shout(){
  fmt.Println("Gol")
  time.Sleep(1 * time.Second)
}

func main(){
  gol.Gol(wait_n_shout)
}
