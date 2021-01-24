package mediator

import "testing"

func TestMediator(t *testing.T) {
	chatRoom := new(ChatRoom)

	robert := NewUser("Robert", chatRoom)
	john := NewUser("John", chatRoom)

	robert.SendMessage("Hi, John!")
	john.SendMessage("Hi, Robert!")
}
