package main


import (
  "fmt"

  golzmq "github.com/abhishekkr/gol/golzmq"
)


func get_n_put(messages []string) string{
  var ret_val string
  axn, msg := messages[0], messages[1]
  if axn == "get" {
    ret_val = fmt.Sprintf("GET: %s", msg)
  } else if axn == "put" {
    ret_val = fmt.Sprintf("PUT: %s", msg)
  } else {
    ret_val = fmt.Sprintf("unhandled request sent: %s", msg)
  }
  return ret_val
}


func main(){
  socket := golzmq.ZmqReplySocket("127.0.0.1", 9797, 9898)
  golzmq.ZmqReply(socket, get_n_put)
}
