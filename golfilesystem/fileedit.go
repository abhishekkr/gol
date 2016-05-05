package golfilesystem

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

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
