package v1

import (
	v1Routes "bitbucket.org/asadventure/be-uploader-service/api/v1/http"
	alive "bitbucket.org/asadventure/be-uploader-service/internal/alive/domain/v1"
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

type Controller struct {
	*domain.DefaultController
	app   domain.IApp
	model alive.IModel
}

func NewController(app domain.IApp, model alive.IModel) domain.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		app:               app,
		model:             model,
	}
}

func (c *Controller) Register() {
	v1Routes.Alive.SetRoute(c.app.Http().Router(), c.Get)
	v1Routes.PublicAliveUploader.SetRoute(c.app.Http().Router(), c.GetPublic)
}

/*
	 swagger:route GET /alive alive alive

	 Internal service status check.

		Produces:
		- application/json

		Responses:
		  200: SwaggerAliveResponse
		  400: ErrorResponse
*/
func (c *Controller) Get(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	alive, err := c.model.Get(ctx)
	c.Json(ctx, alive.FromDomainToApi(), err)
}

/*
	 swagger:route GET /p/alive/uploader alive public_alive

	 Public service status check.

		Produces:
		- application/json

		Security:
		  BasicAuth:

		Responses:
		  200: SwaggerPublicAliveResponse
		  400: ErrorResponse
*/
func (c *Controller) GetPublic(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	alive, err := c.model.GetPublic(ctx)
	c.Json(ctx, alive.FromDomainToApi(), err)
}
