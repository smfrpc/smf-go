package demo_gen

import (
	"context"

	"github.com/crackcomm/go-smf/example/demo"
)

// StorageClient - Storage Client implementation.
type StorageClient struct {
	Storage
}

// Get -
func (s *StorageClient) Get(ctx context.Context, req []byte) (*demo.Response, error) {
	res, err := s.RawGet(ctx, req)
	if err != nil {
		return nil, err
	}
	return demo.GetRootAsResponse(res, 0), nil
}

// RawGet -
func (s *StorageClient) RawGet(ctx context.Context, req []byte) ([]byte, error) {
	return nil, nil
}
