package model

type pagination struct {
	currentPage, prev, next, pageSize, total int
	queryString                              string
	id                                       int
}
