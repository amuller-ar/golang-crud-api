package domain

const (
	DefaultPage     = 1
	DefaultPageSize = 10
)

type SearchParameters struct {
	Status      *string
	BoundingBox *BoundingBox
	Page        int
	PageSize    int
}

type PaginatedResponse struct {
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	Total      int64      `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []Property `json:"data"`
}

type Pagination struct {
	Page       int
	Limit      int
	TotalRows  int64
	TotalPages int
	Sort       string
	Rows       interface{}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
