package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxy(t *testing.T) {
	expected := "SubjectTask"
	subject := NewConcreteSubject()
	result := subject.DoTask()
	assert.Equal(t, expected, result, "they should be equal")

	expected = "ProxyPreTask-SubjectTask-ProxyAfterTask"
	proxy := NewProxy(subject)
	result = proxy.DoTask()
	assert.Equal(t, expected, result, "they should be equal")
}
