package log

/*
func main() {
    //Close log writer when exit
    defer func(){
        if file, ok := Logger.Out.(*os.File); ok {
            file.Sync()
            file.Close()
        } else if handler, ok := Logger.Out.(io.Closer); ok {
            handler.Close()
        }
    }()
}
*/

var defaultLogger *Logger

func init() {
	defaultLogger = NewLogger()
}

func SetDefaultLogger(logger *Logger) {
	defaultLogger = logger
}

func AddTarget(target *Target) {
	defaultLogger.AddTarget(target)
}

func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

func Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}

func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

func Flush() error {
	return defaultLogger.Flush()
}

func WithTargets(targets ...*Target) error {
	return defaultLogger.WithTargets(targets...)
}
func RemoveAllTargets() {
	defaultLogger.RemoveAllTargets()
}
