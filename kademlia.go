package main

import "fmt"

func handlePing(content []byte) {
	fmt.Println("handle ping rpc request here")
}

func handleStore(content []byte) {
	fmt.Println("handle store rpc request here")
}

func handleFindNode(content []byte) {
	fmt.Println("handle findNode rpc request here")
}

func handleFindValue(content []byte) {
	fmt.Println("handle findValue rpc request here")
}
