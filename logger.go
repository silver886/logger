// Package logger provides both console and file system logging.
package logger

import (
	"io/ioutil"
	"log"
	"os"

	colorable "github.com/mattn/go-colorable"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var logger *logrus.Logger

// New create a new logger.
func New(filename string, dev, force bool) *logrus.Logger {
	// If Logger had been created and not force create, return the original Logger
	if logger != nil && !force {
		return logger
	}

	// Create log file in temp folder.
	logFile, err := ioutil.TempFile("", filename+".*.log")
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

	return logger
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
