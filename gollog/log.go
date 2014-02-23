package gollog


import (
  "os"
  "fmt"
  "log"
)


var LogFile string


// Log_it logs any string passed to it, into LogFile set as filepath for this package.
func Log_it(msg string){
    f, err := os.OpenFile(LogFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v\n", err)
    }
    defer f.Close()

    log.SetOutput(f)
    log.Println(msg)
}
