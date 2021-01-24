package mediator

import (
	"fmt"
	"time"
)

// ChatRoom ...
type ChatRoom struct{}

// ShowMessage ...
func (*ChatRoom) ShowMessage(user *User, message string) {
	fmt.Printf("%s [ %s ]: %s\n", time.Now(), user.GetName(), message)
}

// User ...
type User struct {
	name     string
	mediator *ChatRoom
}

// NewUser ...
func NewUser(name string, mediator *ChatRoom) *User {
	return &User{name, mediator}
}

// GetName ...
func (u *User) GetName() string {
	return u.name
}

// SendMessage ...
func (u *User) SendMessage(message string) {
	u.mediator.ShowMessage(u, message)
}
