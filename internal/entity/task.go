package entity

import "time"

type TaskEntity struct {
	TaskId      string
	OwnerId     string
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskPaginationEntity struct {
	Page    int64
	Limit   int64
	OwnerId string
}
