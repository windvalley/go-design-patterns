package strategy

import (
	"testing"
)

func TestNewLogManager(t *testing.T) {
	fileLogger := new(FileLogger)
	logManager := NewLogManager(fileLogger)
	logManager.Debug()
	logManager.Info()

	elasticsearchLogger := new(ElasticsearchLogger)
	logManager.Logger = elasticsearchLogger
	logManager.Debug()
	logManager.Info()
}
