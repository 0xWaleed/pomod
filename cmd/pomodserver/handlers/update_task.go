package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/0xWaleed/pomod/cmd/pomodserver/dto"
)

type updateTaskHandler struct {
	s *pomodoServer
}

func (s *updateTaskHandler) handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		taskId := c.Params("id")

		//todo lock to avoid data race
		targetIndex := -1
		for index, task := range s.s.tasks {
			if task.ID == taskId {
				targetIndex = index
			}
		}

		if targetIndex == -1 {
			return errors.New("invalid task id")
		}

		var payload dto.UpdateTaskDto
		if err := c.BodyParser(&payload); err != nil {
			//todo change to better response
			return errors.New("invalid task body")
		}

		fmt.Println("received payload:", payload)

		//todo: dirty patch, clean up the code
		if payload.WorkLength != nil {
			s.s.tasks[targetIndex].Options.WorkLength = *payload.WorkLength * time.Second
		}

		if payload.LongBreakLength != nil {
			s.s.tasks[targetIndex].Options.LongBreakLength = *payload.LongBreakLength * time.Second
		}

		if payload.ShortBreakLength != nil {
			s.s.tasks[targetIndex].Options.ShortBreakLength = *payload.ShortBreakLength * time.Second
		}

		return nil
	}
}
