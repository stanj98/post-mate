package main

import "github.com/gofiber/fiber/v2"

func main() {

	/*
		Plan:
		1. Create note structure - perform CRUD and duplicate notes
		2. Use mongodb instead of point 1
		3. Refer to this: https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34

	*/

	app := fiber.New()

	app.Static("/", "./public") 

	app.Get("/api/view-notes", func(c *fiber.Ctx) error {
		return c.SendString("All notes")
	})

	

	app.Listen(":3000")

}