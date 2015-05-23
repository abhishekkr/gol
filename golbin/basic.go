package golbin

import (
	"strings"
)

/*
Uptime proxies linux system command for 'uptime'.
*/
func Uptime() string {
	kon := Console{Command: "uptime"}
	kon.Run()
	return strings.Fields(kon.StdOutput)[0]
}
