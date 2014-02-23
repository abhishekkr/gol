package golzmq

import (
  "fmt"
  "strings"

  zmq "github.com/alecthomas/gozmq"
)


func ZmqRequestSocket(ip string, request_ports ...int) *zmq.Socket {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REQ)
  for _, _port := range request_ports {
    socket.Connect(fmt.Sprintf("tcp://%s:%d", ip, _port))
  }
  return socket
}


func ZmqRequest(_socket *zmq.Socket, dat ...string) (string, error){
  _msg := strings.Join(dat, " ")
  _socket.Send([]byte(_msg), 0)
  response, err := _socket.Recv(0)
  return string(response), err
}
