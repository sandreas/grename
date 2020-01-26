package log

import (
	"io"
)

func NewFileWrapperTarget(filename string, level, toLevel int) *Target {
	if filename == "" {
		return nil
	}
	wrapper := NewFileWrapper(filename)
	return &Target{
		Writer:        wrapper,
		Level:         level,
		ToLevel:       toLevel,
		Formatter:     new(FileFormatter),
		FlushCallback: wrapper.Flush,
	}
}

func NewColorTerminalTarget(writer io.Writer, level, toLevel int) *Target {
	return &Target{
		Writer:        writer,
		Level:         level,
		ToLevel:       toLevel,
		Formatter:     new(ColorTerminalFormatter),
		FlushCallback: nil,
	}
}

type Target struct {
	io.Writer
	Level         int
	ToLevel       int
	Formatter     Formatter
	FlushCallback func() error
}

func (target *Target) Format(level int, v ...interface{}) ([]byte, error) {
	if target.Formatter == nil {
		return nil, nil
	}
	return target.Formatter.Format(level, v...)
}

func (target *Target) Flush() error {
	if target.FlushCallback == nil {
		return nil
	}
	return target.FlushCallback()
}
