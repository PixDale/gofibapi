package main

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var SavedResponses sync.Map // map thread safe

func main() {
	app := fiber.New()

	SavedResponses = sync.Map{}

	app.Get("/", helloHandler)
	app.Get("/fib/:input", fibHandler)
	app.Get("/all", allHandler)

	app.Listen(":3000")
}
