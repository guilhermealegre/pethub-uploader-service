package uploader_service_uploader

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type UploaderClientMock struct {
	mock.Mock
}

func NewUploaderClientMock() *UploaderClientMock {
	return &UploaderClientMock{}
}

func (c *UploaderClientMock) ImageUploadGrpc(ctx context.Context, in *ImageUploadGrpcRequest, opts ...grpc.CallOption) (*ImageUploadGrpcResponse, error) {
	var optsList = []interface{}{ctx, in}
	for _, v := range opts {
		optsList = append(optsList, v)
	}

	args := c.Called(optsList...)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ImageUploadGrpcResponse), args.Error(1)
}
