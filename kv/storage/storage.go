package storage

import (
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
	"github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
)

// Storage represents the internal-facing server part of TinyKV, it handles sending and receiving from other TinyKV nodes.
// As part of that responsibility, it also reads and writes data to disk (or semi-permanent memory).
type Storage interface {
	Start() error
	Stop() error
	Write(ctx *kvrpcpb.Context, batch []Modify) error
	Reader(ctx *kvrpcpb.Context) (StorageReader, error)
}

// StorageReader reads from Storage.
type StorageReader interface {
	// GetCF gets the column family, and returns nil for the value when the key doesn't exist
	GetCF(cf string, key []byte) ([]byte, error)
	IterCF(cf string) engine_util.DBIterator
	Close()
}
