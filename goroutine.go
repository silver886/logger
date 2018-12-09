package logger

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Entry add some function on logrus.Entry
type Entry struct {
	*logrus.Entry
	wg *sync.WaitGroup
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
	return &Entry{Entry: logger.Logger.WithError(err), wg: &logger.wg}
}
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: logger.Logger.WithField(key, value), wg: &logger.wg}
}
func (logger *Logger) WithFields(fields logrus.Fields) *Entry {
	return &Entry{Entry: logger.Logger.WithFields(fields), wg: &logger.wg}
}
func (logger *Logger) WithTime(t time.Time) *Entry {
	return &Entry{Entry: logger.Logger.WithTime(t), wg: &logger.wg}
}

func (entry *Entry) WithError(err error) *Entry {
	return &Entry{Entry: entry.Entry.WithError(err), wg: entry.wg}
}
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: entry.Entry.WithField(key, value), wg: entry.wg}
}
func (entry *Entry) WithFields(fields logrus.Fields) *Entry {
	return &Entry{Entry: entry.Entry.WithFields(fields), wg: entry.wg}
}
func (entry *Entry) WithTime(t time.Time) *Entry {
	return &Entry{Entry: entry.Entry.WithTime(t), wg: entry.wg}
}

// Wait wait for all logging complete
func (logger *Logger) Wait() {
	logger.wg.Wait()
}

// Go routine wrapper

func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Tracef(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Debugf(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Infof(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Printf(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Printf(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Warnf(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Warningf(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Errorf(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Fatalf(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Panicf(format, args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Trace(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Trace(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Debug(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Debug(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Info(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Info(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Print(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Print(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Warn(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Warn(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Warning(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Warning(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Error(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Error(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Fatal(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Fatal(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Panic(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Panic(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Traceln(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Traceln(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Debugln(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Debugln(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Infoln(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Infoln(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Println(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Println(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Warnln(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Warnln(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Warningln(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Warningln(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Errorln(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Errorln(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Fatalln(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Fatalln(args...)
		logger.wg.Done()
	}()
}
func (logger *Logger) Panicln(args ...interface{}) {
	logger.wg.Add(1)
	go func() {
		logger.Logger.WithTime(time.Now()).Panicln(args...)
		logger.wg.Done()
	}()
}

func (entry *Entry) Trace(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Trace(args...)
		} else {
			entry.Entry.Trace(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Debug(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Debug(args...)
		} else {
			entry.Entry.Debug(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Print(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Print(args...)
		} else {
			entry.Entry.Print(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Info(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Info(args...)
		} else {
			entry.Entry.Info(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Warn(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Warn(args...)
		} else {
			entry.Entry.Warn(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Warning(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Warning(args...)
		} else {
			entry.Entry.Warning(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Error(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Error(args...)
		} else {
			entry.Entry.Error(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Fatal(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Fatal(args...)
		} else {
			entry.Entry.Fatal(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Panic(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Panic(args...)
		} else {
			entry.Entry.Panic(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Debugf(format, args...)
		} else {
			entry.Entry.Debugf(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Tracef(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Tracef(format, args...)
		} else {
			entry.Entry.Tracef(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Infof(format, args...)
		} else {
			entry.Entry.Infof(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Printf(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Printf(format, args...)
		} else {
			entry.Entry.Printf(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Warnf(format, args...)
		} else {
			entry.Entry.Warnf(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Warningf(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Warningf(format, args...)
		} else {
			entry.Entry.Warningf(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Errorf(format, args...)
		} else {
			entry.Entry.Errorf(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Fatalf(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Fatalf(format, args...)
		} else {
			entry.Entry.Fatalf(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Panicf(format string, args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Panicf(format, args...)
		} else {
			entry.Entry.Panicf(format, args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Traceln(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Traceln(args...)
		} else {
			entry.Entry.Traceln(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Debugln(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Debugln(args...)
		} else {
			entry.Entry.Debugln(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Infoln(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Infoln(args...)
		} else {
			entry.Entry.Infoln(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Println(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Println(args...)
		} else {
			entry.Entry.Println(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Warnln(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Warnln(args...)
		} else {
			entry.Entry.Warnln(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Warningln(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Warningln(args...)
		} else {
			entry.Entry.Warningln(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Errorln(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Errorln(args...)
		} else {
			entry.Entry.Errorln(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Fatalln(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Fatalln(args...)
		} else {
			entry.Entry.Fatalln(args...)
		}
		entry.wg.Done()
	}()
}
func (entry *Entry) Panicln(args ...interface{}) {
	entry.wg.Add(1)
	go func() {
		if entry.Time.IsZero() {
			entry.Entry.WithTime(time.Now()).Panicln(args...)
		} else {
			entry.Entry.Panicln(args...)
		}
		entry.wg.Done()
	}()
}
