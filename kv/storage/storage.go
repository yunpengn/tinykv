package storage

import (
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
	"github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
)

// Storage represents the internal-facing server part of TinyKV, it handles sending and receiving from other TinyKV nodes.
// As part of that responsibility, it also reads and writes data to disk (or semi-permanent memory).
type Storage interface {
	// Start starts the Storage.
	Start() error

	// Stop stops the Storage.
	Stop() error

	// Write performs a batch of modifications to the Storage.
	Write(ctx *kvrpcpb.Context, batch []Modify) error

	// Reader returns a StorageReader for the Storage.
	Reader(ctx *kvrpcpb.Context) (StorageReader, error)
}

// StorageReader reads from Storage.
type StorageReader interface {
	// GetCF gets the value for the given key under the given column family, and returns nil for the value when the key doesn't exist
	GetCF(cf string, key []byte) ([]byte, error)

	// IterCF iterates over the key-value pairs in the given column family.
	IterCF(cf string) engine_util.DBIterator

	// Close closes the StorageReader.
	Close()
}
