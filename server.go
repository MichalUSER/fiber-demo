package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	type Person struct {
		name string
		Age  uint64
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/:name/:age", func(c *fiber.Ctx) error {
		age, _ := strconv.ParseUint(c.Params("age"), 10, 8)
		name := c.Params("name")
		msg := Person{
			name: name,
			Age:  age,
		}
		return c.JSON(msg)
	})

	app.Static("/web", "./web")

	app.Listen(":3000")
}
