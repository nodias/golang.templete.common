package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(ctx context.Context) *logrus.Entry {
	return Log.WithContext(ctx)
}

var Log = &logrus.Logger{
	Out:   os.Stderr,
	Hooks: make(logrus.LevelHooks),
	Level: logrus.InfoLevel,
	Formatter: &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "agro", //non-ECS
		},
	},
}
