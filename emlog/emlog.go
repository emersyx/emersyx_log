package emlog

import (
	"io"
	"log"
	"os"
)

// Emlog is an implementation of a logger based on the standard log.Logger type. The Emlog implementation offers
// additional logging functionality
type Emlog struct {
	logger *log.Logger
	level uint
}

// Print calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func (el Emlog) Print(level uint, v ...interface{}) {
	if level <= el.level {
		el.logger.Print(v...)
	}
}

// Printf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func (el Emlog) Printf(level uint, format string, v ...interface{}) {
	if level <= el.level {
		el.logger.Printf(format, v...)
	}
}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func (el Emlog) Println(level uint, v ...interface{}) {
	if level <= el.level {
		el.logger.Println(v...)
	}
}

// NewEmlog returns an Emlog instance with the default format, which writes messages to standard output if the stdout
// argument is true and/or the specified file if the path argument is given. If the file cannot be opened or created,
// then an error is returned. The component argument is prepended to logs for easier filtering, while the level argument
// controls the verbosity.
func NewEmlog(stdout bool, path string, component string, level uint) (Emlog, error) {
	var emlog Emlog
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
