package controllers

import "github.com/gofiber/fiber/v2"

func GeneratePages(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}
