package main

import (
	"log"

	"github.com/JerryLegend254/fiber-api/database"
	"github.com/JerryLegend254/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Post("/users", routes.CreateUser)
	app.Get("/users", routes.GetUsers)
	app.Get("/users/:id", routes.GetUser)
	app.Put("/users/:id", routes.UpdateUser)

    log.Fatal(app.Listen(":3000"))
}
