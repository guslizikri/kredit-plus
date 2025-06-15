package util

type MetaPagination struct {
	Next  *int `json:"next"`
	Prev  *int `json:"prev"`
	Total int  `json:"total"`
}

func BuildPaginationMeta(page, limit, total int) MetaPagination {
	offset := (page - 1) * limit
	var next *int
	if offset+limit < total {
		n := page + 1
		next = &n
	}
	var prev *int
	if page > 1 {
		p := page - 1
		prev = &p
	}
	return MetaPagination{
		Next:  next,
		Prev:  prev,
		Total: total,
	}
}
