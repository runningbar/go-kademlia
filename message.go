package main

type messageType int

const (
	// PING ping一个节点
	PING messageType = iota

	// STORE 要求一个节点存储一个<key, value>
	STORE

	// FINDNODE 递归的寻找指定节点
	FINDNODE

	// FINDVALUE 递归的寻找指定<key, value>
	FINDVALUE
)

// Message 代表KAD网络中的元消息
type Message struct {
	Type    messageType `json:"type"`
	Content []byte      `json:"content"`
}
