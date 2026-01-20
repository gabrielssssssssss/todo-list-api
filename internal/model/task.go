package model

type TaskModel struct {
	TaskId      string
	Title       string
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

type TaskModelResponse struct {
	TaskId      string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TaskPaginationModel struct {
	Page  int64
	Limit int64
	Total int64
}

type TaskPaginationModelResponse struct {
	Data []struct {
		Id          int64  `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}
