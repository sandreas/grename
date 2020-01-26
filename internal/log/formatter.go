package log

import (
	"fmt"
	"time"
)

/*
type MyJSONFormatter struct {
}

log.SetFormatter(new(MyJSONFormatter))

func (f *MyJSONFormatter) Format(entry *Entry) ([]byte, error) {
  // Note this doesn't include Time, Level and Message which are available on
  // the Entry. Consult `godoc` on information about those fields or read the
  // source of the official loggers.
  serialized, err := json.Marshal(entry.Data)
    if err != nil {
      return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
    }
  return append(serialized, '\n'), nil
}
*/

// https://github.com/apsdehal/go-logger/blob/master/logger.go
// https://stackoverflow.com/questions/2616906/how-do-i-output-coloured-text-to-a-linux-terminal
const (
	TerminalColorBlack = iota + 30
	TerminalColorRed
	TerminalColorGreen
	TerminalColorYellow
	TerminalColorBlue
	TerminalColorMagenta
	TerminalColorCyan
	TerminalColorWhite
)

const (
	TerminalPrefixBrighten = "1;"
)

type Formatter interface {
	Format(level int, v ...interface{}) ([]byte, error)
}

/*
type RawFormatter struct {
}

func (f *RawFormatter) Format(level int, v ...interface{}) ([]byte, error){
	return  []byte(fmt.Sprint(v...)), nil
}
*/

type FileFormatter struct {
}

func (f *FileFormatter) Format(level int, v ...interface{}) ([]byte, error) {
	logLine := fmt.Sprintf("%s [%s] %s", time.Now().Format(time.RFC3339), levelToString(level), fmt.Sprintln(v...))
	return []byte(logLine), nil
}

type ColorTerminalFormatter struct {
}

func (f *ColorTerminalFormatter) Format(level int, v ...interface{}) ([]byte, error) {
	var color int
	var brightenPrefix string
	switch level {
	case LevelTrace:
		color = TerminalColorBlack
		brightenPrefix = TerminalPrefixBrighten
	case LevelDebug:
		color = TerminalColorWhite
	case LevelWarn:
		color = TerminalColorYellow
	case LevelError:
		color = TerminalColorRed
		brightenPrefix = TerminalPrefixBrighten
	case LevelFatal:
		color = TerminalColorMagenta
	case LevelInfo:
	default:
		color = TerminalColorWhite
		brightenPrefix = TerminalPrefixBrighten
	}

	return []byte(fmt.Sprintf("\033[%s%dm", brightenPrefix, color) + fmt.Sprintln(v...) + "\033[0m"), nil
}
