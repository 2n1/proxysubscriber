package db

import "math"

type Paginate struct {
	TotalRecord     int         `json:"total_record"`
	TotalPage       int         `json:"total_page"`
	Page            int         `json:"page"`
	PageSize        int         `json:"page_size"`
	Data            interface{} `json:"data"`
	Query           map[string]interface{}
	PageList        []int
	PageListDisplay []int
}

func NewPaginate(totalRecord, page, pageSize int, data interface{}) *Paginate {
	pageTotal := int(math.Ceil(float64(totalRecord) / float64(pageSize)))
	var pl []int
	for i := 0; i < pageTotal; i++ {
		pl = append(pl, i+1)
	}
	return &Paginate{
		TotalRecord:     totalRecord,
		TotalPage:       pageTotal,
		Page:            page,
		PageSize:        pageSize,
		Data:            data,
		Query:           nil,
		PageList:        pl,
	}
}
