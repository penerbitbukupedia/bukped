package controller

import (
	"gocroot/config"
	"gocroot/helper"
	"gocroot/model"

	"github.com/aiteung/atdb"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func UploadFotoProfil(ctx *fiber.Ctx) error {
	var h model.Login
	err := ctx.ReqHeaderParser(&h)
	if err != nil {
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "token tidak ada"})
	}
	phonenumber := watoken.DecodeGetId(config.PublicKey, h.Login)
	if phonenumber == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "token tidak dikenali"})
	}
	author := atdb.GetOneDoc[model.Author](config.Mongoconn, "author", bson.M{"phone": phonenumber})
	if author.Phone == "" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}
	author.Phone = phonenumber
	// Parse the form file
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// save to github
	content, _, err := helper.GithubUpload(config.GitHubAuthorName, config.GitHubAuthorEmail, file, "penerbitbukupedia", "foto", "main", author.ID.String()+".jpg", true)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"content": content, "error": err.Error()})
	}
	res, err := atdb.UpdateDoc(config.Mongoconn, "author", bson.D{{"_id", author.ID}}, bson.D{{"$set", bson.D{{"photo", *content.Content.DownloadURL}}}})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"error": err.Error(), "result": res.MatchedCount})

}
