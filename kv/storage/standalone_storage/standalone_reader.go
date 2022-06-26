package standalone_storage

import (
	"github.com/Connor1996/badger"
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
	"github.com/pingcap-incubator/tinykv/log"
)

// standAloneReader reads from StandAloneStorage.
type standAloneReader struct {
	txn *badger.Txn
}

// GetCF ...
func (s *standAloneReader) GetCF(_ string, key []byte) ([]byte, error) {
	// Gets the item.
	item, err := s.txn.Get(key)
	if err != nil {
		log.Warnf("Unable to get the item for key %#v", key)
		return nil, err
	}

	// Copies the value of the ite.
	return item.ValueCopy(nil)
}

// IterCF ...
func (s *standAloneReader) IterCF(cf string) engine_util.DBIterator {
	return &standAloneIter{iter: s.txn.NewIterator(badger.DefaultIteratorOptions)}
}

// Close ...
func (s *standAloneReader) Close() {
	s.txn.Discard()
}
