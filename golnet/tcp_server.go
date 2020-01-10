package golnet

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var (
	TcpServerHalt = false
)

type RequestParamFunc func(request []byte) (reply []byte)

func TCPServer(connection_string string, request_handler RequestParamFunc) {
	server, err := net.Listen("tcp", connection_string)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()
	for {
		if TcpServerHalt {
			server.Close()
			return
		}
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(connection, request_handler)
	}
}

func handleRequest(conn net.Conn, request_handler RequestParamFunc) {
	request, read_err := ioutil.ReadAll(conn)
	if read_err != nil {
		fmt.Println("ERROR: Golnet TCP server failed to read from client.", read_err)
	}

	reply := request_handler(request)
	_, write_err := conn.Write(reply)
	if write_err != nil {
		fmt.Println("ERROR: Golnet TCP server failed to write to client.", write_err)
	}
	conn.Close()
}
