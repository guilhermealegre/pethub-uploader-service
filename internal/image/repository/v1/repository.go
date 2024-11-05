package v1

import (
	"bytes"

	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"

	"bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/request"
	domainImage "bitbucket.org/asadventure/be-uploader-service/internal/image/domain/v1"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

type Repository struct {
	app domain.IApp
}

func NewRepository(app domain.IApp) domainImage.IRepository {
	return &Repository{
		app: app,
	}
}

func (r *Repository) Upload(ctx dCtx.IContext, req *request.ImageUploadRequest) (uploader *domainImage.Upload, err error) {
	inputObject := &s3.PutObjectInput{}
	inputObject.Bucket = &r.app.S3().Config().Bucket
	inputObject.Body = bytes.NewReader(req.Image)
	inputObject.ContentType = &req.ContentType

	// Define key and filename
	uploader = &domainImage.Upload{}
	uploader.ImagePath = domainImage.PrepareKeyForImageUpload(req.Section, req.FileName)
	inputObject.Key = &uploader.ImagePath

	// Upload Object to S3 Bucket
	_, err = r.app.S3().Client().PutObject(ctx.RequestContext(), inputObject, func(*s3.Options) {})
	if err != nil {
		return nil, err
	}
	return uploader, nil
}
