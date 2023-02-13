package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/logger"

	"github.com/0xWaleed/pomod/cmd/pomodserver/handlers"
)

var l *logger.Logger

func main() {
	l = logger.Init("MAIN", false, false, os.Stdout)
	defer l.Close()

	app := fiber.New()

	handlers.NewPomodoServer(app)

	log.Fatal(app.Listen("localhost:8080"))
}
