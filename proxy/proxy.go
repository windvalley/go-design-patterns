package proxy

// Subject subject
type Subject interface {
	DoTask() string
}

type concreteSubject struct{}

func (*concreteSubject) DoTask() string {
	return "SubjectTask"
}

// Proxy proxy
type Proxy struct {
	subject *concreteSubject
}

// DoTask do task
func (p *Proxy) DoTask() string {
	if p.subject == nil {
		p.subject = new(concreteSubject)
	}

	result := "ProxyPreTask-"

	result += p.subject.DoTask()

	result += "-ProxyAfterTask"

	return result
}
