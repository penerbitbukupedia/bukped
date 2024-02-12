package controller

import (
	"gocroot/helper"

	"github.com/gofiber/fiber/v2"
)

func UploadFotoProfil(ctx *fiber.Ctx) error {
	// Parse the form file
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// save to github
	content, response, err := helper.GithubUpload(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"content": content, "response": response})

}
