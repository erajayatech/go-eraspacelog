package eraspacelog

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var eraspaceLog *logrus.Logger

type Fields = logrus.Fields

func SetupLogger(env string) {
	loger := logrus.New()

	formatter := Formatter{
		ChildFormatter: &logrus.JSONFormatter{},
		Line:           true,
		File:           true,
	}

	if env == "local" {
		formatter = Formatter{
			ChildFormatter: &logrus.TextFormatter{
				ForceColors:   true,
				FullTimestamp: true,
			},
			Line: true,
			File: true,
		}
	}

	loger.SetFormatter(&formatter)
	loger.SetOutput(os.Stdout)
	loger.SetLevel(logrus.InfoLevel)

	otelHook := NewOtelTraceHook(&TraceHookConfig{
		RecordStackTraceInSpan: true,
		EnableLevels:           logrus.AllLevels,
		ErrorSpanLevel:         logrus.ErrorLevel,
	})

	eraspaceLog = loger
	eraspaceLog.AddHook(otelHook)
}

func WithContext(ctx context.Context, field map[string]interface{}) *logrus.Entry {
	return eraspaceLog.WithContext(ctx).WithFields(field)
}
