package logger

import (
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/config"
	"log"
	"os"
)

type ServiceLogger struct {
	InfoLogger *log.Logger
	WarnLogger *log.Logger
	ErrLogger  *log.Logger
	loggerFile *os.File
}

func New(cfg *config.Config) (*ServiceLogger, error) {
	osFlags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(cfg.Logger.FilePath, osFlags, 0666)
	if err != nil {
		return nil, err
	}

	logFlags := log.Ldate | log.Ltime | log.Lshortfile
	serviceLogger := &ServiceLogger{
		InfoLogger: log.New(file, "INFO: ", logFlags),
		WarnLogger: log.New(file, "WARN: ", logFlags),
		ErrLogger:  log.New(file, "ERROR: ", logFlags),
	}

	return serviceLogger, nil
}

func (l *ServiceLogger) Info(msg ...string) {
	l.InfoLogger.Println(msg)
}

func (l *ServiceLogger) Warn(msg ...string) {
	l.WarnLogger.Println(msg)
}

func (l *ServiceLogger) Error(msg ...string) {
	l.ErrLogger.Println(msg)
}
