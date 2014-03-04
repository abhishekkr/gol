package main

import (
  "fmt"
  golbin "github.com/abhishekkr/gol/golbin"
)

func main() {
  kon := golbin.Console{Command: "echo A B", StdInput: "CDE"}
  kon.Run()
  fmt.Printf("StdOutput: %q\n", kon.StdOutput)
  kon = golbin.Console{Command: "cecho C D"}
  kon.Run()
  fmt.Printf("StdOutput: %q\n", kon.StdOutput)
  fmt.Println(golbin.Run2("echo A B"))
}
