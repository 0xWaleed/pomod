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
		var tasksDto []dto.GetTaskDto

		tasks := s.s.tasks
		for _, task := range tasks {
			tasksDto = append(tasksDto, dto.GetTaskDto{
				ID:    task.ID,
				Title: task.Title,
			})
		}

		return c.JSON(tasksDto)
	}
}