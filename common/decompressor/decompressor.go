package decompress

import (
	"context"
	"errors"
	"github/leezone/simulator/common/logger"
	"io"
	"os"
	"sync"
)

var ErrNotSupportDepression = errors.New("do not support decompression")

const (
	ArchiveFileSuccess = "archive file success"
	OpenFileErr        = "open archive file error"
	CreateFileErr      = "create file error"
	SaveFileErr        = "save file err"
)

type FileContext struct {
	IsDir bool
	Path  string
	Ctx   context.Context
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

	funcs := make([]func(FileContext, io.Reader) error, len(iterateFuncs)+1)
	funcs[0] = saveFile
	copy(funcs[1:], iterateFuncs)

	return d.strategy.Decompress(file, dest, funcs...)
}

func GetDecompressorByExtension(ext string) (Decompressor, error) {
	switch ext {
	case ".zip":
		return &ZipDecompressor{}, nil
	default:
		return nil, nil
	}
}

func saveFile(ctx FileContext, reader io.Reader) error {
	if ctx.IsDir {
		err := os.MkdirAll(ctx.Path, 0x644)
		if err != nil {
			logger.Logger.ErrorWithStack(CreateFileErr, err, ctx)
			return errors.New(CreateFileErr)
		}
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
