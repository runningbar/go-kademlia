package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
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
		KeyValueDB: &Database{Path: "/Users/yanghongxing/ethdev/bootnode-dir/vntdb/<peer.ID 1kHRjixCiatk1W8JythP4VkYsasDt1hF3hFrXzZFo35k6ii>"},
	}

	db, err := leveldb.OpenFile("/Users/yanghongxing/ethdev/datadir1/vntdb", nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		keyByte := iter.Key()
		valueByte := iter.Value()
		fmt.Printf("key = %s, value = %v\n", string(keyByte), string(valueByte))
	}

	id, err := n.generateNodeID()
	if err != nil {
		return nil, err
	}
	n.ID = id
	return n, nil
}

func testStore(addr string) {

}

func main() {
	n, err := newNode()
	if err != nil {
		fmt.Println("newNode error:", err)
		return
	}

	value, err := n.KeyValueDB.get([]byte("/F53C62DFNRWG6MI"))
	if err != nil {
		fmt.Println("get error:", err)
	} else {
		fmt.Println("value =", string(value))
	}

	/*n, err := newNode()
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

	var port = flag.Int("p", 8888, "define kad node udp port")
	var remoteAddr = flag.String("r", "127.0.0.1:8888", "remote udp addr")
	flag.Parse()
	fmt.Printf("port = %d, and remoteAddr = %s\n", *port, *remoteAddr)

	go startListen(*port)
	testStore(*remoteAddr)
	*/
}
