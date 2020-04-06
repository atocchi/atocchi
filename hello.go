package main

import "github.com/gofiber/fiber"

func main() {
  app := fiber.New()


  app.Static("/", "./public") 

  app.Get("/hello", func(c *fiber.Ctx) {
	c.Send("hello World")
  })

  

  app.Listen(3001)
}