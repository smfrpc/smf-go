package demo_gen

import (
	"context"

	"github.com/crackcomm/go-smf/example/demo"
)

// SmfStorageClient - SmfStorage Client implementation.
type SmfStorageClient struct{}

// Get - method description.
func (s *SmfStorageClient) Get(ctx context.Context, req []byte) (*demo.Response, error) {
	res, err := s.RawGet(ctx, req)
	if err != nil {
		return nil, err
	}
	return demo.GetRootAsResponse(res, 0), nil
}

// RawGet - Raw method description.
func (s *SmfStorageClient) RawGet(ctx context.Context, req []byte) ([]byte, error) {
	return nil, nil
	// Server would do following:
	//   return s.SmfStorage.Get(ctx, demo.GetRootAsRequest(req, 0))
	// Client does the opposite, Get -> RawGet
	//   // TODO(crackcomm): here we are sending etc.
}
