package server

import (
	"context"

	"github.com/pingcap-incubator/tinykv/kv/storage"
	"github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
)

// The functions below are Server's Raw API. (implements TinyKvServer).
// Some helper methods can be found in sever.go in the current directory

// RawGet returns the corresponding Get response based on RawGetRequest's CF and Key fields.
func (server *Server) RawGet(_ context.Context, req *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error) {
	// Gets a reader.
	reader, err := server.storage.Reader(req.GetContext())
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// Reads the key-value pair.
	value, err := reader.GetCF(req.GetCf(), req.GetKey())
	if err != nil {
		return nil, err
	}
	return &kvrpcpb.RawGetResponse{Value: value}, nil
}

// RawPut puts the target data into storage and returns the corresponding response.
func (server *Server) RawPut(_ context.Context, req *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error) {
	// Composes the modification.
	modify := storage.Modify{Data: storage.Put{
		Key:   req.GetKey(),
		Value: req.GetValue(),
		Cf:    req.GetCf(),
	}}

	// Writes to the underlying storage.
	if err := server.storage.Write(req.GetContext(), []storage.Modify{modify}); err != nil {
		return nil, err
	}
	return &kvrpcpb.RawPutResponse{}, nil
}

// RawDelete deletes the target data from storage and returns the corresponding response.
func (server *Server) RawDelete(_ context.Context, req *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error) {
	// Composes the modification.
	modify := storage.Modify{Data: storage.Delete{
		Key: req.GetKey(),
		Cf:  req.GetCf(),
	}}

	// Writes to the underlying storage.
	if err := server.storage.Write(req.GetContext(), []storage.Modify{modify}); err != nil {
		return nil, err
	}
	return &kvrpcpb.RawDeleteResponse{}, nil
}

// RawScan scans the data starting from the start key up to limit. and returns the corresponding result.
func (server *Server) RawScan(_ context.Context, req *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error) {
	// Gets a reader.
	reader, err := server.storage.Reader(req.GetContext())
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// Gets an iterator.
	iter := reader.IterCF(req.GetCf())
	defer iter.Close()

	// Initializes the scan.
	var result []*kvrpcpb.KvPair
	iter.Seek(req.GetStartKey())

	// Performs a scan.
	for i := uint32(0); i < req.GetLimit(); i++ {
		item := iter.Item()

		// Reads the item.
		value, itemErr := item.ValueCopy(nil)
		if itemErr != nil {
			return nil, itemErr
		}
		result = append(result, &kvrpcpb.KvPair{Key: item.KeyCopy(nil), Value: value})

		// Points to the next item.
		iter.Next()
	}
	return &kvrpcpb.RawScanResponse{Kvs: result}, nil
}
