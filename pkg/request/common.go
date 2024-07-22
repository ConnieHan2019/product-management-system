package request

// ListFilter is a struct that represents the options to list products
type ListFilter struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
}
