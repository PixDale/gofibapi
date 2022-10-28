package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func helloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func allHandler(c *fiber.Ctx) error {
	queries := make([]FibResult, 0)
	addToQueries := func(key interface{}, value interface{}) bool {
		queries = append(queries, value.(FibResult))
		return true
	}
	SavedResponses.Range(addToQueries)
	return c.Status(fiber.StatusOK).JSON(AllResponse{
		Queries: queries,
	})
}

func fibHandler(c *fiber.Ctx) error {
	const TIMEOUT = 500 * time.Millisecond
	inputStr := c.Params("input")
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	response := callFib(uint64(input), TIMEOUT)
	return c.Status(fiber.StatusOK).JSON(response)
}
