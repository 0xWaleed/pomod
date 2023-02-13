package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/0xWaleed/pomod/cmd/pomodserver/dto"
)

type getActiveTaskHandler struct {
	s *pomodoServer
}

func (s *getActiveTaskHandler) handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentTask := s.s.currentTask
		if currentTask == nil {
			return errors.New("there is no active task, enjoy your time")
		}

		return c.JSON(dto.GetTaskDto{
			ID:    currentTask.ID,
			Title: currentTask.Title,
		})
	}
}
