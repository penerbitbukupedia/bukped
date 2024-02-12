package controller

import (
	"fmt"
	"gocroot/config"
	"gocroot/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadFotoProfil(ctx *fiber.Ctx) error {
	// Parse the form file
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Save the uploaded file to the server
	id := uuid.New()
	fname := id.String() + ".jpg"
	err = helper.SaveUploadedFile(file, config.UploadDir, fname)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	filename := fmt.Sprintf("%s/%s", config.UploadDir, fname)
	// save to github
	content, response, err := helper.GithubUpload(filename)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"content": content, "response": response, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"content": content, "response": response})

}
