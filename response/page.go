package response

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
	PageNow   int `json:"page_now"`
	TotalData int `json:"total_data"`
	NextPage  int `json:"next_page"`
	PrevPage  int `json:"prev_page"`
}

func NewPagination(page, limit, totalData int) Pagination {
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}
	totalPage := totalData / limit
	if totalData%limit > 0 {
		totalPage++
	}
	pageNow := page
	if pageNow > totalPage {
		pageNow = totalPage
	}
	if pageNow < 1 {
		pageNow = 1
	}
	nextPage := pageNow + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}
	prevPage := pageNow - 1
	if prevPage < 1 {
		prevPage = 1
	}
	return Pagination{
		Page:      page,
		Limit:     limit,
		TotalPage: totalPage,
		PageNow:   pageNow,
		TotalData: totalData,
		NextPage:  nextPage,
		PrevPage:  prevPage,
	}
}
