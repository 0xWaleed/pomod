package handlers

import (
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/0xWaleed/pomod"
	"github.com/0xWaleed/pomod/cmd/pomodserver/models"
)

type pomodoServer struct {
	mu               sync.Mutex
	tasks            []pomod.Task
	currentTask      *pomod.Task
	activeSession    *pomod.Session
	clientCollection models.UserCollection
	abortTask        chan bool
	// todo, ensure to lock this to avoid data race
	isTaskRunning bool
}

func (s *pomodoServer) createTaskHandler() fiber.Handler {
	c := createTaskHandler{
		s,
	}
	return c.handler()
}

func (s *pomodoServer) getTasksHandler() fiber.Handler {
	c := getTasksHandler{s}
	return c.handler()
}

func (s *pomodoServer) getActiveTaskHandler() fiber.Handler {
	c := getActiveTaskHandler{s}
	return c.handler()
}

func (s *pomodoServer) activateTaskHandler() fiber.Handler {
	c := activateTaskHandler{s}
	return c.handler()
}

func (s *pomodoServer) startActiveTask() fiber.Handler {
	c := startActiveTaskHandler{s}
	return c.handler()
}

func (s *pomodoServer) subscribeHandler() fiber.Handler {
	c := subscribeTaskHandler{s}
	return c.handler()
}

func (s *pomodoServer) updateTaskHandler() fiber.Handler {
	c := updateTaskHandler{s}
	return c.handler()
}

func NewPomodoServer(app *fiber.App) {
	s := pomodoServer{
		abortTask: make(chan bool),
	}
	app.Post("/tasks", s.createTaskHandler())

	app.Get("/tasks", s.getTasksHandler())
	app.Get("/task", s.getActiveTaskHandler())
	app.Put("/tasks/:id", s.updateTaskHandler())

	app.Post("/tasks/:id/activate", s.activateTaskHandler())
	app.Post("/tasks/:id/start", s.startActiveTask())
	app.Get("/tasks/subscribe", s.subscribeHandler())
}
