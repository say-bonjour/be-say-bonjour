package view

type pageable struct {
	PageRequest
	TotalPages int64       `json:"totalPages"`
	TotalRows  int64       `json:"totalRows"`
	Data       interface{} `json:"data"`
}

func NewPageable(pageRequest PageRequest, totalRows int64) pageable {
	if len(pageRequest.Sort) == 0 {
		pageRequest.Sort = "id"
	}

	if pageRequest.Desc {
		pageRequest.Sort += " desc"
	}

	if pageRequest.Size > 1000 {
		pageRequest.Size = 1000
	}

	return pageable{PageRequest: pageRequest, TotalPages: totalRows, TotalRows: (totalRows / pageRequest.Size) + 1}
}

func (request pageable) GetOffset() int {
	return int((request.Page - 1) * request.Size)
}
