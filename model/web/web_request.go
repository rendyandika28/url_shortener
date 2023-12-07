package web

type QueryRequest struct {
	SortBy  string `query:"sort_by"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
	Search  string `query:"search"`
}
