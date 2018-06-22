package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const keySpaceBits = 160

// Node 代表kademlia网络的参与者
type Node struct {
	ID         NodeID
	KeyValueDB *Database
}

// NodeID 节点ID
type NodeID [keySpaceBits / 8]byte

func (n *Node) generateNodeID() (NodeID, error) {
	b := make([]byte, keySpaceBits/8)
	if _, err := rand.Read(b); err != nil {
		return [keySpaceBits / 8]byte{}, err
	}
	var id [keySpaceBits / 8]byte
	copy(id[:], b[:])
	return id, nil
}

func (id NodeID) getHexString() string {
	h := hex.EncodeToString(id[:])
	fmt.Println("nodeID =", h)
	return h
}

func newNode() (*Node, error) {
	n := &Node{
		KeyValueDB: &Database{Path: "./keyvaluedb"},
	}
	id, err := n.generateNodeID()
	if err != nil {
		return nil, err
	}
	n.ID = id
	return n, nil
}

func main() {
	n, err := newNode()
	if err != nil {
		fmt.Println("newNode error:", err)
		return
	}
	n.ID.getHexString()

	value, err := n.KeyValueDB.get([20]byte{223})
	if err != nil {
		fmt.Println("get error:", err)
	} else {
		fmt.Println("value =", string(value))
	}

	err = n.KeyValueDB.put([20]byte{223}, []byte("helloworld!"))
	if err != nil {
		fmt.Println("put error:", err)
	} else {
		value, _ := n.KeyValueDB.get([20]byte{223})
		fmt.Println("now value =", string(value))
	}
}
