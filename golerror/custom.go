package golerror

import "fmt"

type csutomError struct {
	ExitCode    int
	ExitMessage string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[error] %d - %s", e.ExitCode, e.ExitMessage)
}

func CustomError(exitcode int, msg string) error {
	return &argError{exitcode, msg}
}
