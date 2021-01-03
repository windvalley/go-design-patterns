package facade

import "fmt"

// ModuleFoo ...
type ModuleFoo interface {
	Run()
}

// ModuleBar ...
type ModuleBar interface {
	Fly()
}

// Facade ...
type Facade struct {
	foo ModuleFoo
	bar ModuleBar
}

// NewFacade ...
func NewFacade() *Facade {
	return &Facade{
		foo: NewConcreteModuleFoo(),
		bar: NewConcreteModuleBar(),
	}
}

// API1 ...
func (f *Facade) API1() {
	f.foo.Run()
}

// API2 ...
func (f *Facade) API2() {
	f.bar.Fly()
}

// API3 ...
func (f *Facade) API3() {
	f.foo.Run()
	f.bar.Fly()
}

type concreteModuleFoo struct{}

// NewConcreteModuleFoo ...
func NewConcreteModuleFoo() ModuleFoo {
	return &concreteModuleFoo{}
}

func (*concreteModuleFoo) Run() {
	fmt.Println("running...")
}

type concreteModuleBar struct{}

// NewConcreteModuleBar ...
func NewConcreteModuleBar() ModuleBar {
	return &concreteModuleBar{}
}

func (*concreteModuleBar) Fly() {
	fmt.Println("flying...")
}
