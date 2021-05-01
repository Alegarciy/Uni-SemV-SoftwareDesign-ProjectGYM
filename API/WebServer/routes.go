package WebServer

import (
	"API/WebServer/Controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){
	app.Get("/", Controllers.Start)

	general := app.Group("/general")
	client := app.Group("/client")
	admin := app.Group("/admin")

	general.Post("/login", Controllers.Login)
	general.Get("/activeSchedule", Controllers.GetActiveSchedule)

	client.Get("/userInfo", Controllers.GetUserInfo)
	client.Get("/reservedSessions", Controllers.GetReservedSessions)
	client.Post("/bookSession", Controllers.BookSession)

	admin.Get("/preliminarySchedule", Controllers.GetPreliminarySchedule)
	admin.Post("/insertPreliminarySession", Controllers.InsertPreliminarySession)
	admin.Post("/deletePreliminarySession", Controllers.DeletePreliminarySession)
	admin.Post("/confirmPreliminarySchedule", Controllers.ConfirmPreliminarySchedule)
}


