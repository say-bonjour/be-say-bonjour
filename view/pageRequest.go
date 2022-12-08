package view

type PageRequest struct {
	Page    int64  `form:"page" json:"page" binding:"required"`
	Size    int64  `form:"size" json:"size" binding:"required"`
	Keyword string `form:"keyword" json:"-"`
	Sort    string `form:"sort" json:"-"`
	Desc    bool   `form:"desc" json:"-"`
}
