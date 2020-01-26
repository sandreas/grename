package log

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	LevelAll   = int(^uint(0) >> 1)
	LevelTrace = 600
	LevelDebug = 500
	LevelInfo  = 400
	LevelWarn  = 300
	LevelError = 200
	LevelFatal = 100
	LevelOff   = 0
)

func levelToString(level int) string {
	switch level {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO "
	case LevelWarn:
		return "WARN "
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "-----"
	}
}

func NewLogger() *Logger {
	return &Logger{}
}

type Logger struct {
	logTargets []*Target
}

func (logger *Logger) logInternalLn(level int, v ...interface{}) {
	logger.logInternal(level, v...)
}
func (logger *Logger) logfInternal(level int, format string, v ...interface{}) {
	logger.logInternal(level, fmt.Sprintf(format, v...))
}

func (logger *Logger) logInternal(level int, v ...interface{}) {
	for _, target := range logger.logTargets {
		if target.ToLevel <= level && target.Level >= level {
			if content, err := target.Formatter.Format(level, v...); err == nil {
				target.Write(content)
			}
		}
	}
}

func (logger *Logger) AddTarget(target *Target) {
	if target == nil {
		return
	}
	logger.logTargets = append(logger.logTargets, target)
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.logInternalLn(LevelDebug, v...)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.logInternalLn(LevelInfo, v...)
}

func (logger *Logger) Warn(v ...interface{}) {
	logger.logInternalLn(LevelWarn, v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.logInternalLn(LevelError, v...)
}

func (logger *Logger) Fatal(v ...interface{}) {
	logger.logInternalLn(LevelFatal, v...)
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.logfInternal(LevelDebug, format, v...)
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.logfInternal(LevelInfo, format, v...)
}

func (logger *Logger) Warnf(format string, v ...interface{}) {
	logger.logfInternal(LevelWarn, format, v...)
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.logfInternal(LevelError, format, v...)
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.logfInternal(LevelFatal, format, v...)
}

func (logger *Logger) Flush() error {
	var err error
	for _, target := range logger.logTargets {
		err = target.Flush()
	}
	return err
}

func (logger *Logger) WithTargets(targets ...*Target) error {
	nilTargets := make([]string, 0)
	for index, target := range targets {
		if target == nil {
			nilTargets = append(nilTargets, strconv.Itoa(index))
		} else {
			logger.AddTarget(target)
		}
	}

	if len(nilTargets) > 0 {
		return fmt.Errorf("nil targets found - this means that one of the targets provided as parameter could not be created, missing indexes are: %s", strings.Join(nilTargets, ", "))
	}
	return nil
}

func (logger *Logger) RemoveAllTargets() {
	logger.logTargets = make([]*Target, 0)
}
