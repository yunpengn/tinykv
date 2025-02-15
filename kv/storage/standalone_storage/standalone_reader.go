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
	keyWithCF := engine_util.KeyWithCF(cf, key)

	// Gets the item.
	item, err := s.txn.Get(keyWithCF)
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, nil
		}
		return nil, err
	}

	// Copies the value of the ite.
	return item.ValueCopy(nil)
}

// IterCF ...
func (s *standAloneReader) IterCF(cf string) engine_util.DBIterator {
	return engine_util.NewCFIterator(cf, s.txn)
}

// Close ...
func (s *standAloneReader) Close() {
	s.txn.Discard()
}
