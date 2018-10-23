# go-kademlia

基于`golang`实现的`P2P kademlia DHT`。力求完全符合[论文](http://www.ic.unicamp.br/~bit/ensino/mo809_1s13/papers/P2P/Kademlia-%20A%20Peer-to-Peer%20Information%20System%20Based%20on%20the%20XOR%20Metric%20.pdf)的原本设计

# 备注

* 借鉴了`go-ethereum`中`devp2p`的实现

* 使用[`leveldb`](https://github.com/syndtr/goleveldb)存储`<key, value>`对

* 使用`json`做消息序列化，不用`grpc`和`protobuf`是因为`grpc`不支持`UDP`