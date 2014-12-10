package golnet

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type TCPClient struct {
	Connection net.Conn
}

func CreateTCPClient(connection_string string) TCPClient {
	connection, err := net.Dial("tcp", connection_string)
	if err != nil {
		fmt.Println("ERROR: Golnet TCP client failed to connect to server.", err)
		os.Exit(1)
	}
	return TCPClient{
		Connection: connection,
	}
}

func (client TCPClient) Request(send []byte) []byte {
	_, write_err := client.Connection.Write(send)
	if write_err != nil {
		fmt.Println("ERROR: Golnet TCP client failed to write to server.", write_err)
	}
	client.Connection.(*net.TCPConn).CloseWrite()

	read_buffer, read_err := ioutil.ReadAll(client.Connection)
	if read_err != nil {
		fmt.Println("ERROR: Golnet TCP client failed to read from server.", read_err)
	}
	return read_buffer
}
