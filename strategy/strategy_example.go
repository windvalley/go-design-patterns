package strategy

import "fmt"

// Logger ...
type Logger interface {
	Debug()
	Info()
}

// LogManager ...
type LogManager struct {
	Logger
}

// NewLogManager ...
func NewLogManager(logger Logger) *LogManager {
	return &LogManager{logger}
}

// FileLogger ...
type FileLogger struct {
}

// Debug ...
func (*FileLogger) Debug() {
	fmt.Println("debug level log to file")
}

// Info ...
func (*FileLogger) Info() {
	fmt.Println("info level log to file")
}

// ElasticsearchLogger ...
type ElasticsearchLogger struct {
}

// Debug ...
func (*ElasticsearchLogger) Debug() {
	fmt.Println("debug level log to elasticsearch")
}

// Info ...
func (*ElasticsearchLogger) Info() {
	fmt.Println("info level log to elasticsearch")
}
