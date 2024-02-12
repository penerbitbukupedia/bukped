package controller

import (
	"gocroot/config"
	"gocroot/model"

	"github.com/aiteung/atdb"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPhoneNumber(c *fiber.Ctx) error {
	var author model.Author
	author.Phone = watoken.DecodeGetId(config.PublicKey, c.Params("login"))
	return c.JSON(author)
}

func GetHeaderUserData(c *fiber.Ctx) error {
	var h model.Login
	err := c.ReqHeaderParser(&h)
	if err != nil {
		return fiber.ErrServiceUnavailable
	}
	phonenumber := watoken.DecodeGetId(config.PublicKey, h.Login)
	if phonenumber == "" {
		return fiber.ErrForbidden
	}
	author := atdb.GetOneDoc[model.Author](config.Mongoconn, "author", bson.M{"phone": phonenumber})
	author.Phone = phonenumber
	return c.JSON(author)
}

func PostDaftarAuthor(c *fiber.Ctx) error {
	var h model.Login
	err := c.ReqHeaderParser(&h)
	if err != nil {
		return fiber.ErrServiceUnavailable
	}
	phonenumber := watoken.DecodeGetId(config.PublicKey, h.Login)
	if phonenumber == "" {
		return fiber.ErrForbidden
	}
	author := atdb.GetOneDoc[model.Author](config.Mongoconn, "author", bson.M{"phone": phonenumber})
	author.Phone = phonenumber
	if err := c.BodyParser(&author); err != nil {
		if err != nil {
			return fiber.ErrBadRequest
		}
		errors := model.ValidateAuthorStruct(author)
		if errors != nil {
			return c.JSON(errors)
		}
	}
	if author.ID.String() == "000000000000000000000000" {
		insertedid := atdb.InsertOneDoc(config.Mongoconn, "author", author)
		return c.JSON(insertedid)
	}
	return c.JSON(author)
}
