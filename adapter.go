package stdlog

import (
	"log"

	"go.melnyk.org/mlog"
)

type loggerWriter struct {
	logger mlog.Logger
	level  mlog.Level
}

func (l *loggerWriter) Write(msg []byte) (int, error) {
	l.logger.Event(l.level, func(e mlog.Event) {
		e.String("msg", string(msg))
	})
	return len(msg), nil
}

// NewStdLog returns a *log.Logger which writes to the supplied mlog Logger
func NewStdLog(logger mlog.Logger, level mlog.Level) *log.Logger {
	return log.New(&loggerWriter{logger: logger, level: level}, "" /* prefix */, 0 /* flags */)
}

// RedirectStdLog redirects output from the standard library's package-global
// logger to the supplied logger.
//
// It returns a function to restore the original prefix, flags and output.
func RedirectStdLog(logger mlog.Logger, level mlog.Level) func() {
	flags := log.Flags()
	prefix := log.Prefix()
	output := log.Writer()

	log.SetFlags(0)
	log.SetPrefix("")
	log.SetOutput(&loggerWriter{logger: logger, level: level})

	return func() {
		log.SetFlags(flags)
		log.SetPrefix(prefix)
		log.SetOutput(output)
	}
}
