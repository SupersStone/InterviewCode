package orm

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// DBLog 定义自己的Writer
type DBLog struct {
	log *logrus.Logger
}

// Printf 实现gorm/logger.Writer接口
func (m *DBLog) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	m.log.Info(logstr)
}

// NewDBLog new obj
func NewDBLog() *DBLog {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &DBLog{log: log}
}
