package controller

import (
	"gocroot/config"
	"gocroot/model"

	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/watoken"
)

func GetPhoneNumber(c *fiber.Ctx) error {
	var author model.Author
	author.Phone = watoken.DecodeGetId(config.PublicKey, c.Params("login"))
	return c.JSON(author)
}

func GetHeaderPhoneNumber(c *fiber.Ctx) error {
	var h model.Login
	c.ReqHeaderParser(&h)
	var author model.Author
	author.Phone = watoken.DecodeGetId(config.PublicKey, h.Login)
	return c.JSON(author)
}
