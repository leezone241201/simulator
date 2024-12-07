package service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github/leezone/simulator/common/constant"
	"github/leezone/simulator/common/logger"
	"github/leezone/simulator/common/syserror"
	"github/leezone/simulator/dao"
	"github/leezone/simulator/svc"
)

func Upload(file *multipart.FileHeader) error {
	// 获取基础文件名,防止../../等格式的路径穿越
	fileName := filepath.Base(file.Filename)

	fileStream, err := file.Open()
	if err != nil {
		logger.Logger.ErrorWithStack("open file error", err, nil)
		return errors.New("读取文件失败")
	}
	defer fileStream.Close()

	filePath := filepath.Join(constant.StaticDir, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		logger.Logger.ErrorWithStack("open file error", err, nil)
		return errors.New("创建文件失败")
	}
	defer f.Close()

	_, err = io.Copy(f, fileStream)
	if err != nil {
		logger.Logger.ErrorWithStack("open file error", err, nil)
		return errors.New("保存文件失败")
	}

	daoFile := &dao.File{Path: filePath}
	err = svc.Svc.FileDB.CreateFile(daoFile)
	if err != nil {
		return errors.New(syserror.ErrMap[syserror.InternalErrCode])
	}

	err = AddServerTmplate(nil, strconv.Itoa(int(daoFile.ID)))
	if err != nil {
		return errors.New(syserror.ErrMap[syserror.InternalErrCode])
	}

	return nil
}
