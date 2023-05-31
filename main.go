package main

import (
	"log"

	"github.com/rizkyriahutabarat/rizkyria/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/rizkyriahutabarat/rizkyria/url"

	"github.com/gofiber/fiber/v2"
	

	_ "github.com/rizkyriahutabarat/rizkyria/docs"

// @title TES SWAG
// @version 1.0
// @description This is a sample server.

// @contact.name API Support
// @contact.url http://github.com/rizkyriahutabarat/
// @contact.email 1214023@std.ulbi.ac.id

// @host hutabarat.herokuapp.com
// @BasePath /
// @schemes https http
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}

