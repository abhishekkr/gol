package golbin

import (
  "strings"
)

func Uptime() string{
  kon := Console{Command: "uptime"}
  kon.Run()
  return strings.Fields(kon.StdOutput)[0]
}
