package model

type TaskModel struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type TaskPaginationModel struct {
	Data  []TaskPaginationData `json:"data"`
	Page  int64                `json:"page"`
	Limit int64                `json:"limit"`
	Total int64                `json:"total"`
}

type TaskPaginationData struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
