package log

import "fmt"

const (
	LevelAll = int(^uint(0) >> 1)
	LevelTrace = 600
	LevelDebug = 500
	LevelInfo = 400
	LevelWarn = 300
	LevelError = 200
	LevelFatal = 100
	LevelOff = 0
)

func NewLogger() *Logger {
	return &Logger{}
}

type Logger struct  {
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
		if target.Level <= level && target.ToLevel >= level  {
			if content, err := target.Formatter.Format(level, v...); err == nil {
				target.Write(content)
			}
		}
	}
}

func (logger *Logger) AddTarget(target *Target) {
	if target.ToLevel < target.Level{
		target.ToLevel = LevelAll
	}
	logger.logTargets = append(logger.logTargets,target)
}


func (logger *Logger) Debug(v ...interface{})  {
	logger.logInternalLn(LevelInfo, v...)
}

func (logger *Logger) Info(v ...interface{})  {
	logger.logInternalLn(LevelInfo, v...)
}

func (logger *Logger) Warn(v ...interface{})  {
	logger.logInternalLn(LevelInfo, v...)
}

func (logger *Logger) Error(v ...interface{})  {
	logger.logInternalLn(LevelInfo, v...)
}

func (logger *Logger) Fatal(v ...interface{})  {
	logger.logInternalLn(LevelInfo, v...)
}

func (logger *Logger) Debugf(format string, v ...interface{})  {
	logger.logfInternal(LevelDebug,format, v...)
}

func (logger *Logger) Infof(format string, v ...interface{})  {
	logger.logfInternal(LevelInfo, format, v...)
}

func (logger *Logger) Warnf(format string, v ...interface{})  {
	logger.logfInternal(LevelWarn, format, v...)
}

func (logger *Logger) Errorf(format string, v ...interface{})  {
	logger.logfInternal(LevelError, format, v...)
}

func (logger *Logger) Fatalf(format string, v ...interface{})  {
	logger.logfInternal(LevelFatal, format, v...)
}

func (logger *Logger) Flush() []error {
	errors := make([]error, 0)
	for _, target := range logger.logTargets {
		if err := target.Flush(); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
