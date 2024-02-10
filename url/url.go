package url

import (
	"gocroot/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage)

	page.Get("/auth/phonenumber/:login", controller.GetPhoneNumber)
}
