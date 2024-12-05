package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, code int, codeStr string, data ...interface{}) {
	result := gin.H{
		"code": code,
		"msg":  codeStr,
	}
	if len(data) > 0 {
		successWithData(result, data)
	}
	ctx.JSON(http.StatusOK, result)
}

func successWithData(result gin.H, data ...interface{}) {
	result["total"] = len(data)
	result["data"] = data
}

func SuccessFile() {

}

func Failure(ctx *gin.Context, code int, codeStr string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  codeStr,
	})
}
