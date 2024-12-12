package service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github/leezone/simulator/common/constant"
	"github/leezone/simulator/common/logger"
	"github/leezone/simulator/common/syserror"
	"github/leezone/simulator/dao"
	"github/leezone/simulator/svc"
)

var unMergeFilesMap map[string]map[string]struct{}

var ErrUpload = errors.New("file processing failed")
var ErrSaveFile = errors.New("save file error")
var ErrMergeFile = errors.New("merge files error")

func saveFile(file *multipart.FileHeader, dstDir string) (string, error) {
	// 获取基础文件名,防止../../等格式的路径穿越
	fileName := filepath.Base(file.Filename)

	var err error
	defer func() {
		if err != nil {
			logger.Logger.ErrorWithStack(ErrSaveFile.Error(), err, nil)
			err = ErrUpload
		}
	}()

	fileStream, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileStream.Close()

	filePath := filepath.Join(dstDir, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, fileStream)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func mergeFiles(dstDirName, dstFileName string, files []string) (string, error) {
	var tempFile *os.File
	var err error
	defer func() {
		if err != nil {
			logger.Logger.ErrorWithStack(ErrMergeFile.Error(), err, files)
			err = ErrUpload
		}
	}()

	dstFilePath := filepath.Join(dstDirName, dstFileName)
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	for _, fileName := range files {
		tempFile, err = os.Open(fileName)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(dstFile, tempFile)
		if err != nil {
			return "", err
		}
	}

	return dstFilePath, nil
}

func sortFiles(fileMap map[string]struct{}) []string {
	sortedFileNames := make([]string, 0, len(fileMap))
	for fileName := range fileMap {
		sortedFileNames = append(sortedFileNames, fileName)
	}

	sort.Strings(sortedFileNames)
	return sortedFileNames
}

func Upload(file *multipart.FileHeader, dstFileName string, chunks int) error {
	var filePath string
	var err error

	if chunks > 1 {
		filePath, err = saveFile(file, constant.TempDir)
		if err != nil {
			return err
		}

		// merge files
		unMergeFilesMap[dstFileName][filePath] = struct{}{}
		if len(unMergeFilesMap[dstFileName]) != chunks {
			return nil
		}
		// sort files
		sortedFiles := sortFiles(unMergeFilesMap[dstFileName])
		filePath, err = mergeFiles(constant.StaticDir, dstFileName, sortedFiles)

	} else {
		filePath, err = saveFile(file, constant.StaticDir)
	}
	if err != nil {
		return err
	}

	// todo, tomorrow?
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

// 多文件合并
