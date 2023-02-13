package handlers

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/0xWaleed/pomod"
)

type startActiveTaskHandler struct {
	s *pomodoServer
}

func (s *startActiveTaskHandler) handler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		currentTask := s.s.currentTask

		if currentTask == nil {
			return errors.New("there is no active task")
		}

		session := currentTask.CurrentSession

		isDone := session.Done()

		taskHasSession := session != nil

		currentSessionIsActive := session == s.s.activeSession

		alreadyStartedAndNotFinished := taskHasSession && currentSessionIsActive && !isDone

		if alreadyStartedAndNotFinished {
			return errors.New("task already has an active session")
		}

		go func() {

			log.Println("Starting the session", session.Task.Title)
			s.s.mu.Lock()
			s.s.isTaskRunning = true
			s.s.mu.Unlock()

			if session != nil && isDone {
				currentTask.Next()
			}

			s.s.activeSession = session

			for {
				select {
				case <-s.s.abortTask:
					s.s.mu.Lock()
					log.Println("start: Aborting current task", s.s.currentTask.Title)
					s.s.isTaskRunning = false
					s.s.mu.Unlock()
					return
				default:
					session := currentTask.Tick()

					s.broadcast(session)

					if session.Done() {
						log.Println("Done Ticking:", session)
						s.s.mu.Lock()
						s.s.isTaskRunning = false
						s.s.mu.Unlock()
						return
					}
				}

			}

		}()

		return nil
	}
}

func (s *startActiveTaskHandler) broadcast(session *pomod.Session) {
	users := s.s.clientCollection.GetAll()

	for _, user := range users {
		conn := user.C
		remaining := session.Remaining()
		minutes := remaining.Truncate(time.Minute) / time.Minute
		seconds := remaining.Truncate(time.Second) / time.Second % 60
		typeAsString := session.Type.String()
		isDone := session.Done()
		title := session.Task.Title

		err := conn.WriteJSON(map[string]any{
			"minutes": minutes,
			"seconds": seconds,
			"type":    typeAsString,
			"isDone":  isDone,
			"title":   title,
		})

		if err != nil {
			log.Println()
		}
	}
}
