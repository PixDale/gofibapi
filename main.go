package main

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

//var SavedResponses map[uint64]FibResult

var SavedResponses sync.Map

func main() {
	app := fiber.New()
	//SavedResponses = make(map[uint64]FibResult)
	SavedResponses = sync.Map{}

	app.Get("/", helloHandler)
	app.Get("/fib/:input", fibHandler)
	app.Get("/all", allHandler)

	app.Listen(":3000")
}
