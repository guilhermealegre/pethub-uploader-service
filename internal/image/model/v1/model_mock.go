package v1

import (
	"bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/request"
	v1 "bitbucket.org/asadventure/be-uploader-service/internal/image/domain/v1"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/stretchr/testify/mock"
)

func NewModelMock() *ModelMock {
	return &ModelMock{}
}

type ModelMock struct {
	mock.Mock
}

func (m *ModelMock) Upload(ctx dCtx.IContext, req *request.ImageUploadRequest) (*v1.Upload, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*v1.Upload), args.Error(1)
}
