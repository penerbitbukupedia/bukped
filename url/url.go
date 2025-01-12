package url

import (
	"gocroot/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage)

	page.Get("/auth/phonenumber/:login", controller.GetPhoneNumber)
	page.Get("/auth/userdata", controller.GetHeaderUserData)

	page.Post("/auth/daftar", controller.PostDaftarAuthor)
	page.Post("/auth/upload/image/profil", controller.UploadFotoProfil)
}
