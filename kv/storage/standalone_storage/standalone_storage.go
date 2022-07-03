package standalone_storage

import (
	"os"

	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"

	"github.com/Connor1996/badger"
	"github.com/pingcap-incubator/tinykv/kv/config"
	"github.com/pingcap-incubator/tinykv/kv/storage"
	"github.com/pingcap-incubator/tinykv/kv/util"
	"github.com/pingcap-incubator/tinykv/log"
	"github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
)

// StandAloneStorage is an implementation of `Storage` for a single-node TinyKV instance. It does not
// communicate with other nodes and all data is stored locally.
type StandAloneStorage struct {
	db *badger.DB
}

// NewStandAloneStorage creates a new instance of standalone storage.
func NewStandAloneStorage(conf *config.Config) *StandAloneStorage {
	// Creates the directory if it does not exist yet.
	if util.DirExists(conf.DBPath) {
		log.Infof("Going to use an existing data directory at %s", conf.DBPath)
	} else if err := os.MkdirAll(conf.DBPath, os.ModePerm); err != nil {
		log.Fatalf("Unable to create data directory at %s", conf.DBPath)
	}

	// Prepares the options.
	opts := badger.DefaultOptions
	opts.Dir = conf.DBPath
	opts.ValueDir = conf.DBPath

	// Opens a Badger DB.
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Unable to open a badger DB due to %s", err)
	}

	// Returns the new storage instance.
	return &StandAloneStorage{db: db}
}

// Start ...
func (s *StandAloneStorage) Start() error {
	return nil
}

// Stop ...
func (s *StandAloneStorage) Stop() error {
	if err := s.db.Close(); err != nil {
		log.Warnf("Unable to close the badger DB due to %s", err)
		return err
	}
	return nil
}

// Reader ...
func (s *StandAloneStorage) Reader(_ *kvrpcpb.Context) (storage.StorageReader, error) {
	// Creates a transaction for consistent snapshot.
	txn := s.db.NewTransaction(false)

	// Returns a new reader.
	return &standAloneReader{txn: txn}, nil
}

// Write ...
func (s *StandAloneStorage) Write(_ *kvrpcpb.Context, batch []storage.Modify) error {
	return s.db.Update(func(txn *badger.Txn) error {
		// Iterates over each operation.
		for _, modify := range batch {
			keyWithCF := engine_util.KeyWithCF(modify.Cf(), modify.Key())

			// Performs the modification.
			switch modify.Data.(type) {
			case storage.Put:
				if err := txn.Set(keyWithCF, modify.Value()); err != nil {
					return err
				}
			case storage.Delete:
				if err := txn.Delete(keyWithCF); err != nil {
					return err
				}
			default:
				log.Infof("Skip this modification since it is not supported %#v", modify.Data)
			}
		}
		return nil
	})
}
