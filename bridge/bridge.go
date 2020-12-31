package bridge

import "fmt"

// MessageBox abstracter
type MessageBox interface {
	SendMessage(content string)
}

// SendImplementer implementer
type SendImplementer interface {
	Send(content string)
}

// CommonMessageBox common object
type CommonMessageBox struct {
	sender SendImplementer
}

// NewCommonMessageBox instantiate an object
func NewCommonMessageBox(sender SendImplementer) MessageBox {
	return &CommonMessageBox{sender}
}

// SendMessage implement interface
func (c *CommonMessageBox) SendMessage(content string) {
	c.sender.Send(content)
}

// ImportantMessageBox ...
type ImportantMessageBox struct {
	sender SendImplementer
}

// NewImportantMessageBox ...
func NewImportantMessageBox(sender SendImplementer) MessageBox {
	return &ImportantMessageBox{sender}
}

// SendMessage ...
func (c *ImportantMessageBox) SendMessage(content string) {
	content = "[Important]" + content
	c.sender.Send(content)
}

// Mail implementer
type Mail struct{}

// Send implement interface
func (*Mail) Send(content string) {
	fmt.Printf("send content(%s) by mail\n", content)
}

// Phone ...
type Phone struct{}

// Send ...
func (*Phone) Send(content string) {
	fmt.Printf("send content(%s) by phone\n", content)
}
