package v1

import (
	"github.com/gin-gonic/gin"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
)

type IController interface {
	GetPublic(ctx *gin.Context)
}

type IModel interface {
	Get(ctx dCtx.IContext) (*Alive, error)
	GetPublic(ctx dCtx.IContext) (*PublicAlive, error)
}
