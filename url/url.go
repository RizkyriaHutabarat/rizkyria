package url

import (
	"github.com/rizkyriahutabarat/rizkyria/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)
	page.Get("/presensi", controller.GetPresensiold)
	page.Get("/user", controller.GetUser)
	page.Get("/pendidikan", controller.GetPendidikan)
	page.Get("/pengalaman", controller.GetPengalaman)
	page.Get("/skill", controller.GetSkill)
	page.Get("/profile", controller.GetProfile)
	page.Get("/all", controller.GetAll)
	page.Get("/inuser", controller.InsertUser)
	page.Get("/inpendidikan", controller.InsertPendidikan)
	page.Get("/inpengalaman", controller.InsertPengalaman)
	page.Get("/inskill", controller.InsertSkill)
	page.Get("/inprofile", controller.InsertProfile)
}
