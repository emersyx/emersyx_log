package emlog

import (
	"io"
	"log"
	"os"
	"runtime"
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

// printCaller uses the runtime.Caller function to recover information from the call stack and print the location where
// the logging call was executed. If the information can be recovered, then it is printed. Otherwise, the function does
// nothing.
func (el EmersyxLogger) printCaller() {
	// argument values for runtime.Caller:
	// 0 will simply show this line
	// 1 will show the location of the call in either EmersyxLogger.print, EmersyxLogger.printf, EmersyxLogger.println
	// 2 will show the location of the call in one of the public EmersyxLogger methods
	// 3 will show the location of the call to one of the public EmersyxLogger methods, which is what we want
	_, file, line, ok := runtime.Caller(3)
	if ok {
		el.logger.Printf("%s:%d ", file, line)
	}
}

// print calls send messages to the standard logger. Arguments are handled in the manner of fmt.Print.
func (el EmersyxLogger) print(level uint, v ...interface{}) {
	if level <= el.level {
		el.printCaller()
		el.logger.Print(v...)
	}
}

// printf calls send messages to the standard logger. Arguments are handled in the manner of fmt.Printf.
func (el EmersyxLogger) printf(level uint, format string, v ...interface{}) {
	if level <= el.level {
		el.printCaller()
		el.logger.Printf(format, v...)
	}
}

// println calls send messages to the standard logger. Arguments are handled in the manner of fmt.Println.
func (el EmersyxLogger) println(level uint, v ...interface{}) {
	if level <= el.level {
		el.printCaller()
		el.logger.Println(v...)
	}
}

// Fatal calls the print method with the ELFatal level and the given arguments. After the logging message gets
// printed, the os.Exit(1) function is called.
func (el EmersyxLogger) Fatal(v ...interface{}) {
	el.print(ELFatal, v...)
	os.Exit(1)
}

// Fatalf calls the printf method with the ELFatal level and the given arguments. After the logging message gets
// printed, the os.Exit(1) function is called.
func (el EmersyxLogger) Fatalf(format string, v ...interface{}) {
	el.printf(ELFatal, format, v...)
	os.Exit(1)
}

// Fatalln calls the println method with the ELFatal level and the given arguments. After the logging message gets
// printed, the os.Exit(1) function is called.
func (el EmersyxLogger) Fatalln(v ...interface{}) {
	el.println(ELFatal, v...)
	os.Exit(1)
}

// Error calls the print method with the ELError level and the given arguments.
func (el EmersyxLogger) Error(v ...interface{}) {
	el.print(ELError, v...)
}

// Errorf calls the printf method with the ELError level and the given arguments.
func (el EmersyxLogger) Errorf(format string, v ...interface{}) {
	el.printf(ELError, format, v...)
}

// Errorln calls the println method with the ELError level and the given arguments.
func (el EmersyxLogger) Errorln(v ...interface{}) {
	el.println(ELError, v...)
}

// Info calls the print method with the ELInfo level and the given arguments.
func (el EmersyxLogger) Info(v ...interface{}) {
	el.print(ELInfo, v...)
}

// Infof calls the printf method with the ELInfo level and the given arguments.
func (el EmersyxLogger) Infof(format string, v ...interface{}) {
	el.printf(ELInfo, format, v...)
}

// Infoln calls the println method with the ELInfo level and the given arguments.
func (el EmersyxLogger) Infoln(v ...interface{}) {
	el.println(ELInfo, v...)
}

// Debug calls the print method with the ELDebug level and the given arguments.
func (el EmersyxLogger) Debug(v ...interface{}) {
	el.print(ELDebug, v...)
}

// Debugf calls the printf method with the ELDebug level and the given arguments.
func (el EmersyxLogger) Debugf(format string, v ...interface{}) {
	el.printf(ELDebug, format, v...)
}

// Debugln calls the println method with the ELDebug level and the given arguments.
func (el EmersyxLogger) Debugln(v ...interface{}) {
	el.println(ELDebug, v...)
}

// NewEmersyxLogger returns an EmersyxLogger instance with the default format, which writes messages to specified
// io.Writer instance. The component argument is prepended to logs for easier filtering, while the level argument
// controls the verbosity.
func NewEmersyxLogger(writer io.Writer, component string, level uint) (*EmersyxLogger) {
	emlog := new(EmersyxLogger)

	emlog.logger = log.New(
		writer,
		"["+component+"] ",
		log.Ldate|log.Ltime,
	)

	if emlog.logger == nil {
		return nil
	}

	emlog.level = level

	return emlog
}
