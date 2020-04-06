// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Listen(8080)
}
