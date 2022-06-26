package standalone_storage

import (
	"github.com/Connor1996/badger"
)

// standAloneItem is an item returned from standAloneIter.
type standAloneItem struct {
	item *badger.Item
}

// Key ...
func (s *standAloneItem) Key() []byte {
	return s.item.Key()
}

// KeyCopy ...
func (s *standAloneItem) KeyCopy(dst []byte) []byte {
	return s.item.KeyCopy(dst)
}

// Value ...
func (s *standAloneItem) Value() ([]byte, error) {
	return s.item.Value()
}

// ValueSize ...
func (s *standAloneItem) ValueSize() int {
	return s.item.ValueSize()
}

// ValueCopy ...
func (s *standAloneItem) ValueCopy(dst []byte) ([]byte, error) {
	return s.item.ValueCopy(dst)
}
