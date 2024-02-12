package controller

import (
	"gocroot/config"
	"gocroot/helper"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(ctx *fiber.Ctx) error {
	// Parse the form file
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// save to github
	err = helper.GithubUpload(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"filename": config.UploadDir + file.Filename, "content": file.Filename})

}
