package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/0xWaleed/pomod/cmd/pomodserver/dto"
)

type getTasksHandler struct {
	s *pomodoServer
}

func (s *getTasksHandler) handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasksDto := make([]dto.GetTaskDto, 0)

		tasks := s.s.tasks
		for _, task := range tasks {
			tasksDto = append(tasksDto, dto.GetTaskDto{
				ID:      task.ID,
				Title:   task.Title,
				Options: dto.CreateTaskOptionsDto(&task.Options),
			})
		}

		return c.JSON(tasksDto)
	}
}
