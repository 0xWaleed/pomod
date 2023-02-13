package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/0xWaleed/pomod"
	"github.com/0xWaleed/pomod/cmd/pomodserver/dto"
)

type createTaskHandler struct {
	server *pomodoServer
}

func (s *createTaskHandler) handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var taskDto dto.CreateTaskDto
		err := c.BodyParser(&taskDto)

		taskOptions := pomod.TaskOptions{
			WorkLength:       time.Duration(taskDto.WorkLength) * time.Second,
			LongBreakLength:  time.Duration(taskDto.LongBreakLength) * time.Second,
			ShortBreakLength: time.Duration(taskDto.ShortBreakLength) * time.Second,
			AutoSwitch:       false,
			Interval:         time.Second,
			LongBreakAfter:   3,
		}

		if err != nil {
			return err
		}

		if taskDto.Title == "" {
			return errors.New("expected title to be a valid string")
		}

		err = s.ensureThereIsNoTaskWithSameTitle(&taskDto)

		if err != nil {
			return errors.New("task with same title already exist")
		}

		task := pomod.NewTask(taskDto.Title, taskOptions)

		s.server.tasks = append(s.server.tasks, task)

		return nil
	}
}

func (s *createTaskHandler) ensureThereIsNoTaskWithSameTitle(taskDto *dto.CreateTaskDto) error {
	for _, task := range s.server.tasks {
		if task.Title == taskDto.Title {
			return errors.New("task already exist with same name")
		}
	}
	return nil
}
