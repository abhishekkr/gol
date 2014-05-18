package golservice

/*
gives you ability to daemon-ize your service via VividCoretx/godaemon
while enabling it to stop/status itself as well like service
*/

import (
	"flag"
	"fmt"

	"github.com/VividCortex/godaemon"
)

type Funk func()

var (
	daemon        = flag.String("daemon", "status", "status|start|stop")
	DaemonPIDFile = flag.String("daemon-pid", "", "path for pidfile of daemon")
	DaemonLogFile = flag.String("daemon-log", "", "path for dumping current status of daemon")
)

/*
full start|stop|status via flag daemon, can just pass service as Funk in main to it
*/
func Daemon(toRun Funk) {
	godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	flag.Parse()

	switch *daemon {
	case "stop":
		Stop()

	case "status":
		Status()

	default:
		Start(toRun)
	}
}

// start for passed Funk typed method call
func Start(toRun Funk) {
	flag.Parse()
	if PersistPID(*DaemonPIDFile) {
		LogDaemon("Started daemon.")
		toRun()
		LogDaemon("Daemon finished task.")
	} else {
		LogDaemon("Daemon seem to already run. Start failed.")
	}
}

// stop for the given process name's stored pid
func Stop() {
	flag.Parse()

	var status string
	if KillPID(*DaemonPIDFile) {
		status = "Status: Stopped."
	} else {
		status = fmt.Sprintf("Failed to stop. Status: %s", StatusPID(*DaemonPIDFile))
	}

	LogDaemon(status)
}

// status for the given process name's stored pid
func Status() {
	flag.Parse()

	LogDaemon(fmt.Sprintf("Status: %s", StatusPID(*DaemonPIDFile)))
}
