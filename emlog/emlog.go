package emlog

import (
	"log"
	"os"
)

// NewEmlogStdout retuns a *log.Logger object with the default format which prints messages to standard output.
func NewEmlogStdout(component string) *log.Logger {
	var logger *log.Logger
	logger = log.New(
		os.Stdout,
		"["+component+"]",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)
	return logger
}

// NewEmlogFile retuns a *log.Logger object with the default format which writes messages to the specified file. The
// file is opened in append or create mode. If the file cannot be opened or created, then an error is returned.
func NewEmlogFile(component string, path string) (*log.Logger, error) {
	var logger *log.Logger

	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger = log.New(
		f,
		"["+component+"]",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)
}
