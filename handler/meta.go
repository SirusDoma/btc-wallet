package handler

type Meta struct {
	Status int `json:"http_status"`
}

type PaginationMeta struct {
	Meta
	Offset uint `json:"offset"`
	Limit  uint `json:"limit"`
	Total  uint `json:"total"`
}
