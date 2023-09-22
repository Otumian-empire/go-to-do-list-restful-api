package web

var (
	DEFAULT_PAGE_NUMBER = 1
	DEFAULT_PAGE_SIZE   = 20
)

type PaginationParameters struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
}

func GetPaginationParams(count, pageNumber, pageSize int) PaginationParameters {
	pagination := PaginationParameters{
		PageNumber: DEFAULT_PAGE_NUMBER,
		PageSize:   DEFAULT_PAGE_SIZE,
		TotalCount: count,
	}

	if pageNumber*pageSize < count {
		pagination.PageNumber = pageNumber + 1
	}

	if pageSize <= count {
		pagination.PageSize = pageSize
	}

	return pagination
}

func CleanPaginationParams(pageNumber, pageSize int) PaginationParameters {
	if pageNumber < 1 {
		pageNumber = DEFAULT_PAGE_NUMBER
	}

	if pageSize < 1 {
		pageSize = DEFAULT_PAGE_SIZE
	}

	return PaginationParameters{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		TotalCount: 0,
	}

}
