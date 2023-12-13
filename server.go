package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	app.Static("/static", "./public") 

	app.Listen(":3000")

}