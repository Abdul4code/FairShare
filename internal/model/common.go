package model

// GroupQuery holds the filter and pagination parameters
// for querying groups from the database.
type MetaData struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PageSize    int `json:"page_size"`
	Total       int `json:"total"`
}
