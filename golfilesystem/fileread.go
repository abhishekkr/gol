package golfilesystem

import (
	"fmt"
	"os"
)

func ReadBinaryFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	filesize := fileinfo.Size()

	buffer := make([]byte, filesize)
	if _, err := file.Read(buffer); err != nil {
		return []byte{}, err
	}
	return buffer, nil
}
