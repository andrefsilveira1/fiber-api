package main

import (
	"api/src/config"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Person struct {
	Name string `json:"name" form:"name"`
	Pass string `json:"pass" form:"pass"`
}

func main() {
	config.Load()
	fmt.Print("Listening on port: $d \n", config.Port)
	fmt.Println(config.StringConexaoBanco)
	fmt.Println("Rodando API")
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
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return err
		}

		if erro := user.Prepare("register"); erro != nil {
			fmt.Println("Status Bad request")
			return erro
		}
		db, erro := db.Connect()
		if erro != nil {
			fmt.Println("Algo deu errado:", erro)
			return erro
		}
		defer db.Close()
		repositorie := repositories.NewRepository(db)
		Id, erro := repositorie.CreateUser(user)
		if erro != nil {
			fmt.Println("Status Internal Server Error", erro)
			return erro
		}
		fmt.Println("Id do usuário:", Id)

		return c.Render("send", fiber.Map{
			"Nome": user.Name,
			"Pass": user.Email,
			"Id":   Id,
		})
	})

	fmt.Println("Rodando aplicação")
	app.Listen(":3000")

}
