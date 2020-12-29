package proxy

// Subject ...
type Subject interface {
	DoTask() string
}

type concreteSubject struct{}

// NewConcreteSubject ...
func NewConcreteSubject() Subject {
	return &concreteSubject{}
}

func (*concreteSubject) DoTask() string {
	return "SubjectTask"
}

type proxy struct {
	subject Subject
}

// NewProxy ...
func NewProxy(s Subject) Subject {
	return &proxy{s}
}

func (p *proxy) DoTask() string {
	result := "ProxyPreTask-"

	result += p.subject.DoTask()

	result += "-ProxyAfterTask"

	return result
}
