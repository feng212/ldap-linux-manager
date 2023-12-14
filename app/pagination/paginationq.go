package pagination

type PaginationQ struct {
	Ok    bool        `json:"ok"`
	Size  int         `form:"size" json:"size"`
	Page  int         `form:"page" json:"page"`
	Data  interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total int         `json:"total"`
}
