package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/0xWaleed/pomod"
	"github.com/0xWaleed/pomod/cmd/pomodserver/dto"
)

type activateTaskHandler struct {
	s *pomodoServer
}

func (s *activateTaskHandler) handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		s.s.mu.Lock()
		defer s.s.mu.Unlock()

		id := c.Params("id")

		tasks := s.s.tasks

		var targetTask *pomod.Task

		for index, task := range tasks {
			if task.ID == id {
				targetTask = &tasks[index]
				break
			}
		}

		if targetTask == nil {
			c.Status(http.StatusNotFound)
			return errors.New("not found")
		}

		if s.s.currentTask != nil && s.s.isTaskRunning {
			log.Println("activate: Aborting current task", s.s.currentTask.Title)
			s.s.abortTask <- true
		}

		s.s.currentTask = targetTask
		return c.JSON(dto.GetTaskDto{
			ID:    targetTask.ID,
			Title: targetTask.Title,
		})
	}
}
