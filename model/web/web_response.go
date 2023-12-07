package web

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

type Pagination struct {
	Page      int   `json:"page"`
	PerPage   int   `json:"per_page"`
	TotalData int64 `json:"total_data"`
	TotalPage int   `json:"total_page"`
}

type PaginationLinks struct {
	PrevUrl string `json:"prev_url"`
	NextUrl string `json:"next_url"`
}

type AddsDatabaseResponse struct {
	TotalData int64
}
