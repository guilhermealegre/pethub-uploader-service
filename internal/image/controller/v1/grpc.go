package v1

import (
	"context"

	uploader "bitbucket.org/asadventure/be-uploader-service/api/v1/grpc/uploader_service_uploader"

	requestModel "bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/request"
	v1 "bitbucket.org/asadventure/be-uploader-service/internal/image/domain/v1"
	"github.com/gin-gonic/gin"
	domainContext "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

type GrpcController struct {
	*domain.DefaultController
	app   domain.IApp
	model v1.IModel
	uploader.UnimplementedUploaderServer
}

func NewGrpcController(app domain.IApp, model v1.IModel) *GrpcController {
	return &GrpcController{app: app, model: model}
}

func (g *GrpcController) Register() {
	server, _ := g.app.Grpc().GetServer()
	uploader.RegisterUploaderServer(server, g)
}

func (g *GrpcController) ImageUploadGrpc(ctx context.Context, grpcRequest *uploader.ImageUploadGrpcRequest) (*uploader.ImageUploadGrpcResponse, error) {
	// New context
	newCtx := domainContext.NewContext(&gin.Context{})
	// load domain request
	domainRequest := &requestModel.ImageUploadRequest{}
	domainRequest.LoadFromGrpc(grpcRequest)
	// Upload File
	resp, err := g.model.Upload(newCtx, domainRequest)
	if err != nil {
		return nil, err
	}
	// Convert response and Return
	return resp.FromDomainToGrpc(), err
}
