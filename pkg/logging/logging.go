package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}

	return err
}

func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func init() {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Func), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		FullTimestamp: true,
	}

	if _, err := os.Stat("logs"); err != nil {
		if !os.IsNotExist(err) {
			err = os.Mkdir("logs", 0644)
		} else {
			panic(err)
		}
	}

	allFiles, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 755)
	if err != nil {
		panic(err)
	}

	log.SetOutput(io.Discard)

	log.AddHook(&writeHook{
		Writer:    []io.Writer{allFiles, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	log.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(log)
}
