package common

import (
	"fmt"
	"github.com/koding/kite"
)

var (
	ServerPort = map[string]int{
		"user":     6800,
		"passport": 6801,
	}
)

func GetServerPort(serverName string) int {
	serverPort, ok := ServerPort[serverName]
	if !ok {
		panic(fmt.Sprintf("server port not found. severName: %v", serverName))
	}
	return serverPort
}

func GetServerClient(serverName string) *kite.Client {
	k := kite.New(serverName, "1.0.0")
	serverPort, ok := ServerPort[serverName]
	if !ok {
		panic(fmt.Sprintf("server port not found. severName: %v", serverName))
	}
	addr := fmt.Sprintf("http://localhost:%v/kite", serverPort)
	client := k.NewClient(addr)
	client.Dial()
	return client
}
