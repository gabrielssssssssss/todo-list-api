package entity

import "time"

type TaskEntity struct {
	Id          string    `json:"id"`
	OwnerId     string    `json:"ownerId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Token       string    `json:"token"`
}

type TaskPaginationEntity struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
