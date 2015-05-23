package golbin

import (
	"fmt"
	"strings"
)

/*
MemInfo returns asked field value from /proc/meminfo.
*/
func MemInfo(field string) string {
	kon := Console{Command: "cat /proc/meminfo"}
	kon.Run()
	kon = Console{Command: fmt.Sprintf("grep %s", field), StdInput: kon.StdOutput}
	kon.Run()
	return strings.Fields(kon.StdOutput)[1]
}
