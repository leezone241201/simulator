package decompress

import (
	"archive/tar"
	"fmt"
	"io"
	"os"

	"github/leezone/simulator/common/logger"
)

type TarDecompressor struct{}

func decompress(reader io.Reader, dest string, iterateFuncs ...func(FileContext, io.Reader) error) (interface{}, error) {
	tarReader := tar.NewReader(reader)
	var tarHeader *tar.Header
	var err error

	var results = make([]DecompressResult, 0)
	var singleResult DecompressResult
	var ctx FileContext

	for {
		tarHeader, err = tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Logger.ErrorWithStack(GetHeaderErr, err, nil)
			break
		}

		singleResult.FileName = tarHeader.Name
		ctx.Path = fmt.Sprintf("%s%c%s", dest, os.PathSeparator, singleResult.FileName)
		ctx.CurrentIndex++

		switch tarHeader.Typeflag {
		case tar.TypeDir:
			ctx.IsDir = true
		default:
			ctx.IsDir = false
		}

		for _, iterateFunc := range iterateFuncs {
			err = iterateFunc(ctx, tarReader)
			if err != nil {
				singleResult.Msg = err.Error()
				logger.Logger.ErrorWithStack(CustomerExecErr, err, nil)
				goto OneFinished
			}
		}

		singleResult.Msg = ArchiveFileSuccess

	OneFinished:
		results = append(results, singleResult)
	}

	return results, nil
}

func (t *TarDecompressor) Decompress(file *os.File, dest string, iterateFuncs ...func(FileContext, io.Reader) error) (interface{}, error) {
	return decompress(file, dest, iterateFuncs...)
}
