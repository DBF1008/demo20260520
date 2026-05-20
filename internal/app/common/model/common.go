package model


type PageReq struct {
	DateRange []string `p:"dateRange"`
	PageNum   int      `p:"pageNum"`
	PageSize  int      `p:"pageSize"`
	OrderBy   string   `p:"orderBy"`
}


type ListRes struct {
	CurrentPage int         `json:"currentPage"`
	Total       interface{} `json:"total"`
}
