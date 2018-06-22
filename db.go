package main

import "github.com/syndtr/goleveldb/leveldb"

// Database 封装leveldb
type Database struct {
	Path string // db的路径
}

func (db Database) get(key [keySpaceBits / 8]byte) (value []byte, err error) {
	d, err := leveldb.OpenFile(db.Path, nil)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	value, err = d.Get(key[:], nil)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (db Database) put(key [keySpaceBits / 8]byte, value []byte) (err error) {
	d, err := leveldb.OpenFile(db.Path, nil)
	if err != nil {
		return err
	}
	defer d.Close()
	err = d.Put(key[:], value, nil)
	if err != nil {
		return err
	}
	return nil
}
