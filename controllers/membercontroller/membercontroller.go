package membercontroller

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message" : "method get / Index",
	})
}

func Show(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message" : "method get / Show",
	})
}

func Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message" : "method post / Create",
	})
}

func Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message" : "method put / Update",
	})
}

func Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message" : "method delete / Delete",
	})
}