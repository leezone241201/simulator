package decompress

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github/leezone/simulator/common/logger"
)

type ZipDecompressor struct{}

func (z *ZipDecompressor) Decompress(file *os.File, dest string, iterateFuncs ...func(FileContext, io.Reader) error) (interface{}, error) {
	zipReader, err := zip.OpenReader(file.Name())
	if err != nil {
		logger.Logger.ErrorWithStack(OpenFileErr, err, file)
		return nil, err
	}
	defer zipReader.Close()

	var results = make([]DecompressResult, 0, len(zipReader.File))
	var singleResult DecompressResult
	var ctx FileContext

	ctx.TotalFileCount = len(zipReader.File)
	for i, zipFile := range zipReader.File {
		singleResult.FileName = zipFile.Name

		file, err := zipFile.Open()
		if err != nil {
			singleResult.Msg = OpenFileErr
			logger.Logger.ErrorWithStack(OpenFileErr, err, nil)
			goto OneFinished
		}

		ctx.IsDir = zipFile.FileInfo().IsDir()
		ctx.Path = fmt.Sprintf("%s%c%s", dest, os.PathSeparator, singleResult.FileName)
		ctx.CurrentIndex = i

		for _, iterateFunc := range iterateFuncs {
			err = iterateFunc(ctx, file)
			if err != nil {
				file.Close()
				singleResult.Msg = err.Error()
				logger.Logger.ErrorWithStack("exec customer function error", err, nil)
				goto OneFinished
			}
		}

		file.Close()
		singleResult.Msg = ArchiveFileSuccess

	OneFinished:
		results = append(results, singleResult)
	}

	return results, nil
}
