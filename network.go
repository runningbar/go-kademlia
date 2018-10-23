package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

func sendRequest(remoteAddr string, message []byte) error {
	addr, err := net.ResolveUDPAddr("udp", remoteAddr)
	if err != nil {
		fmt.Println("sendRequest resolveUDPAddr error:", err)
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("sendRequest dialUDP error:", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("sendRequest write error:", err)
		return err
	}
}

func startListen(port int) error {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("resolveUDPAddr error:", err)
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listenUDP error:", err)
		return err
	}
	defer conn.Close()

	for {
		handleMessage(conn)
	}
}

func handleMessage(conn *net.UDPConn) {
	var data []byte
	_, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("readFromUDP error:", err)
		return
	}
	fmt.Println("remoteAddr is:", remoteAddr)

	message := &Message{}
	err = json.Unmarshal(data, message)
	if err != nil {
		fmt.Println("unmarshal json error:", err)
		return
	}
	switch message.Type {
	case PING:
		handlePing(message.Content)
	case STORE:
		handleStore(message.Content)
	case FINDNODE:
		handleFindNode(message.Content)
	case FINDVALUE:
		handleFindValue(message.Content)
	}
}
