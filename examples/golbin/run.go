package main

import (
  "fmt"
  golbin "github.com/abhishekkr/gol/golbin"
)

func main() {
  kon := golbin.Console{Command: "date"}
  kon.Run()
  fmt.Printf("output: %q\n", kon.StdOutput)

  kon = golbin.Console{Command: "echo A B", StdInput: "CDE"}
  kon.Run()
  fmt.Printf("output: %q\n", kon.StdOutput)

	kon = golbin.Console{Command: "cecho C D"}
  kon.Run()
  fmt.Printf("output: %q\n", kon.StdOutput)

  fmt.Printf("output: %q\n", golbin.ExecOutput("uptime"))
}
