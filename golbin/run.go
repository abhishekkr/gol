package golbin

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

/*
Console is structure to contain Command Line, Input and Output.
*/
type Console struct {
	Command, StdInput, StdOutput string
}

/*
startCommand runs passed command string and returns it.
*/
func startCommand(sysCommand string) *exec.Cmd {
	cmdTokens := strings.Split(sysCommand, " ")
	cmd := cmdTokens[0]
	if len(cmdTokens) == 1 {
		return exec.Command(cmd)
	}
	return exec.Command(cmd, strings.Join(cmdTokens[1:], " "))
}

/*
Run executes command from Console field with its input
and sets the output or error whatever gets prompted.
*/
func (konsole *Console) Run() {
	cmd := startCommand(konsole.Command)
	if konsole.StdInput != "" {
		cmd.Stdin = strings.NewReader(konsole.StdInput)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err == nil {
		konsole.StdOutput = out.String()
	} else {
		konsole.StdOutput = fmt.Sprintf("Error: %s", err.Error())
	}
}

/*
ExecOutput can be passed a command to quickly get its output or error.
*/
func ExecOutput(cmdline string) string {
	cmd := startCommand(cmdline)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err == nil {
		return out.String()
	}
	return fmt.Sprintf("Error: %s", err.Error())
}
