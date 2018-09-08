// Package logger provides both console and file system logging.
package logger

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	colorable "github.com/mattn/go-colorable"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var loggers = map[string]*logrus.Logger{}

// New create a new logger.
func New(name string, dev bool) (logger *logrus.Logger) {
	// If Logger had been created, return nil
	if loggers[name] != nil {
		return
	}

	// Create log file in temp folder.
	logFile, err := ioutil.TempFile("", name+".*.log")
	if err != nil {
		log.Fatal(err)
	}
	logFileName := logFile.Name()

	// Create logger.
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	// Enable color logging in Windows console.
	logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	logger.SetOutput(colorable.NewColorableStdout())

	// Formatter for file logging.
	fileFormatter := &prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true, DisableColors: true, SpacePadding: 64}

	// Add file hook.
	logger.Hooks.Add(lfshook.NewHook(
		lfshook.PathMap{
			logrus.DebugLevel: logFileName,
			logrus.InfoLevel:  logFileName,
			logrus.WarnLevel:  logFileName,
			logrus.ErrorLevel: logFileName,
			logrus.FatalLevel: logFileName,
			logrus.PanicLevel: logFileName,
		},
		fileFormatter,
	))

	// Enable developing logging.
	if dev {
		os.RemoveAll("dev_full.log")
		os.RemoveAll("dev_warn.log")
		os.RemoveAll("dev_erro.log")
		os.RemoveAll("dev_crit.log")

		// Full log file.
		logger.Hooks.Add(lfshook.NewHook(
			lfshook.PathMap{
				logrus.DebugLevel: "dev_full.log",
				logrus.InfoLevel:  "dev_full.log",
				logrus.WarnLevel:  "dev_full.log",
				logrus.ErrorLevel: "dev_full.log",
				logrus.FatalLevel: "dev_full.log",
				logrus.PanicLevel: "dev_full.log",
			},
			fileFormatter,
		))

		// Warning only log file.
		logger.Hooks.Add(lfshook.NewHook(
			lfshook.PathMap{
				logrus.WarnLevel: "dev_warn.log",
			},
			fileFormatter,
		))

		// Error only log file.
		logger.Hooks.Add(lfshook.NewHook(
			lfshook.PathMap{
				logrus.ErrorLevel: "dev_erro.log",
			},
			fileFormatter,
		))

		// Critical only log file.
		logger.Hooks.Add(lfshook.NewHook(
			lfshook.PathMap{
				logrus.FatalLevel: "dev_crit.log",
				logrus.PanicLevel: "dev_crit.log",
			},
			fileFormatter,
		))
	}

	// Store logger
	loggers[name] = logger

	return
}

// Get find the logger from storage
func Get(name string) (logger *logrus.Logger) {
	logger = loggers[name]
	return
}

// List shows created loggers
func List() (keys []string) {
	for k := range loggers {
		keys = append(keys, k)
	}
	return
}

// Has returns true if the logger created before
func Has(name string) (ok bool) {
	_, ok = loggers[name]
	return
}

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
func DebugInfo(fidles logrus.Fields) (debugInfo logrus.Fields) {
	if pc, file, line, ok := runtime.Caller(1); ok {
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
