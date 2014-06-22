package gollog_example

import "os"

type ReceiveStringReturnNil func(log *os.File, msg string)

func PassAndLog(foo ReceiveStringReturnNil, logr *os.File) {
	foo(logr, "passed and logged")
}
