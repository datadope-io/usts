package usts

import (
	"log"
	"os"
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Logger interface
//_______________________________________________________________________

// Logger interface is to abstract the logging from Usts. Gives control to
// the USTS users, choice of the logger.
type Logger interface {
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Tracef(format string, v ...interface{})
}

func createLogger() *logger {
	l := &logger{l: log.New(os.Stderr, "", log.Ldate|log.Ltime)}
	return l
}

var _ Logger = (*logger)(nil)

type logger struct {
	l *log.Logger
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.output("USTS ERROR: "+format, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.output("USTS WARN: "+format, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.output("USTS INFO: "+format, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.output("USTS DEBUG: "+format, v...)
}

func (l *logger) Tracef(format string, v ...interface{}) {
	l.output("USTS TRACE: "+format, v...)
}

func (l *logger) output(format string, v ...interface{}) {
	if len(v) == 0 {
		l.l.Print(format)
		return
	}
	l.l.Printf(format, v...)
}
