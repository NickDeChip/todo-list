package utility

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LogErrorAndSendStatus(c *fiber.Ctx, err error, status int) error {
	log.Printf("%s \n", err.Error())
	return c.SendStatus(status)
}
