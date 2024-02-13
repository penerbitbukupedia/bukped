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
	fileName := "anu.jpg"
	// save to github
	content, _, err := helper.GithubUpload(file, fileName)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"content": content, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(content)

}
