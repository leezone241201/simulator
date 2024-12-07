package apistruct

type Page struct {
	PageNum   int `json:"pageNum" form:"pageNum" binding:"gte=0"`
	PageCount int `json:"pageCount" form:"pageCount" binding:"gte=0"`
}
