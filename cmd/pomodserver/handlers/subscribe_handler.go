package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"github.com/0xWaleed/pomod/cmd/pomodserver/models"
)

type subscribeTaskHandler struct {
	s *pomodoServer
}

func (s *subscribeTaskHandler) handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		upgraded := websocket.IsWebSocketUpgrade(c)
		log.Println(c, upgraded)
		if !upgraded {
			return fiber.ErrUpgradeRequired
		}

		return websocket.New(func(socket *websocket.Conn) {
			// user is only live in this block
			user := models.User{
				ID: socket.Params("id"),
				C:  socket,
			}

			s.s.clientCollection.Add(&user)

			var (
				mt  int
				msg []byte
				err error
			)

			for {
				if mt, msg, err = socket.ReadMessage(); err != nil {
					log.Println("closed", mt, msg, err)
					s.s.clientCollection.Remove(&user)
					break
				}
				log.Println("received message:", mt, msg, err)
			}
		})(c)
	}
}
