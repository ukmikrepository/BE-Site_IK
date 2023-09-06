package model

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type Meta struct {
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
	Total  int64 `json:"total"`
}
