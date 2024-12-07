package controller

import (
	"github/leezone/simulator/apistruct"
	"github/leezone/simulator/common/constant"
	"github/leezone/simulator/common/logger"
	"github/leezone/simulator/common/response"
	"github/leezone/simulator/common/syserror"
	"github/leezone/simulator/service"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Upload(ctx *gin.Context) {
	files, err := ctx.MultipartForm()
	if err != nil {
		logger.Logger.ErrorWithStack("get upload file error", err, nil)
		response.Failure(ctx, syserror.UploadFileErrCode, syserror.ErrMap[syserror.UploadFileErrCode])
		return
	}

	uploadFiles := files.File["uploadFile"]
	res := make([]apistruct.UploadResponse, 0, len(uploadFiles))

	for _, file := range uploadFiles {
		if _, exist := constant.AllowUploadSuffix[filepath.Ext(file.Filename)]; !exist {
			logger.Logger.Debug("upload not allowed file", zap.String("filename", file.Filename))
			res = append(res, apistruct.UploadResponse{
				FileName: file.Filename,
				Result:   syserror.ErrMap[syserror.UploadFileNotAllowedErrCode],
			})
			continue
		}

		err = service.Upload(file)
		if err != nil {
			res = append(res, apistruct.UploadResponse{
				FileName: file.Filename,
				Result:   err.Error(),
			})
			continue
		}
	}

	response.Success(ctx, syserror.SuccessCode, "上传成功!", res)
}
