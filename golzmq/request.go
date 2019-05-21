package golzmq

import (
	"fmt"
	"strings"

	zmq "github.com/alecthomas/gozmq"
)

/* returns ZMQ Request Socket created at given IP-string and Ports-int-array */
func ZmqRequestSocket(ip string, requestPorts []int) *zmq.Socket {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REQ)
	for _, _port := range requestPorts {
		socket.Connect(fmt.Sprintf("tcp://%s:%d", ip, _port))
	}
	return socket
}

/* performs ZMQ Request at given Socket for given string data */
func ZmqRequest(_socket *zmq.Socket, dat ...string) (string, error) {
	_msg := strings.Join(dat, " ")
	_socket.Send([]byte(_msg), 0)
	response, err := _socket.Recv(0)
	return string(response), err
}

/* performs ZMQ Request at given Socket for given message as []byte */
func ZmqRequestByte(_socket *zmq.Socket, _msg []byte) ([]byte, error) {
	_socket.Send(_msg, 0)
	return _socket.Recv(0)
}
