package golzmq

import (
	"fmt"
	"strings"

	zmq "github.com/alecthomas/gozmq"
)

/* returns ZMQ Reply Socket created at given IP-string and Ports-int-array */
func ZmqReplySocket(ip string, reply_ports []int) *zmq.Socket {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	for _, _port := range reply_ports {
		socket.Bind(fmt.Sprintf("tcp://%s:%d", ip, _port))
	}

	return socket
}

/*
performs ZMQ Reply of received message at given Socket processed by given FUNC
where FUNC take String-Array as argument and return result as String
*/
func ZmqReply(_socket *zmq.Socket, _compute_reply RecieveArrayReturnString) error {
	_msg, err := _socket.Recv(0)
	if err != nil {
		return err
	}

	_msg_array := strings.Fields(string(_msg))
	return_value := _compute_reply(_msg_array)

	err = _socket.Send([]byte(return_value), 0)
	return err
}

/*
performs ZMQ Reply of received message at given Socket processed by given FUNC
where FUNC take []byte as argument and return result as []byte
*/
func ZmqReplyByte(_socket *zmq.Socket, _compute_reply RecieveByteArrayReturnByte) error {
	_msg, err := _socket.Recv(0)
	if err != nil {
		return err
	}

	return_value := _compute_reply(_msg)

	err = _socket.Send(return_value, 0)
	return err
}
