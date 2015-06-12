package golbin

import (
	"errors"
	"fmt"
	"os"

	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
)

type ConsolePipe struct {
	Konsole  Console
	PipePath string
}

/*
Run executes command from console with a pipe to get back.
*/
func (konsolePipe *ConsolePipe) Run() error {
	if konsolePipe.PipePath == "" {
		return errors.New(fmt.Sprintf("Pipe path not configured."))
	}
	if golfilesystem.PathExists(konsolePipe.PipePath) {
		return errors.New(fmt.Sprintf("%s path already exists!", konsolePipe.PipePath))
	}
	ExecOutput(fmt.Sprintf("mkfifo %s", konsolePipe.PipePath))

	konsolePipe.Konsole.Command = fmt.Sprintf("%s < %s", konsolePipe.Konsole.Command, konsolePipe.PipePath)
	konsolePipe.Konsole.Run()
	return nil
}

/*
Pipe pipes given input to PipePath.
*/
func (konsolePipe *ConsolePipe) Pipe(input string) error {
	if !golfilesystem.PathExists(konsolePipe.PipePath) {
		return errors.New(fmt.Sprintf("%s path doesn't exists!", konsolePipe.PipePath))
	}
	cmd := fmt.Sprintf("echo '%s' > %s", input, konsolePipe.PipePath)
	ExecOutput(cmd)
	return nil
}

/*
Clean cleans up the process and then the pipe.
*/
func (konsolePipe *ConsolePipe) Clean() {
	konsolePipe.Konsole.Process.Kill()
	os.Remove(konsolePipe.PipePath)
}
