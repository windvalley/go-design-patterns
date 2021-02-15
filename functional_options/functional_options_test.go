package functionaloptions

import (
	"fmt"
	"testing"
	"time"
)

func TestFunctionalOptions(t *testing.T) {
	server1, _ := NewServer("localhost", 80)
	server2, _ := NewServer("localhost", 80, WithProtocol("udp"), WithTimeout(20*time.Second), WithMaxConns(3000))
	_, err := NewServer("localhost", 80, WithProtocol("udp"), WithTimeout(20*time.Second), WithMaxConns(20000))

	// default options
	fmt.Println(server1)

	// override default options
	fmt.Println(server2)

	// error case
	fmt.Println(err)
}
