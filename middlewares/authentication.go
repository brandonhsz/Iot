package middlewares

import (
	st "main/structs"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	auth := new(st.AuthStruct)

	err := c.BodyParser(auth)
	if err != nil {
		return err
	}
	//in a future add jwt validation
	if auth.Secret == os.Getenv("SECRET") {
		return c.Next()
	} else {
		return c.Status(401).SendString("Unauthorized")
	}

}
