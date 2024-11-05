package v1

import (
	"bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/request"
	domainImage "bitbucket.org/asadventure/be-uploader-service/internal/image/domain/v1"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
)

type Model struct {
	app        domain.IApp
	repository domainImage.IRepository
}

func NewModel(app domain.IApp, repository domainImage.IRepository) domainImage.IModel {
	return &Model{
		app:        app,
		repository: repository,
	}
}

func (m *Model) Upload(ctx dCtx.IContext, req *request.ImageUploadRequest) (*domainImage.Upload, error) {
	return m.repository.Upload(ctx, req)
}
