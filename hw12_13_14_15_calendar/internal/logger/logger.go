package logger

import (
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/config"
	"log"
	"os"
)

type Logger struct {
	InfoLogger *log.Logger
	WarnLogger *log.Logger
	ErrLogger  *log.Logger
	loggerFile *os.File
}

func New(cfg *config.LoggerConfig) (*Logger, error) {
	logFlags := log.Ldate | log.Ltime | log.Lshortfile
	if !cfg.EnableFileLog {
		return &Logger{
			InfoLogger: log.New(os.Stdout, "INFO: ", logFlags),
			WarnLogger: log.New(os.Stdout, "WARN: ", logFlags),
			ErrLogger:  log.New(os.Stdout, "ERROR: ", logFlags),
		}, nil
	}

	osFlags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(cfg.FilePath, osFlags, 0666)
	if err != nil {
		return nil, err
	}

	serviceLogger := &Logger{
		InfoLogger: log.New(file, "INFO: ", logFlags),
		WarnLogger: log.New(file, "WARN: ", logFlags),
		ErrLogger:  log.New(file, "ERROR: ", logFlags),
	}

	return serviceLogger, nil
}

func (l *Logger) Info(msg ...string) {
	l.InfoLogger.Println(msg)
}

func (l *Logger) Warn(msg ...string) {
	l.WarnLogger.Println(msg)
}

func (l *Logger) Error(msg ...string) {
	l.ErrLogger.Println(msg)
}
