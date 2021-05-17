package pagination

type Pagination struct {
	CurrentPage  int  `json:"current_page"`
	PreviousPage bool `json:"previous_page"`
	NextPage     bool `json:"next_page"`
	Limit        int  `json:"limit"`
	Offset       int  `json:"offset"`
	TotalPages   int  `json:"total_pages"`
	TotalRows    int  `json:"total_rows"`
}

func GetPaginationData(CurrentPage, SizePerPage, TotalRow int) *Pagination {
	pagination := Pagination{}
	pagination.Limit = SizePerPage
	pagination.TotalRows = TotalRow
	pagination.TotalPages = TotalRow / pagination.Limit
	pagination.CurrentPage = CurrentPage
	pagination.Offset = pagination.Limit * (CurrentPage - 1)

	if pagination.Offset == 0 {
		pagination.PreviousPage = false
	} else if pagination.CurrentPage <= pagination.TotalPages {
		pagination.PreviousPage = true
	}

	if pagination.CurrentPage == pagination.TotalPages {
		pagination.NextPage = false
	} else {
		pagination.NextPage = true
	}

	return &pagination
}
