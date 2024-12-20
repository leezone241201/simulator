package syserror

const SuccessCode = 1000

// 客户端错误
const (
	ParameterErrCode = 4000 + iota
)

// 服务器错误
const (
	UploadFileErrCode = 5000 + iota
	UploadFileNotAllowedErrCode
	TimeOutErrCode
	InternalErrCode
)

var ErrMap = map[int]string{
	ParameterErrCode: "参数错误",

	UploadFileErrCode:           "上传文件失败!",
	UploadFileNotAllowedErrCode: "不支持的上传类型,仅支持zip,rar,json,7z格式!",
	TimeOutErrCode:              "请求超时,请稍后重试!",
	InternalErrCode:             "系统内部错误,请稍后重试!",
}
