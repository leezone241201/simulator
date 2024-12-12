package constant

const StaticDir = "./static/uploadFiles"
const TempDir = "./static/temp"

var AllowUploadSuffix = map[string]struct{}{
	".zip":    {},
	".tar":    {},
	".tar.gz": {},

	".json": {},
}
