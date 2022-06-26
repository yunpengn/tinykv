package standalone_storage

import (
	"github.com/Connor1996/badger"
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
)

// standAloneReader reads from StandAloneStorage.
type standAloneReader struct {
	txn *badger.Txn
}

// GetCF ...
func (s *standAloneReader) GetCF(cf string, key []byte) ([]byte, error) {
	return nil, nil
}

// IterCF ...
func (s *standAloneReader) IterCF(cf string) engine_util.DBIterator {
	return &standAloneIter{}
}

// Close ...
func (s *standAloneReader) Close() {
	s.txn.Discard()
}
