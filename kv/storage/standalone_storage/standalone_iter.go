package standalone_storage

import (
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
)

// standAloneReader iterates over a standAloneReader.
type standAloneIter struct {
}

func (s standAloneIter) Item() engine_util.DBItem {
	return nil
}

func (s standAloneIter) Valid() bool {
	return false
}

func (s standAloneIter) Next() {
}

func (s standAloneIter) Seek(_ []byte) {
}

func (s standAloneIter) Close() {
}
