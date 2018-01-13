package emlog

import (
	"io"
	"log"
	"os"
)

const (
	// ELNone completely turns off logging
	ELNone = 0
	// ELFatal sets logging to print messages of fatal events, which surely lead to a crash
	ELFatal = 1
	// ELError sets logging to print error messages, which may lead to a crash
	ELError = 2
	// ELInfo sets logging to print informational messages about the progress of tasks
	ELInfo = 3
	// ELDebug logs verbose debug messages
	ELDebug = 4
)

// EmersyxLogger is an implementation of a logger based on the standard log.Logger type. The EmersyxLogger
// implementation offers additional logging functionality
type EmersyxLogger struct {
	logger *log.Logger
	level  uint
}

// Print calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func (el EmersyxLogger) Print(level uint, v ...interface{}) {
	if level <= el.level {
		el.logger.Print(v...)
	}
}

// Printf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func (el EmersyxLogger) Printf(level uint, format string, v ...interface{}) {
	if level <= el.level {
		el.logger.Printf(format, v...)
	}
}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func (el EmersyxLogger) Println(level uint, v ...interface{}) {
	if level <= el.level {
		el.logger.Println(v...)
	}
}

// NewEmersyxLogger returns an EmersyxLogger instance with the default format, which writes messages to standard output
// if the stdout argument is true and/or the specified file if the path argument is given. If the file cannot be opened
// or created, then an error is returned. The component argument is prepended to logs for easier filtering, while the
// level argument controls the verbosity.
func NewEmersyxLogger(stdout bool, path string, component string, level uint) (EmersyxLogger, error) {
	var emlog EmersyxLogger
	var sinks []io.Writer

	if stdout {
		sinks = append(sinks, os.Stdout)
	}

	if len(path) > 0 {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return emlog, err
		}
		sinks = append(sinks, f)
	}

	emlog.logger = log.New(
		io.MultiWriter(sinks...),
		"["+component+"] ",
		log.Ldate|log.Ltime,
	)
	emlog.level = level

	return emlog, nil
}
