package golfilesystem

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreateBinaryFile(filename string, blob []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = f.Write(blob)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	return err
}

func AppendToFile(filename, txt string) {
	if filename == "" {
		log.Fatal("Filename can't be empty.")
	}
	fileHandle, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, txt)
	writer.Flush()
}
