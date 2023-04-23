package domain

type NextToken struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}