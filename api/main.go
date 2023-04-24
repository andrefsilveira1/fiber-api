package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "../public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":    "Portfólio",
			"Subtitle": "This portfolio was developed by André Freitas Silveira",
		})
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("Value: " + c.Params("value"))
		//Getting request from the parameters
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("Path: " + c.Params("*"))
		//Getting request from the parameters
	})

	fmt.Println("Rodando aplicação")
	app.Listen(":3000")
}
