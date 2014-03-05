package golbin

import (
  "strings"
)

func MemInfo(field string) string{
  kon := Console{Command: "cat /proc/meminfo"}
  kon.Run()
  kon = Console{Command: "grep MemFree", StdInput: kon.StdOutput}
  kon.Run()
  return strings.Fields(kon.StdOutput)[1]
}
