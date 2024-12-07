package apistruct

type FileRequest struct {
}

type FileResponse struct {
}

type UploadResponse struct {
	FileName string `json:"fileName"`
	Result   string `json:"result"`
}
