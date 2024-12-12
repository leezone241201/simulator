package decompress

import (
	"context"
	"errors"
	"io"
	"os"
	"sync"

	"github/leezone/simulator/common/logger"
)

var ErrNotSupportDepression = errors.New("do not support decompression")

const (
	ArchiveFileSuccess = "archive file success"
	OpenFileErr        = "open archive file error"
	CreateFileErr      = "create file error"
	SaveFileErr        = "save file error"
	GetHeaderErr       = "get file header error"
	CustomerExecErr    = "exec customer function error"
)

type FileContext struct {
	IsDir          bool
	Path           string
	TotalFileCount int
	CurrentIndex   int
	Ctx            context.Context
}

type DecompressResult struct {
	FileName string
	Msg      string
}

type Decompressor interface {
	Decompress(*os.File, string, ...func(FileContext, io.Reader) error) (interface{}, error)
}

type DecompressionContext struct {
	mu       sync.RWMutex
	strategy Decompressor
}

func (d *DecompressionContext) SetStrategy(decompressor Decompressor) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.strategy = decompressor
}

func (d *DecompressionContext) Decompress(file *os.File, dest string, iterateFuncs ...func(FileContext, io.Reader) error) (interface{}, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.strategy == nil {
		return nil, ErrNotSupportDepression
	}

	funcs := make([]func(FileContext, io.Reader) error, 0, len(iterateFuncs)+1)
	funcs = append(funcs, saveFile)
	copy(funcs[1:], iterateFuncs)

	return d.strategy.Decompress(file, dest, funcs...)
}

func GetDecompressorByExtension(ext string) (Decompressor, error) {
	switch ext {
	case ".zip":
		return &ZipDecompressor{}, nil
	case ".tar":
		return &TarDecompressor{}, nil
	case ".tar.gz":
		return &TarGzDecompressor{}, nil
	default:
		return nil, ErrNotSupportDepression
	}
}

func saveFile(ctx FileContext, reader io.Reader) error {
	if ctx.IsDir {
		err := os.MkdirAll(ctx.Path, 0x644)
		if err != nil {
			logger.Logger.ErrorWithStack(CreateFileErr, err, ctx)
			return errors.New(CreateFileErr)
		}
		return nil
	}

	file, err := os.Create(ctx.Path)
	if err != nil {
		logger.Logger.ErrorWithStack(CreateFileErr, err, ctx)
		return errors.New(CreateFileErr)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		logger.Logger.ErrorWithStack(SaveFileErr, err, ctx)
		return errors.New(SaveFileErr)
	}

	return nil
}
