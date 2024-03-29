package models

type Pagination struct {
	Limit        int         `json:"limit,omitempty"`
	Page         int         `json:"page,omitempty"`
	Sort         string      `json:"sort,omitempty"`
	TotalResults int64       `json:"total_results"`
	TotalPages   int         `json:"total_pages"`
	Results      interface{} `json:"results"`
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
		p.Sort = "updated_at desc"
	}
	return p.Sort
}

func (p *Pagination) GetRand() string {
	if p.Sort == "" {
		p.Sort = "RAND()"
	}
	return p.Sort
}
