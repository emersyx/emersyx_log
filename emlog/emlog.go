package emlog

import(
    "os"
    "log"
)

func NewEmlogStdout(component string) *log.Logger {
    var logger *log.Logger
    logger = golog.New(
        os.Stdout,
        "[" + component + "]",
        golog.Ldate | golog.Ltime | golog.Lmicroseconds | golog.Lshortfile
    )
    return logger
}

func NewEmlogFile(component string, path string) (*log.Logger, error) {
    var logger *log.Logger

    f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }

    logger = golog.New(
        f,
        "[" + component + "]",
        golog.Ldate | golog.Ltime | golog.Lmicroseconds | golog.Lshortfile
    )
}
