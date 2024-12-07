package middleware

import (
	"context"
	"github/leezone/simulator/common/response"
	"github/leezone/simulator/common/syserror"
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout(ctx *gin.Context) {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	success := make(chan struct{})
	go func() {
		ctx.Next()
		close(success)
	}()

	select {
	case <-timeoutCtx.Done():
		response.Failure(ctx, syserror.TimeOutErrCode, syserror.ErrMap[syserror.TimeOutErrCode])
		ctx.Abort()
	case <-success:
		// TODO,记录API操作日志
		return
	}
}
