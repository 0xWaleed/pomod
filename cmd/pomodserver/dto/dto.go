package dto

import (
	"time"

	"github.com/0xWaleed/pomod"
)

type CreateTaskDto struct {
	Title            string `json:"title"`
	WorkLength       int64  `json:"workLength"`
	ShortBreakLength int64  `json:"shortBreakLength"`
	LongBreakLength  int64  `json:"longBreakLength"`
}

type TaskOptionsDto struct {
	WorkLength       time.Duration `json:"workLength"`
	ShortBreakLength time.Duration `json:"shortBreakLength"`
	LongBreakLength  time.Duration `json:"longBreakLength"`
	LongBreakAfter   int           `json:"longBreakAfter"`
}

func CreateTaskOptionsDto(o *pomod.TaskOptions) TaskOptionsDto {
	return TaskOptionsDto{
		WorkLength:       o.WorkLength / time.Second,
		ShortBreakLength: o.ShortBreakLength / time.Second,
		LongBreakLength:  o.LongBreakLength / time.Second,
		LongBreakAfter:   o.LongBreakAfter,
	}
}

type GetTaskDto struct {
	ID      string         `json:"id"`
	Title   string         `json:"title"`
	Options TaskOptionsDto `json:"options"`
}

type UpdateTaskDto struct {
	WorkLength       *time.Duration `json:"workLength"`
	ShortBreakLength *time.Duration `json:"shortBreakLength"`
	LongBreakLength  *time.Duration `json:"longBreakLength"`
}
