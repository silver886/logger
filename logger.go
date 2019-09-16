package logger

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sync"

	colorable "github.com/mattn/go-colorable"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Logger add some attributes, base on logrus.Logger.
type Logger struct {
	*logrus.Logger
	wg      sync.WaitGroup
	devInit bool
	Path    string
}

var loggers = map[string]*Logger{}

// New create a new logger.
func New(name string, base logrus.Level, level []logrus.Level, dev bool) (logger *Logger) {
	// If Logger had been created, return nil.
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
	logger = &Logger{
		Logger: logrus.New(),
		Path:   logFileName,
	}

	// Enable color logging in Windows console.
	logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	logger.SetOutput(colorable.NewColorableStdout())

	// Update logger config.
	logger.Config(base, level, dev)

	// Store logger.
	loggers[name] = logger

	return
}

// Get find the logger from storage.
func Get(name string) (logger *Logger) {
	logger = loggers[name]
	return
}

// List shows created loggers.
func List() (keys []string) {
	for k := range loggers {
		keys = append(keys, k)
	}
	return
}

// Has returns true if the logger created before.
func Has(name string) (ok bool) {
	_, ok = loggers[name]
	return
}

// Config updates the config of the logger.
func (logger *Logger) Config(base logrus.Level, level []logrus.Level, dev bool) error {
	// If Logger is nil, return an error.
	if logger == nil {
		return errors.New("Logger cannot be Nil")
	}

	// Set base logging level.
	logger.SetLevel(base)

	// Clear all logging levels.
	logger.Hooks = make(logrus.LevelHooks)

	// Formatter for file logging.
	fileFormatter := &prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true, DisableColors: true, SpacePadding: 64}

	// Add file hook.
	for _, val := range level {
		logger.Hooks.Add(lfshook.NewHook(
			lfshook.PathMap{
				val: logger.Path,
			},
			fileFormatter,
		))
	}

	// Enable developing logging.
	if dev {
		if logger.devInit == false {
			os.RemoveAll("dev_full.log")
			os.RemoveAll("dev_warn.log")
			os.RemoveAll("dev_erro.log")
			os.RemoveAll("dev_crit.log")
			logger.devInit = true
		}

		// Full log file.
		logger.Hooks.Add(lfshook.NewHook(
			lfshook.PathMap{
				logrus.TraceLevel: "dev_full.log",
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

	return nil
}
