package demo_gen

import (
	"context"

	"github.com/crackcomm/go-smf/example/demo"
	"github.com/crackcomm/go-smf/rpc"
)

// Storage - Storage service interface.
type Storage interface {
	// Get - method description.
	Get(context.Context, *demo.Request) ([]byte, error)
}

// StorageService - Storage service implementation.
type StorageService struct {
	Storage
}

// NewStorageService - Creates a new Storage service.
func NewStorageService(s Storage) *StorageService {
	return &StorageService{Storage: s}
}

// ServiceName - Returns smf service name.
func (s *StorageService) ServiceName() string {
	return "SmfStorage"
}

// ServiceID - Returns smf service ID.
func (s *StorageService) ServiceID() uint32 {
	return 212494116
}

// MethodHandle - Returns method handle for request ID.
// The handle is nil if the request ID is not recognized.
func (s *StorageService) MethodHandle(id uint32) rpc.RawHandle {
	switch id {
	case 212494116 ^ 1719559449:
		return s.RawGet
	default:
		return nil
	}
}

// RawGet - Calls underlying storage interface by casting request to *demo.Request.
func (s *StorageService) RawGet(ctx context.Context, req []byte) ([]byte, error) {
	return s.Storage.Get(ctx, demo.GetRootAsRequest(req, 0))
}
