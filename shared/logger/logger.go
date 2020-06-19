package logger

import (
	"context"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/nodias/golang.templete.common/models"
)

var log *logrus.Logger
var config models.TomlConfig

func Init() {
	config = *models.GetConfig()
	log = &logrus.Logger{
		Out:   os.Stderr,
		Hooks: make(logrus.LevelHooks),
		Level: logrus.Level(config.Logconfig.Loglevel),
		Formatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "log.level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "function.name", //non-ECS
			},
		},
	}
	fpLog, err := os.OpenFile(config.Logconfig.Logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)
	log.Debug("logger init - success")
}

var instance *logrus.Logger
var once sync.Once

func Get() *logrus.Logger {
	once.Do(func() {
		instance = log
	})
	return instance
}

func New(ctx context.Context) *logrus.Logger {
	return Get()
}
