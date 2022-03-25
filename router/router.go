package router

import (
	"fmt"
	ct "main/controllers"
	mw "main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	fmt.Println("Server/Router is running...")

	app.Post("/", mw.Auth, ct.Index)

	app.Post("/reboot", mw.Auth, ct.BashReboot)

}
