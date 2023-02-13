package handlers

import "github.com/gofiber/fiber/v2"

type stopActiveTaskHandler struct {
	s *pomodoServer
}

func (s *stopActiveTaskHandler) handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
