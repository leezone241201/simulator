package constant

const StaticDir = "./static"

var AllowUploadSuffix = map[string]struct{}{
	".zip":    {},
	".7z":     {},
	".rar":    {},
	".tar.gz": {},

	".json": {},
}
