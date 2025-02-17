package main

import (
	"github.com/gofiber/fiber/v2"
   
	"github.com/gofiber/fiber/v2/middleware/logger"
   
   )

   type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
   }

   var users = make(map[string]User)
   
   func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/",func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
		}
		users[user.ID] = user

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":"User register successfully",
			"user":user,
		})
	})

	app.Get("/user/:id",func(c *fiber.Ctx) error {
		id := c.Params("id")
		user, exists := users[id]
		if !exists {
			return c.Status(fiber.StatusNotFound).SendString("user not found")
		}
	 	return c.JSON(user)
	})

	app.Patch("/user/:id",func(c *fiber.Ctx) error {
		id := c.Params("id")
		var updateUser User
		if err := c.BodyParser(&updateUser); err != nil{
			return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
		}
		users[id] = updateUser

	 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":"User updaeted successfully",
			"user":updateUser,
		})
	})


	app.Delete("/user/:id",func(c *fiber.Ctx) error {
		id := c.Params("id")
		if _,exists := users[id]; !exists {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}
		delete(users, id)

	 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":"User deleted successfully",
		})
	})

	app.Listen(":8080")
   }