package database

type PageInfo struct {
	PageNumber int
	PageSize   int
}

func (pageInfo PageInfo) Offset() int {
	return (pageInfo.PageNumber - 1) * pageInfo.PageSize
}
