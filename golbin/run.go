package golbin

import (
  "bytes"
  "os/exec"
  "strings"
  "fmt"
)

type Console struct {
    Command, StdInput, StdOutput string
}

func start_Command(sys_Command string) *exec.Cmd{
  cmd_tokens := strings.Split(sys_Command, " ")
  cmd := cmd_tokens[0]
  args := strings.Join(cmd_tokens[1:], " ")
  return exec.Command(cmd, args)
}

func (konsole *Console) Run() {
  cmd := start_Command(konsole.Command)

  if konsole.StdInput != ""{ cmd.Stdin = strings.NewReader(konsole.StdInput) }

  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err == nil {
    konsole.StdOutput = out.String()
  } else {
    konsole.StdOutput = fmt.Sprintf("Error: %s", err.Error())
  }
}

func ExecOutput(cmdline string) string {
  cmd := start_Command(cmdline)

  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err == nil {
    return out.String()
  }
  return fmt.Sprintf("Error: %s", err.Error())
}
