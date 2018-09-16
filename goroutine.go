package logger

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	waitGroup sync.WaitGroup
)

// Entry add some function on logrus.Entry
type Entry struct {
	*logrus.Entry
}

// Wait wait for all logging complete
func Wait() {
	waitGroup.Wait()
}

// logger.Entry wrapper

func WithError(err error) *Entry {
	return &Entry{Entry: logrus.WithError(err)}
}
func WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: logrus.WithField(key, value)}
}
func WithFields(fields logrus.Fields) *Entry {
	return &Entry{Entry: logrus.WithFields(fields)}
}
func WithTime(t time.Time) *Entry {
	return &Entry{Entry: logrus.WithTime(t)}
}

func (logger *Logger) WithError(err error) *Entry {
	return &Entry{Entry: logger.Logger.WithError(err)}
}
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: logger.Logger.WithField(key, value)}
}
func (logger *Logger) WithFields(fields logrus.Fields) *Entry {
	return &Entry{Entry: logger.Logger.WithFields(fields)}
}
func (logger *Logger) WithTime(t time.Time) *Entry {
	return &Entry{Entry: logger.Logger.WithTime(t)}
}

func (entry *Entry) WithError(err error) *Entry {
	return &Entry{Entry: entry.Entry.WithError(err)}
}
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: entry.Entry.WithField(key, value)}
}
func (entry *Entry) WithFields(fields logrus.Fields) *Entry {
	return &Entry{Entry: entry.Entry.WithFields(fields)}
}
func (entry *Entry) WithTime(t time.Time) *Entry {
	return &Entry{Entry: entry.Entry.WithTime(t)}
}

// Go routine wrapper

func (logger *Logger) Debugf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Debugf(format, args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Infof(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Infof(format, args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Printf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Printf(format, args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Warnf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warnf(format, args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Warningf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warningf(format, args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Errorf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Errorf(format, args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Fatalf(format, args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Panicf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Panicf(format, args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Debug(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Debug(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Info(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Info(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Print(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Print(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Warn(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warn(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Warning(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warning(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Error(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Error(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Fatal(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Fatal(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Panic(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Panic(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Debugln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Debugln(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Infoln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Infoln(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Println(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Println(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Warnln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warnln(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Warningln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warningln(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Errorln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Errorln(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Fatalln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Fatalln(args...)
		waitGroup.Done()
	}()
}
func (logger *Logger) Panicln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		logger.Logger.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Panicln(args...)
		waitGroup.Done()
	}()
}

func (entry *Entry) Debug(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Debug(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Print(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Print(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Info(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Info(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Warn(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warn(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Warning(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warning(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Error(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Error(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Fatal(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Fatal(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Panic(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Panic(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Debugf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Debugf(format, args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Infof(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Infof(format, args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Printf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Printf(format, args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Warnf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warnf(format, args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Warningf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warningf(format, args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Errorf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Errorf(format, args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Fatalf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Fatalf(format, args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Panicf(format string, args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Panicf(format, args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Debugln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Debugln(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Infoln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Infoln(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Println(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Println(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Warnln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warnln(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Warningln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Warningln(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Errorln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Errorln(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Fatalln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Fatalln(args...)
		waitGroup.Done()
	}()
}
func (entry *Entry) Panicln(args ...interface{}) {
	waitGroup.Add(1)
	go func() {
		entry.Entry.WithField("_commit_time", time.Now().Format(time.RFC3339Nano)).Panicln(args...)
		waitGroup.Done()
	}()
}
