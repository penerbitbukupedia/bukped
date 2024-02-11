package main

import (
	"log"
	"time"

	"gocroot/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"gocroot/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	// Set the default time zone
	time.Local = loc

	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
