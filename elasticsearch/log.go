package elasticsearch

// Log wraps an application log entry for use by the elasticsearch sdk
type Log struct {
	printf func(format string, args ...interface{})
}

// NewLog returns a new Log value
func NewLog(fn func(format string, args ...interface{})) *Log {
	return &Log{
		printf: fn,
	}
}

// Printf calls the underlying log function
func (l *Log) Printf(format string, args ...interface{}) {
	l.printf(format, args...)
}
