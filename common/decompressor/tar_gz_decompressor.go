package decompress

import (
	"compress/gzip"
	"io"
	"os"

	"github/leezone/simulator/common/logger"
)

type TarGzDecompressor struct{}

func (t *TarGzDecompressor) Decompress(file *os.File, dest string, iterateFuncs ...func(FileContext, io.Reader) error) (interface{}, error) {
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		logger.Logger.ErrorWithStack(OpenFileErr, err, file)
		return nil, err
	}

	return decompress(gzipReader, dest, iterateFuncs...)
}
