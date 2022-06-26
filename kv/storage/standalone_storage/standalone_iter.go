package standalone_storage

import (
	"github.com/Connor1996/badger"
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
)

// standAloneReader iterates over a standAloneReader.
type standAloneIter struct {
	iter *badger.Iterator
}

// Item ...
func (s *standAloneIter) Item() engine_util.DBItem {
	return &standAloneItem{}
}

// Valid ...
func (s *standAloneIter) Valid() bool {
	return s.iter.Valid()
}

// Next ...
func (s *standAloneIter) Next() {
	s.iter.Next()
}

// Seek ...
func (s *standAloneIter) Seek(key []byte) {
	s.iter.Seek(key)
}

// Close ...
func (s *standAloneIter) Close() {
	s.iter.Close()
}
