package main

import (
	"log"

	"github.com/rizkyriahutabarat/rizkyria/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/rizkyriahutabarat/rizkyria/url"

	"github.com/gofiber/fiber/v2"
	

	"github.com/rizkyriahutabarat/rizkyria/docs"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}

