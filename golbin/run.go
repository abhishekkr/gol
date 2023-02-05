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
Run executes command from Console field with its input
and sets the output or error whatever gets prompted.
*/
func (konsole *Console) Run() (err error) {
	cmd := cmdStringToExecCmd(konsole.Command)
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
	cmd := cmdStringToExecCmd(cmdline)

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
		rest = mergeTokensWithQuotes(parts[1:])
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

func mergeTokensWithQuotes(words []string) []string {
	tokens := []string{}
	var inSingleQuotes, inDoubleQuotes bool
	for _, w := range words {
		if inSingleQuotes || inDoubleQuotes {
			tokens = append(tokens[0:len(tokens)-1], tokens[len(tokens)-1]+" "+w)
			if strings.ContainsRune(w, '\'') && inSingleQuotes {
				inSingleQuotes = false
			} else if strings.ContainsRune(w, '"') && inDoubleQuotes {
				inDoubleQuotes = false
			}
			continue
		}

		tokens = append(tokens, w)
		if strings.ContainsRune(w, '\'') {
			if strings.Count(w, "'")%2 == 1 {
				inSingleQuotes = true
			}
		} else if strings.ContainsRune(w, '"') {
			if strings.Count(w, "\"")%2 == 1 {
				inDoubleQuotes = true
			}
		}
	}
	result := make([]string, len(tokens))
	for idx, token := range tokens {
		result[idx] = stripQuotes(token)
	}
	return result
}

func stripQuotes(w string) string {
	if (w[0] == '\'' && w[len(w)-1] == '\'') || (w[0] == '"' && w[len(w)-1] == '"') {
		return w[1 : len(w)-1]
	}
	return w
}
