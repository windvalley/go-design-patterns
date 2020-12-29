package proxy

// Subject ...
type Subject interface {
	DoTask() string
}

// ConcreteSubject ...
type ConcreteSubject struct{}

// NewConcreteSubject ...
func NewConcreteSubject() Subject {
	return &ConcreteSubject{}
}

// DoTask ...
func (*ConcreteSubject) DoTask() string {
	return "SubjectTask"
}

// Proxy ...
type Proxy struct {
	subject Subject
}

// NewProxy ...
func NewProxy(s Subject) Subject {
	return &Proxy{s}
}

// DoTask ...
func (p *Proxy) DoTask() string {
	result := "ProxyPreTask-"

	result += p.subject.DoTask()

	result += "-ProxyAfterTask"

	return result
}
