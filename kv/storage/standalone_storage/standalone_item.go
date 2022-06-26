package standalone_storage

// standAloneItem is an item returned from standAloneIter.
type standAloneItem struct {
}

// Key ...
func (s *standAloneItem) Key() []byte {
	return nil
}

// KeyCopy ...
func (s *standAloneItem) KeyCopy(dst []byte) []byte {
	return nil
}

// Value ...
func (s *standAloneItem) Value() ([]byte, error) {
	return nil, nil
}

// ValueSize ...
func (s *standAloneItem) ValueSize() int {
	return 0
}

// ValueCopy ...
func (s *standAloneItem) ValueCopy(dst []byte) ([]byte, error) {
	return nil, nil
}
