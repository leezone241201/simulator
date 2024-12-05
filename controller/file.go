package controller

import (
	"github/leezone/simulator/common/response"
	"github/leezone/simulator/common/syserror"
	"github/leezone/simulator/service"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	_, err := ctx.FormFile("uploadFile")
	if err != nil {
		response.Failure(ctx, syserror.UploadFileErrCode, syserror.ErrMap[syserror.UploadFileErrCode])
	}
	service.Upload(ctx)
}
