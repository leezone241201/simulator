package controller

import (
	"github/leezone/simulator/apistruct"
	"github/leezone/simulator/common/constant"
	"github/leezone/simulator/common/logger"
	"github/leezone/simulator/common/response"
	"github/leezone/simulator/common/syserror"
	"github/leezone/simulator/service"
	"path/filepath"
	"strconv"

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

	totalChunksStr := ctx.Request.FormValue("totalChunks")
	totalChunks, err := strconv.Atoi(totalChunksStr)
	if err != nil {
		logger.Logger.ErrorWithStack("get totalChunks error", err, totalChunksStr)
		response.Failure(ctx, syserror.ParameterErrCode, syserror.ErrMap[syserror.ParameterErrCode])
		return
	}

	fileName := ctx.Request.FormValue("fileName")
	if fileName == "" {
		logger.Logger.ErrorWithStack("get file name error", err, fileName)
		response.Failure(ctx, syserror.ParameterErrCode, syserror.ErrMap[syserror.ParameterErrCode])
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

		err = service.Upload(file, fileName, totalChunks)
		if err != nil {
			res = append(res, apistruct.UploadResponse{
				FileName: file.Filename,
				Result:   err.Error(),
			})
			continue
		}

		res = append(res, apistruct.UploadResponse{
			FileName: file.Filename,
			Result:   "上传成功",
		})
	}

	response.Success(ctx, syserror.SuccessCode, "上传成功!", res)
}
