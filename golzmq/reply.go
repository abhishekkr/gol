package golzmq

import (
  "fmt"
  "strings"

  zmq "github.com/alecthomas/gozmq"
)


type RecieveArrayReturnString func(msg_array []string) string


func ZmqReplySocket(ip string, req_port int, rep_port int) *zmq.Socket {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REP)
  socket.Bind(fmt.Sprintf("tcp://%s:%d", ip, req_port))
  socket.Bind(fmt.Sprintf("tcp://%s:%d", ip, rep_port))

  return socket
}


func ZmqReply(_socket *zmq.Socket, _compute_reply RecieveArrayReturnString){
  for {
    _msg, _ := _socket.Recv(0)
    _msg_array := strings.Fields(string(_msg))
    return_value := _compute_reply(_msg_array)

    _socket.Send([]byte(return_value), 0)
  }
}
