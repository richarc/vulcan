package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/richarc/vulcan/config"
	"github.com/richarc/vulcan/handler"
)

func main() {

	//setup our applications basic configuration
	config.Setup()

	//create a new engine
	engine := html.New("./views", ".html")
	cfg := fiber.Config{
		Views: engine,
	}

	app := fiber.New(cfg)
	app.Use(logger.New())

	app.Get("/", handler.Hello)
	app.Get("/chat", handler.Chat)
	app.Post("/send", handler.Send)
	app.Static("/assets", "./assets")

	log.Fatal(app.Listen(config.PORT))
}
