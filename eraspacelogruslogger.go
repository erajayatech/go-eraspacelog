package eraspacelog

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger(env string) {
	logrus.New()

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

	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	otelHook := NewOtelTraceHook(&TraceHookConfig{
		RecordStackTraceInSpan: true,
		EnableLevels:           logrus.AllLevels,
		ErrorSpanLevel:         logrus.ErrorLevel,
	})

	logrus.AddHook(otelHook)
}
