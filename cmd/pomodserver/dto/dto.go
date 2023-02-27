package dto

import "time"

type CreateTaskDto struct {
	Title            string `json:"title"`
	WorkLength       int64  `json:"workLength"`
	ShortBreakLength int64  `json:"shortBreakLength"`
	LongBreakLength  int64  `json:"longBreakLength"`
}

type GetTaskDto struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type UpdateTaskDto struct {
	WorkLength       *time.Duration `json:"workLength"`
	ShortBreakLength *time.Duration `json:"shortBreakLength"`
	LongBreakLength  *time.Duration `json:"longBreakLength"`
}
