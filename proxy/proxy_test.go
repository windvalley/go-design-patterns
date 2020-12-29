package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxy(t *testing.T) {
	expected := "ProxyPreTask-SubjectTask-ProxyAfterTask"

	proxy := new(Proxy)
	result := proxy.DoTask()

	assert.Equal(t, expected, result, "they should be equal")
}
