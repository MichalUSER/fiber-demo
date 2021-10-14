package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func YearBorn(age uint64, ch chan uint64) {
	d := time.Now()
	year := d.Year()
	ch <- uint64(year) - age
}

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

		ch := make(chan uint64)
		go YearBorn(p.Age, ch)
		v := <-ch

		return c.JSON(fiber.Map{
			"greeting": fmt.Sprintf("Hello %v!", p.Name),
			"yearBorn": v,
		})
	})

	app.Static("/", "./web")

	log.Fatal(app.Listen(":3030"))
}
