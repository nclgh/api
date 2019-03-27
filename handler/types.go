package handler

type QueryPageForm struct {
	Page int64 `form:"page" binding:"min=1"`
	Size int64 `form:"size" binding:"min=1"`
}

type DeleteForm struct {
	Id int64 `form:"id" binding:"required"`
}

type PageInfo struct {
	CurrentPage int64 `json:"current_page"`
	TotalPages  int64 `json:"total_pages"`
	TotalCount  int64 `json:"total_count"`
}
