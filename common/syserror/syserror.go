package syserror

// 服务器错误
const (
	UploadFileErrCode = 5000 + iota
	TimeOutCode
)

var ErrMap = map[int]string{
	UploadFileErrCode: "上传文件失败!",
	TimeOutCode:       "请求超时,请稍后重试!",
}
