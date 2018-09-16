package logger

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

// LevelLog provides an interface for logging with specific format
func LevelLog(entry *logrus.Entry, level logrus.Level, msg string) {
	switch level {
	case logrus.DebugLevel:
		entry.Debugln(msg)
	case logrus.InfoLevel:
		entry.Infoln(msg)
	case logrus.WarnLevel:
		entry.Warnln(msg)
	case logrus.ErrorLevel:
		entry.Errorln(msg)
	case logrus.FatalLevel:
		entry.Fatalln(msg)
	case logrus.PanicLevel:
		entry.Panicln(msg)
	default:
		entry.Debugln(msg)
	}
}

// DebugInfo combain user fidles with debug information
func DebugInfo(skip int, fidles logrus.Fields) (debugInfo logrus.Fields) {
	skip++
	if pc, file, line, ok := runtime.Caller(skip); ok {
		debugInfo = logrus.Fields{
			"file_name":   filepath.Base(file),
			"line_number": line,
		}
		if funcInfo := runtime.FuncForPC(pc); funcInfo != nil {
			debugInfo["function_name"] = funcInfo.Name()
		}
	}

	for key, value := range fidles {
		debugInfo[key] = value
	}

	return
}
