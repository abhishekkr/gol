package golzmq

import (
  "fmt"
  "strings"

  zmq "github.com/alecthomas/gozmq"
)


func ZmqReplySocket(ip string, reply_ports ...int) *zmq.Socket {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REP)
  for _, _port := range reply_ports {
    socket.Bind(fmt.Sprintf("tcp://%s:%d", ip, _port))
  }

  return socket
}


func ZmqReply(_socket *zmq.Socket, _compute_reply RecieveArrayReturnString) error{
  _msg, err := _socket.Recv(0)
  if err != nil { return err }

  _msg_array := strings.Fields(string(_msg))
  return_value := _compute_reply(_msg_array)

  err = _socket.Send([]byte(return_value), 0)
  return err
}
