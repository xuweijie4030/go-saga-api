package controller

import (
	"gosagaapi/app/agreement"
	"gosagaapi/app/logic"
	"github.com/carefreex-io/xhttp"
	"github.com/gin-gonic/gin"
	"sync"
)

type AuthController struct {
}

var (
	authController     *AuthController
	authControllerOnce sync.Once
)

func NewAuthController() *AuthController {
	if authController == nil {
		authControllerOnce.Do(func() {
			authController = &AuthController{}
		})
	}

	return authController
}

func (c *AuthController) Login(ctx *gin.Context) {
	params := agreement.AuthRequest{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		xhttp.BadRequestResponse(ctx, err)
		return
	}

	response, err := logic.NewAuthLogic().Auth(ctx, params)
	if err != nil {
		xhttp.UnauthorizedResponse(ctx, err)
		return
	}

	xhttp.SuccessResponse(ctx, response)
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	params := agreement.RefreshRequest{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		xhttp.BadRequestResponse(ctx, err)
		return
	}

	response, err := logic.NewAuthLogic().Refresh(ctx)
	if err != nil {
		xhttp.UnauthorizedResponse(ctx, err)
		return
	}

	xhttp.SuccessResponse(ctx, response)
}
