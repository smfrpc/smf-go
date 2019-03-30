package demo_gen

import (
	"context"

	"github.com/smfrpc/smf-go/example/demo"
	"github.com/smfrpc/smf-go/src/smf"
)

// SmfStorage - SmfStorage service interface.
type SmfStorage interface {
	// Get - method description.
	Get(context.Context, *demo.Request) ([]byte, error)
}

// SmfStorageService - SmfStorage service implementation.
type SmfStorageService struct {
	SmfStorage
}

// NewSmfStorageService - Creates a new SmfStorage service.
func NewSmfStorageService(s SmfStorage) *SmfStorageService {
	return &SmfStorageService{SmfStorage: s}
}

// ServiceName - Returns smf service name.
func (s *SmfStorageService) ServiceName() string {
	return "SmfStorage"
}

// ServiceID - Returns smf service ID.
func (s *SmfStorageService) ServiceID() uint32 {
	return 212494116
}

// MethodHandle - Returns method handle for request ID.
// The handle is nil if the request ID is not recognized.
func (s *SmfStorageService) MethodHandle(id uint32) smf.RawHandle {
	switch id {
	case 212494116 ^ 1719559449:
		return s.RawGet
	default:
		return nil
	}
}

// RawGet - Calls underlying storage interface by casting request to *demo.Request.
func (s *SmfStorageService) RawGet(ctx context.Context, req []byte) ([]byte, error) {
	return s.SmfStorage.Get(ctx, demo.GetRootAsRequest(req, 0))
}
