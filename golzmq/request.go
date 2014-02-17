package golzmq

import (
  "fmt"
  "strings"

  zmq "github.com/alecthomas/gozmq"
)


func ZmqRequestSocket(ip string, req_port int, rep_port int) *zmq.Socket {
  fmt.Printf("ZMQ REQ/REP Client at port %d and %d\n", req_port, rep_port)
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REQ)
  socket.Connect(fmt.Sprintf("tcp://%s:%d", ip, req_port))
  socket.Connect(fmt.Sprintf("tcp://%s:%d", ip, rep_port))

  return socket
}


func ZmqRequest(_socket *zmq.Socket, dat ...string) string{
  _msg := strings.Join(dat, " ")
  _socket.Send([]byte(_msg), 0)
  response, _ := _socket.Recv(0)

  return string(response)
}
