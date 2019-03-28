package handler

const (
	QueryAllCnt = 100000
)

type QueryPageForm struct {
	Page int64 `form:"page" binding:"min=1"`
	Size int64 `form:"size" binding:"min=1"`
}

type DeleteForm struct {
	Id int64 `form:"id" binding:"required"`
}

type DeleteByCodeForm struct {
	Code string `form:"code" binding:"required"`
}

type PageInfo struct {
	CurrentPage int64 `json:"current_page"`
	TotalPages  int64 `json:"total_pages"`
	TotalCount  int64 `json:"total_count"`
}

type CommonDeviceQueryForm struct {
	Code           string `form:"device_code"`
	Name           string `form:"device_name"`
	Model          string `form:"device_model"`
	Brand          string `form:"device_brand"`
	TagCode        string `form:"device_tag_code"`
	DepartmentCode string `form:"device_department_code"`
}

func (f CommonDeviceQueryForm) ConditionExist() bool {
	return f.Code != "" || f.Name != "" || f.Model != "" || f.Brand != "" || f.TagCode != "" || f.DepartmentCode != ""
}
