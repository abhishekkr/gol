package golerror

import "fmt"

type CustomError struct {
	ExitCode    int
	ExitMessage string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[error] %d - %s", e.ExitCode, e.ExitMessage)
}

func Error(exitcode int, msg string) error {
	return &CustomError{exitcode, msg}
}
