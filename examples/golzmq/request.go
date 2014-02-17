package main


import (
  "fmt"

  golzmq "github.com/abhishekkr/gol/golzmq"
)


func main(){
  _socket := golzmq.ZmqRequestSocket("127.0.0.1", 9797, 9898)

  val := golzmq.ZmqRequest(_socket, "get", "anon")
  fmt.Println(val)
  val = golzmq.ZmqRequest(_socket, "put", "ymous")
  fmt.Println(val)
}
