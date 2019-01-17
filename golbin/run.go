package golbin

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/abhishekkr/gol/golerror"
)

/*
Console is structure to contain Command Line, Input and Output.
*/
type Console struct {
	Command, StdInput, StdOutput string
	Process                      *os.Process
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
func (konsole *Console) Run() (err error) {
	cmd := startCommand(konsole.Command)
	if konsole.StdInput != "" {
		cmd.Stdin = strings.NewReader(konsole.StdInput)
	}
	konsole.Process = cmd.Process

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err == nil {
		konsole.StdOutput = out.String()
	} else {
		konsole.StdOutput = fmt.Sprintf("Stdout: %s\nError: %s", out.String(), err.Error())
	}
	return
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

/*
Run it for distro to manage opening with correct program
*/
func RunWithAssignedApp(runThis string) string {
	var openWith string
	if runtime.GOOS == "linux" {
		openWith = "xdg-open"
	} else if runtime.GOOS == "darwin" {
		openWith = "open"
	} else {
		return fmt.Sprintf("Error: %s is not supported as yet.", runtime.GOOS)
	}

	if !IsSystemCmd(openWith) {
		return fmt.Sprintf("Error: %s not found on this machine.", openWith)
	}

	cmdToRun := fmt.Sprintf("%s %s", openWith, runThis)
	return ExecOutput(cmdToRun)
}

func cmdStringToExecCmd(cmd string) *exec.Cmd {
	parts := strings.Fields(cmd)
	first := parts[0]
	rest := []string{}
	if len(parts) > 1 {
		rest = parts[1:]
	}
	return exec.Command(first, rest...)
}

func execCmd(cmdHandle *exec.Cmd) (out string, err error) {
	var stdout, stderr bytes.Buffer
	cmdHandle.Stdout = &stdout
	cmdHandle.Stderr = &stderr
	err = cmdHandle.Run()

	if err != nil {
		err = golerror.Error(127, fmt.Sprintf("'%s' :: '%s'", stderr.String(), err.Error()))
	}

	out = stdout.String()
	return
}

func Exec(cmd string) (string, error) {
	cmdHandle := cmdStringToExecCmd(cmd)
	return execCmd(cmdHandle)
}

func ExecWithEnv(cmd string, env map[string]string) (string, error) {
	cmdHandle := cmdStringToExecCmd(cmd)
	cmdHandle.Env = os.Environ()
	for envVar, envVal := range env {
		cmdHandle.Env = append(cmdHandle.Env, fmt.Sprintf("%s=%s", envVar, envVal))
	}

	return execCmd(cmdHandle)
}
