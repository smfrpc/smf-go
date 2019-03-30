package demo_gen

import (
	"context"

	"github.com/smfrpc/smf-go/example/demo"
	"github.com/smfrpc/smf-go/src/smf"
)

// SmfStorageClient - SmfStorage Client implementation.
type SmfStorageClient struct {
	*smf.Client
}

// NewSmfStorageClient - Creates new SmfStorage client.
func NewSmfStorageClient(client *smf.Client) *SmfStorageClient {
	return &SmfStorageClient{Client: client}
}

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
	return s.Client.SendRecv(req, 212494116^1719559449)
}
