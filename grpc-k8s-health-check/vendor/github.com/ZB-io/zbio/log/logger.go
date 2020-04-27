package log

var log, _ = newLogger()

type Fields map[string]interface{}

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
	Fatal = "fatal"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Println(args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	WithFields(keyValues Fields) Logger
	Printf(format string, v ...interface{})
}

type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

func newLogger() (Logger, error) {
	config := Configuration{
		EnableConsole:     true,
		ConsoleLevel:      Info,
		ConsoleJSONFormat: false,
		EnableFile:        false,
		FileLevel:         Info,
		FileJSONFormat:    true,
		FileLocation:      "log.log",
	}
	logger, err := newZapLogger(config)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
func Println(v ...interface{}) {
	log.Println(v...)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func WithFields(keyValues Fields) {
	log = log.WithFields(keyValues)
}
