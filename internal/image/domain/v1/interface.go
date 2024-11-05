package v1

import (
	"bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/request"
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
)

type IController interface {
	domain.IController
	Upload(ctx *gin.Context)
}

type IModel interface {
	Upload(ctx dCtx.IContext, req *request.ImageUploadRequest) (*Upload, error)
}

type IRepository interface {
	Upload(ctx dCtx.IContext, req *request.ImageUploadRequest) (*Upload, error)
}
