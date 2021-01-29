package common

type PageResult struct {
	Total int64 `json:"total"`
	Rows []interface{} `json:"rows, omitempty"`
}