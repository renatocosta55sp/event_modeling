package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	coreError "github.org/eventmodeling/ecommerce/pkg/support/errors"
	response "github.org/eventmodeling/ecommerce/pkg/support/server/model"
)

type Response struct{}

func (Response) ResponseNoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (Response) ResponseNotFound(ctx *gin.Context) {
	ctx.Status(http.StatusNotFound)
}

func (Response) ResponseOK(ctx *gin.Context, bodyResponse any) {
	ctx.JSON(http.StatusOK, bodyResponse)
}

func (Response) ResponseListOk(ctx *gin.Context, bodyResponse any, total int64) {
	var list response.List

	list.Pagination = response.Pagination{Total: total}
	list.Data = bodyResponse

	ctx.JSON(http.StatusOK, list)
}

func (r Response) ResponseCreated(ctx *gin.Context, ID uuid.UUID) {
	bodyResponse := response.Created{ID: ID}

	r.ResponseCreatedBody(ctx, bodyResponse)
}

func (Response) ResponseCreatedBody(ctx *gin.Context, bodyResponse any) {
	ctx.JSON(http.StatusCreated, bodyResponse)
}

func (Response) ResponseCreatedInt64(ctx *gin.Context, ID int64) {
	bodyResponse := response.CreatedInt64{ID: ID}

	ctx.JSON(http.StatusCreated, bodyResponse)
}

func (Response) ResponseBadRequest(ctx *gin.Context, err interface{}) {
	errOut := coreError.ConvertTo(err)

	ctx.AbortWithStatusJSON(http.StatusBadRequest, errOut)
}

func (Response) ResponseBadRequestBody(ctx *gin.Context, bodyResponse any) {
	ctx.JSON(http.StatusBadRequest, bodyResponse)
}

func (Response) ResponseUnauthorized(ctx *gin.Context) {
	bodyResponse := coreError.New("error.login.unauthorized")
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, bodyResponse)
}

func (Response) ResponseConflict(ctx *gin.Context, bodyResponse any) {
	ctx.AbortWithStatusJSON(http.StatusConflict, bodyResponse)
}
