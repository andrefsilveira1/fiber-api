package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Person struct {
	Name string `json:"name" form:"name"`
	Pass string `json:"pass" form:"pass"`
}

func main() {
	engine := html.New("./public/views/", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "../partials/template",
	})
	app.Static("/", "../public")
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Render("home/index", fiber.Map{
			"Title":    "Portfólio",
			"Subtitle": "This portfolio was developed by André Freitas Silveira",
		})
	})

	app.Get("value/:value", func(c *fiber.Ctx) error {
		return c.SendString("Value: " + c.Params("value"))

	})

	app.Get("api/", func(c *fiber.Ctx) error {
		return c.Render("api/index", fiber.Map{})
	})

	app.Post("api/enviar", func(c *fiber.Ctx) error {
		person := new(Person)
		if err := c.BodyParser(person); err != nil {
			return err
		}
		log.Println(person.Name)
		log.Println(person.Pass)

		return c.Render("api", fiber.Map{
			"Nome": person.Name,
			"Pass": person.Pass,
		})
	})

	fmt.Println("Rodando aplicação")
	app.Listen(":3000")

}
