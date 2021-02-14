package buildoptions

import (
	"fmt"
	"testing"
	"time"
)

func TestBuildOptions(t *testing.T) {
	serverBuilder := new(ServerBuilder)

	// normal case:
	server, _ := serverBuilder.Create("127.0.0.1", 8080).
		WithProtocol("udp").
		WithMaxConn(200).
		WithTimeout(10 * time.Second).
		Build()

	fmt.Println(server)

	// error case:
	_, err := serverBuilder.Create("127.0.0.1", 8080).
		WithProtocol("udp").
		WithMaxConn(20480).
		WithTimeout(10 * time.Second).
		Build()

	fmt.Println(err)
}
