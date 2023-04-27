package main

import (
	"fmt"
	"log"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Person struct {
	Name string `json:"name" form:"name"`
	Pass string `json:"pass" form:"pass"`
}

func main() {
	engine := html.New("./src/public/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":    "Portfólio",
			"Subtitle": "This portfolio was developed by André Freitas Silveira",
		})
	})

	app.Get("value/:value", func(c *fiber.Ctx) error {
		return c.SendString("Value: " + c.Params("value"))

	})
	app.Get("api/", func(c *fiber.Ctx) error {
		return c.Render("send", fiber.Map{})
	})

	app.Post("api/enviar", func(c *fiber.Ctx) error {
		person := new(Person)
		if err := c.BodyParser(person); err != nil {
			return err
		}
		log.Println(person.Name)
		log.Println(person.Pass)
		_, err := mail.ParseAddress(person.Name)
		if err != nil {
			condition := true
			log.Println("Erro: Email inválido!", err)
			return c.Render("send", fiber.Map{
				"Erro":     err,
				"hasError": condition,
			})
		}

		output := make([]Person, 1)
		output = append(output, *person)
		log.Println(output)

		return c.Render("send", fiber.Map{
			"Nome": person.Name,
			"Pass": person.Pass,
		})
	})

	fmt.Println("Rodando aplicação")
	app.Listen(":3000")

}
