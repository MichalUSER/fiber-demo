package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  uint64 `json:"age"`
	}

	app := fiber.New()

	app.Post("/api", func(c *fiber.Ctx) error {
		p := new(Person)
		if err := c.BodyParser(p); err != nil {
			return err
		}
		return c.SendString(fmt.Sprintf("Hello %v!", p.Name))
	})

	app.Static("/", "./web")

	err := app.Listen(":3030")
	if err != nil {
		log.Fatal(err)
	}
}
