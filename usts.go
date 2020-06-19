package usts

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
	//log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)
}
