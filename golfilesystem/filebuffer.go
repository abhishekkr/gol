package golfilesystem

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func FileBuffer(filepath string) (*bytes.Buffer, error) {
	const chunksize int = 1024
	var err error
	var count int

	data, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer := bytes.NewBuffer(make([]byte, 0))
	part := make([]byte, chunksize)

	for {
		count, err = reader.Read(part)
		if err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		return nil, err
	}
	return buffer, nil
}
