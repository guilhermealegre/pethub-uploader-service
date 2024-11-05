package v1

import (
	v1Routes "bitbucket.org/asadventure/be-uploader-service/api/v1/http"
	"bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/request"
	v1 "bitbucket.org/asadventure/be-uploader-service/internal/image/domain/v1"
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

type Controller struct {
	*domain.DefaultController
	model v1.IModel
}

func NewController(app domain.IApp, model v1.IModel) v1.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	engine := c.App().Http().Router()
	v1Routes.ImageUpload.SetRoute(engine, c.Upload)
	v1Routes.ImagePUpload.SetRoute(engine, c.UploadPublic)
}

/*
	 swagger:route POST /uploader/image/upload uploader ImageUploadRequest

	 Uploads an image.

		Produces:
		- application/json

		Responses:
		  200: ImageResponse
		  400: ErrorResponse
*/
func (c *Controller) Upload(gCtx *gin.Context) {
	c.upload(gCtx)
}

/*
	 swagger:route POST /p/uploader/image/upload uploader PublicImageUploadRequest

	 Uploads an image.

		Produces:
		- application/json

		Responses:
		  200: ImageResponse
		  400: ErrorResponse
*/
func (c *Controller) UploadPublic(gCtx *gin.Context) {
	c.upload(gCtx)
}

func (c *Controller) upload(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	var err error
	req := &request.ImageUploadRequest{}
	err = req.LoadFromApi(ctx)
	if err != nil {
		c.Json(ctx, nil, err)
		return
	}
	// upload file
	obj, err := c.model.Upload(ctx, req)
	// convert and send to client
	c.Json(ctx, obj.FromDomainToApi(), err)
}
