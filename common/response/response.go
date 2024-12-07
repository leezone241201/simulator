package response

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, code int, codeStr string, data ...interface{}) {
	result := gin.H{
		"code": code,
		"msg":  codeStr,
	}
	if len(data) > 0 {
		successWithData(result, data[0])
	}
	ctx.JSON(http.StatusOK, result)
}

func successWithData(result gin.H, data interface{}) {
	if reflect.TypeOf(data).Kind() == reflect.Slice {
		result["total"] = reflect.ValueOf(data).Len()
	}
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
