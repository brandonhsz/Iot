package controllers

import (
	"fmt"
	sv "main/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func BashReboot(c *fiber.Ctx) error {
	fmt.Printf("Rebooting the system...\n")
	time.Sleep(time.Second * 10)
	sv.Reboot()
	return c.SendString("Rebooting the system...")
}
