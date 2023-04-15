package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once = &sync.Once{}
	log  *Logger
)

type Logger struct {
	*logrus.Entry
}

func New() *Logger {
	once.Do(func() {
		l := logrus.New()
		l.SetReportCaller(true)
		l.Formatter = &logrus.TextFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				filename := path.Base(frame.File)
				return fmt.Sprintf("%s %d", filename, frame.Line), frame.Function
			},
			DisableColors: false,
			FullTimestamp: true,
		}

		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		log = &Logger{logrus.NewEntry(l)}
	})

	return log
}
