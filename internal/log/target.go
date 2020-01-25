package log

import (
	"io"
	"os"
)

type Target struct {
	io.Writer
	Level   int
	ToLevel int
	// EnableBuffering bool TODO - do not rely on file writing is fast (perhaps use BufferSize = 0 to disable buffer)
	Formatter    Formatter
	FlushCallback func() error
}

func (target *Target) Format(level int, v ...interface{}) ([]byte, error) {
	if target.Formatter == nil {
		return nil, nil
	}
	return target.Formatter.Format(level, v...)
}

func (target *Target) Flush() error {
	if target.Flush() == nil {
		return nil
	}
	return target.FlushCallback()
}

func RemoveAllTargets(loggers ...*Logger) {
	if loggers == nil {
		loggers = []*Logger{defaultLogger}
	}
	for _, logger := range loggers {
		logger.logTargets = make([]*Target, 0)
	}
}

func AddFileTarget(level int, filename string, loggers ...*Logger) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_CREATE| os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	if loggers == nil {
		loggers = []*Logger{defaultLogger}
	}
	for _, logger := range loggers {
		logger.AddTarget(&Target{
			Writer:        f,
			Level:         level,
			Formatter:     new(FileFormatter),
			FlushCallback: nil,
		})
	}
	return f, nil
}

func AddColorTerminalTarget(level int, loggers ...*Logger) {
	if loggers == nil {
		loggers = []*Logger{defaultLogger}
	}

	if level == LevelOff {
		return
	}

	formatter := new(ColorTerminalFormatter)
	for _, logger := range loggers {
		if level < LevelWarn {
			logger.AddTarget(&Target{
				Writer:        os.Stdout,
				Level:         level,
				ToLevel:       LevelInfo,
				Formatter:     formatter,
				FlushCallback: nil,
			})
		}

		if level < LevelOff {
			logger.AddTarget(&Target{
				Writer:        os.Stderr,
				Level:         level,
				Formatter:     formatter,
				FlushCallback: nil,
			})
		}
	}
}
