package emlog

import(
    "os"
    "log"
)

// NewEmlogStdout retuns a *log.Logger object with the default format which prints messages to standard output.
func NewEmlogStdout(component string) *log.Logger {
    var logger *log.Logger
    logger = golog.New(
        os.Stdout,
        "[" + component + "]",
        golog.Ldate | golog.Ltime | golog.Lmicroseconds | golog.Lshortfile,
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

    logger = golog.New(
        f,
        "[" + component + "]",
        golog.Ldate | golog.Ltime | golog.Lmicroseconds | golog.Lshortfile,
    )
}
